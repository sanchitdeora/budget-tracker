import React from 'react';
import List from '@mui/material/List';

import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import CheckIcon from '@mui/icons-material/Check';
import RadioButtonUncheckedIcon from '@mui/icons-material/RadioButtonUnchecked';

import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Chip from '@mui/material/Chip';

import { DateRangePicker } from 'react-dates';

import ReusableBillDialog from './ReusableBillDialog';
import { capitalizeFirstLowercaseRest, findCategoryById, findFrequencyById, transformDateFormatToMmDdYyyy } from '../../utils/StringUtils';
import { CATEGORY_MAP, BILLS } from '../../utils/GlobalConstants';
import { getMenuItemsByTitle } from '../../utils/menuItems';

import axios from 'axios';
import moment from 'moment';
import './Bills.scss';


class Bills extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allBills: [],
            bill_id: '',
            title: '',
            category: '',
            amount_due: 0,
            due_date: new Date(),
            frequency: '',
            is_paid: false,
            note: '',
            isCreateDialogOpen: false,
            isEditDialogOpen: false,

            filterByCategories: [],
            activeFilterStatus: null,
            filterByDate: {
                isActive: false,
                startDate: null,
                endDate: null,
            },
            filteredBills: [],
        };
        this.props.setNavBarActive(getMenuItemsByTitle(BILLS))
        this.getAllBills()
    };

    // handlers

    cleanBillState = () => {
        this.setState({
            bill_id: '',
            title: '',
            category: '',
            amount_due: 0,
            due_date: new Date(),
            frequency: '',
            is_paid: false,
            note: '',
        })
    }

    handleChange = (event) => {
        let value = event.target.value;
        let name = event.target.name;
        if (name === 'due_date') {
            value = transformDateFormatToMmDdYyyy(value);
            console.log("Onchange | name: "+name+" value: ", value);
        }
        this.setState({
            [name]: value,
        });
    }

    handleFilterDateRangeChange = (input) => {
        console.log("change here:", input, this.state.filterByDate.startDate, this.state.filterByDate.endDate);
        this.setState({focusedInput: input});
        this.calculateFilteredBills();
    }

    handleFilterStatusChange = (e) => {
        console.log(e.target);
        const {
            target: { value },
        } = e;
        this.setState({activeFilterStatus: value});
        this.calculateFilteredBills(null, value);
    }

    handleFilterCategoryChange = (event) => {
        const {
            target: { value },
        } = event;
        console.log("updating filter by category: ", value)
        this.setState({ filterByCategories: value});
        this.calculateFilteredBills(null, value);
    };


    // get bill

    async getAllBills() {
        let res = await axios.get('/api/bills');
        console.log('get all bills: ', res.data.body)
        if (res.data.body != null)
        {
            // sort bills by date
            let sortedBills = res.data.body.sort((p1, p2) => (p1.due_date < p2.due_date) ? 1 : (p1.due_date > p2.due_date) ? -1 : 0);

            this.setState({
                allBills: sortedBills,
                filteredBills: sortedBills,
            });
            this.calculateFilteredBills(sortedBills);
        } else {
            this.setState({
                allBills: [],
            });
        }
    }


    // create bill

    handleCreateBillOpen = () => {
        this.setState({
            isCreateDialogOpen: true
        });
    }

    submitCreateBill = () => {
        let due_date = new Date(this.state.due_date);
        let billBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount_due': parseFloat(this.state.amount_due),
            'due_date': due_date,
            'frequency': this.state.frequency,
            'is_paid': this.state.is_paid,
            'note': this.state.note,
        }
        console.log('The create bill form was submitted with the following data:', billBody,);
        this.postBillRequest(billBody)
        this.handleCreateClose()
    }

    async postBillRequest(billBody) {
        let res = await axios.post('/api/bill', billBody);
        console.log("Post Bill response", res);
        this.getAllBills();
    }

    handleCreateClose = () => {
        this.cleanBillState()
        this.setState({
            isCreateDialogOpen: false
        });
    };


    // edit bill

    handleEditBillOpen = (id) => {
        let bill = this.state.allBills.find(item => item.bill_id === id);
        console.log('Edit bill id: ', id, bill)
        this.setState({
            bill_id: id,
            title: bill.title,
            category: bill.category,
            amount_due: bill.amount_due,
            due_date: bill.due_date,
            frequency: bill.frequency,
            is_paid: bill.is_paid,
            note: bill.note,

            isEditDialogOpen: true
        });
    }

    submitEditBill = () => {
        let due_date = new Date(this.state.due_date);
        let billBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount_due': parseFloat(this.state.amount_due),
            'due_date': due_date,
            'frequency': this.state.frequency,
            'is_paid': this.state.is_paid,
            'note': this.state.note,
        }
        console.log('The edit bill form was submitted with the following data:', billBody);
        this.putBillRequest(billBody)
        this.handleEditClose()
    }

    async putBillRequest(billBody) {
        let res = await axios.put('/api/bill/'+this.state.bill_id, billBody);
        console.log("Put bill response", res);
        this.getAllBills();
    }

    handleEditClose = () => {
        this.cleanBillState()
        this.setState({
            isEditDialogOpen: false
        });
    };

    // delete bill

    handleDeleteBillOpen = (id) => {
        console.log('Delete bill id: ', id)
        this.deleteBillRequest(id)
    }

    async deleteBillRequest(id) {
        let res = await axios.delete('/api/bill/'+id);
        console.log("Delete bill response", res);
        this.getAllBills();
    }

    // bill paid

    handleBillPaid = (id, isPaid) => {
        console.log('Print if bill paid for id: ', id, ' isPaid? ', isPaid)
        isPaid ? this.putBillIsPaidRequest(id) : this.putBillIsUnpaidRequest(id)
    }

    async putBillIsPaidRequest(id) {
        let res = await axios.put('/api/bill/updateIsPaid/' + id);
        console.log("Put bill paid response", res);
        this.getAllBills();
    }

    async putBillIsUnpaidRequest(id) {
        let res = await axios.put('/api/bill/updateIsUnpaid/' + id);
        console.log("Put bill unpaid response", res);
        this.getAllBills();
    }

    // utils

    calculateFilteredBills = (filteredBills, fStatus, fDateObj) => {
        if (filteredBills === undefined || filteredBills === null || filteredBills.length === 0)
            filteredBills = this.state.allBills
        if (fDateObj === undefined || fDateObj === null) {
            fDateObj = this.state.filterByDate
        }
        console.log("Calculating Filter bills", filteredBills, fStatus, fDateObj);

        if (fDateObj.isActive) {
            var startDate = fDateObj.startDate === null ? moment(0) : fDateObj.startDate
            var endDate = fDateObj.endDate === null ? moment() : fDateObj.endDate
            console.log("Filter bills for date======================", fDateObj, startDate, endDate)
            filteredBills = filteredBills.filter(
                x => startDate.isBefore(moment(x.due_date)) && endDate.isAfter(moment(x.due_date))
            )
        }
        if (fStatus !== undefined && fStatus !== null) {
            filteredBills = filteredBills.filter(x => x.is_paid === fStatus)
        }
        this.setState({
            filteredBills: filteredBills,
        })

        // this.setChartData(filteredBills);
    }

    // render functions

    render() {
        return (
            <div className='bills-container'>
                <h2 className='header'>
                    {BILLS}
                </h2>
                {/* {this.renderFilterContainer()} */}
                <div className='bills-filter-by-category'>
                    {this.renderFilterByCategory()}
                </div>
                <div className='bills-filter-by-date'>
                    {this.renderFilterByDates()}
                </div>
                <div className='bills-box'>
                    <div className='bills-create-button'>
                        <IconButton size='large' onClick={this.handleCreateBillOpen}>
                            <AddCircleIcon />
                        </IconButton>
                    </div>
                    {this.renderCreateBillDialogBox()}
                    <List sx={{ width: '100%' }}>
                        {this.state.filteredBills.length ? <p></p> : <p>Oops! No Bills entered</p>}
                        {this.state.filteredBills?.map(data => (
                            this.renderSingleBill(data)
                        ))}
                    </List>
                </div>

                <div className='bills-chart'>
                    Pie Chart by Categories
                    {this.renderBasicPie()}
                </div>
            </div>
        );
    }

    renderSingleBill(bill) {
        return (
            <div className='bill'>
                <Grid container spacing={0}>
                    <Grid item xs={7}>
                        {capitalizeFirstLowercaseRest(bill.title)}
                    </Grid>
                    <Grid item xs={4}>
                        {'$' + bill.amount_due}
                    </Grid>
                    <Grid item xs={1}>
                        <IconButton 
                            size='small'
                            onClick={this.handleEditBillOpen.bind(this, bill.bill_id)}>
                            <ModeEditIcon />
                        </IconButton>
                        {this.renderEditBillDialogBox()}
                    </Grid>
                    <Grid className='secondary-bill-detail' item xs={7}>
                        <i>{bill.note}</i>
                    </Grid>
                    <Grid className='secondary-bill-detail' item xs={4}>
                        {bill.due_date.substring(0, 10)}
                    </Grid>
                    <Grid item xs={1}>
                        <IconButton className='delete-button'
                            size='small'
                            onClick={this.handleDeleteBillOpen.bind(this, bill.bill_id)}>
                                <DeleteIcon />
                        </IconButton>
                    </Grid>
                    <Grid className='tertiary-bill-detail' item xs={7}>
                        {findCategoryById(bill.category)}
                    </Grid>
                    <Grid className='tertiary-bill-detail' item xs={4}>
                        {findFrequencyById(bill.frequency)}
                    </Grid>
                    <Grid item xs={1}>
                        <div>
                            {bill.is_paid ?
                            <IconButton 
                                size='small'
                                onClick={this.handleBillPaid.bind(this, bill.bill_id, false)}>
                                <CheckIcon />
                            </IconButton> 
                            :
                            <IconButton 
                                size='small'
                                onClick={this.handleBillPaid.bind(this, bill.bill_id, true)}>
                                <RadioButtonUncheckedIcon />
                            </IconButton> 
                            }
                        </div>
                    </Grid>
                </Grid>
            </div>
        )
    }

    renderEditBillDialogBox = () => {
        return (
            <ReusableBillDialog
                title={'Edit Bill'}
                isDialogOpen={this.state.isEditDialogOpen}
                currentBill={this.state.allBills.find(item => item.bill_id === this.state.bill_id)}
                handleChange={this.handleChange}
                handleClose={this.handleEditClose}
                submitMethod={this.submitEditBill}
            />
        )
    }

    renderCreateBillDialogBox = () => {
        return (
            <ReusableBillDialog
                title={'Add Bill'}
                isDialogOpen={this.state.isCreateDialogOpen}
                currentBill={{}}
                handleChange={this.handleChange}
                handleClose={this.handleCreateClose}
                submitMethod={this.submitCreateBill}
                />
        )
    }

    renderFilterByCategory() {
        return (
            <FormControl sx={{ m: 1, width: 300 }} className='multiselect-form'>
                <InputLabel id="demo-multiple-chip-label">Filter</InputLabel>
                <Select
                    className='ms-select'
                    labelId="demo-multiple-chip-label"
                    id="demo-multiple-chip"
                    value={this.state.activeFilterStatus}
                    label="Age"
                    onChange={this.handleFilterStatusChange}
                >
                    <MenuItem value={true}>Paid</MenuItem>
                    <MenuItem value={false}>Unpaid</MenuItem>
                    <MenuItem sx={{color: 'red !important' }} value={null}>Clear</MenuItem>
                </Select>
            </FormControl>
        )
    }

    renderFilterByDates() {
        return (
            <DateRangePicker
                startDate={this.state.filterByDate.startDate} // momentPropTypes.momentObj or null,
                startDateId="your_unique_start_date_id" // PropTypes.string.isRequired,
                endDate={this.state.filterByDate.endDate} // momentPropTypes.momentObj or null,
                endDateId="your_unique_end_date_id" // PropTypes.string.isRequired,
                onDatesChange={({ startDate, endDate }) => this.setState({ filterByDate: {isActive: true, startDate: startDate, endDate: endDate} })} // PropTypes.func.isRequired,
                focusedInput={this.state.focusedInput} // PropTypes.oneOf([START_DATE, END_DATE]) or null,
                onFocusChange={focusedInput => this.handleFilterDateRangeChange(focusedInput)} // PropTypes.func.isRequired,
                isOutsideRange={day => (moment().diff(day) < 0)}
                />
                )
            }

    renderBasicPie() {
        return (
            <div></div>
        )
    }
}

export default Bills;