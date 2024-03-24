import React from 'react';
import List from '@mui/material/List';

import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';

import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Chip from '@mui/material/Chip';

import { PieChart, Pie, Tooltip, Cell, ResponsiveContainer, Legend } from 'recharts';
import { DateRangePicker } from 'react-dates';
import { getTxChartData } from './TransactionPieChart';

import ReusableTransactionDialog from './ReusableTransactionDialog';
import { capitalizeFirstLowercaseRest, findCategoryById, transformDateFormatToMmDdYyyy } from '../../utils/StringUtils';
import { getMenuItemsByTitle } from '../../utils/menuItems';
import { CATEGORY_MAP, TRANSACTIONS } from '../../utils/GlobalConstants';

import axios from 'axios';
import moment from 'moment';
import './Transactions.scss';

const RADIAN = Math.PI / 180;
const renderCustomizedLabel = ({ cx, cy, midAngle, innerRadius, outerRadius, percent, index }) => {
    const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
    const x = cx + radius * Math.cos(-midAngle * RADIAN);
    const y = cy + radius * Math.sin(-midAngle * RADIAN);
    // console.log("all chart params: ", cx, cy, midAngle, innerRadius, outerRadius, percent, index, radius, x, y)

    return (
        <text x={x} y={y} fill="#2dfb86" textAnchor={x > cx ? 'start' : 'end'} dominantBaseline="central">
            {`${(percent * 100).toFixed(0)}%`}
        </text>
    );
};

class Transactions extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allTransactions: [],
            transaction_id: '',
            title: '',
            category: '',
            amount: 0,
            type: false,
            account: '',
            date: new Date(),
            note: '',
            isCreateDialogOpen: false,
            isEditDialogOpen: false,

            chartData: [],

            filterByCategories: [],
            filterByDate: {
                isActive: false,
                startDate: null,
                endDate: null,
            },
            filteredTransactions: [],
        };
        this.props.setNavBarActive(getMenuItemsByTitle(TRANSACTIONS))
        this.getAllTransactions()
    };

    // handlers

    setChartData = (data) => {
        this.setState({chartData: getTxChartData(data)});
    }
    
    cleanTransactionState = () => {
        this.setState({
            transaction_id: '',
            title: '',
            category: '',
            amount: 0,
            date: new Date(),
            type: false,
            account: '',
            note: '',
        })
    }
    
    handleChange = (event) => {
        let value = event.target.value;
        let name = event.target.name;
        if (name === 'type') {
            value = value === 'credit' ? true : false  
            // console.log('Name: ' + name + ' value: ' + value)
            this.setState({
                [name]: value,
            });
        }
        if (name === 'due_date') {
            value = transformDateFormatToMmDdYyyy(value);
            // console.log("Onchange | name: "+name+" value: ", value);
        }
        this.setState({
            [name]: value,
        });
    }

    handleFilterDateRangeChange = (input) => {
        console.log("change here:", input, this.state.filterByDate.startDate, this.state.filterByDate.endDate);
        this.setState({focusedInput: input})
        this.calculateFilteredTransactions()
    }

    handleFilterCategoryChange = (event) => {
        const {
            target: { value },
        } = event;
        console.log("updating filter by category: ", value)
        this.setState({ filterByCategories: value});
        this.calculateFilteredTransactions(null, value);
    };

    getSelectedTransaction = () => {
        let transaction = this.state.allTransactions.find(item => item.transaction_id === this.state.transaction_id);
        if (transaction === undefined) {
            transaction = {}
        }
        return transaction;
    }


    // get transaction

    async getAllTransactions() {
        let res = await axios.get('/api/transactions');
        console.log('get all transactions: ', res.data.body)
        if (res.data.body != null)
        {
            // sort transactions by date
            let sortedTransactions = res.data.body.sort((p1, p2) => (p1.date < p2.date) ? 1 : (p1.date > p2.date) ? -1 : 0);

            this.setState({
                allTransactions: sortedTransactions,
                filteredTransactions: sortedTransactions,
            });
            this.calculateFilteredTransactions(sortedTransactions);
        } else {
            this.setState({
                allTransactions: [],
            });
        }

        console.log("data here: ", getTxChartData(res.data.body))
    }


    // create transaction

    handleCreateTransactionOpen = () => {
        this.setState({
            isCreateDialogOpen: true
        });
    }

    submitCreateTransaction = () => {
        let date = new Date(this.state.date);
        let transactionBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount': parseFloat(this.state.amount),
            'date': date,
            'type': this.state.type,
            'account': this.state.account,
            'note': this.state.note,
        }
        console.log('The transaction create form was submitted with the following data:', transactionBody,);
        this.postTransactionRequest(transactionBody)
        this.handleCreateClose()
    }

    async postTransactionRequest(transactionBody) {
        let res = await axios.post('/api/transaction', transactionBody);
        console.log("Post transaction response", res);
        this.getAllTransactions();
    }

    handleCreateClose = () => {
        this.cleanTransactionState()
        this.setState({
            isCreateDialogOpen: false
        });
    };


    // edit transaction

    handleEditTransactionOpen = (id) => {
        console.log('Edit transaction id: ', id)
        let transaction = this.state.allTransactions.find(item => item.transaction_id === id);
        this.setState({
            transaction_id: id,
            title: transaction.title,
            category: transaction.category,
            amount: transaction.amount,
            date: transaction.date,
            type: transaction.type,
            account: transaction.account,
            note: transaction.note,

            isEditDialogOpen: true
        });
    }

    submitEditTransaction = () => {
        let date = new Date(this.state.date);
        let transactionBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount': parseFloat(this.state.amount),
            'date': date,
            'type': this.state.type,
            'account': this.state.account,
            'note': this.state.note,
        }
        console.log('The edit transaction form was submitted with the following data:', transactionBody);
        this.putTransactionRequest(transactionBody)
        this.handleEditClose()
    }

    async putTransactionRequest(transactionBody) {
        let res = await axios.put('/api/transaction/'+this.state.transaction_id, transactionBody);
        console.log("put transaction response", res);
        this.getAllTransactions();
    }

    handleEditClose = () => {
        this.cleanTransactionState()
        this.setState({
            isEditDialogOpen: false
        });
    };

    // delete transaction

    handleDeleteTransactionOpen = (id) => {
        console.log('Delete transaction id: ', id)
        this.deleteTransactionRequest(id)
    }

    async deleteTransactionRequest(id) {
        let res = await axios.delete('/api/transaction/'+id);
        console.log("delete transaction response", res);
        this.getAllTransactions();
    }


    // utils

    calculateFilteredTransactions = (filteredTx, fCategory, fDateObj) => {
        if (filteredTx === undefined || filteredTx === null || filteredTx.length === 0)
            filteredTx = this.state.allTransactions
        if (fCategory === undefined || fCategory === null) {
            fCategory = this.state.filterByCategories
        }
        if (fDateObj === undefined || fDateObj === null) {
            fDateObj = this.state.filterByDate
        }
        console.log("Calculating Filter transactions", filteredTx, fCategory, fDateObj);

        if (fCategory.length > 0) {
            filteredTx = filteredTx.filter(x => fCategory.includes(findCategoryById(x.category)))
        }
        if (fDateObj.isActive) {
            var startDate = fDateObj.startDate === null ? moment(0) : fDateObj.startDate
            var endDate = fDateObj.endDate === null ? moment() : fDateObj.endDate
            console.log("Filter transactions for date======================", fDateObj, startDate, endDate)
            filteredTx = filteredTx.filter(
                x => startDate.isBefore(moment(x.date)) && endDate.isAfter(moment(x.date))
            )
        }
        this.setState({
            filteredTransactions: filteredTx,
        })

        this.setChartData(filteredTx);
    }


    // render functions

render() {
        return (
            <div className='transactions-container'>
                <h2 className='header'>
                    {TRANSACTIONS}
                </h2>
                <div className='transactions-filter-by-date'>
                    {this.renderFilterByDates()}
                </div>                    
                <div className='transactions-filter-by-category'>
                    {this.renderFilterByCategory()}
                </div>
                <div className='transactions-chart'>
                    Pie Chart by Categories
                    {this.renderBasicPie()}
                </div>
                <div className='transactions-box'>
                    
                    <div className='transaction-create-button'>
                        <IconButton size='large' onClick={this.handleCreateTransactionOpen}>
                            <AddCircleIcon />
                        </IconButton>
                    </div>
                    {this.renderCreateTransactionDialogBox()}
                    <List sx={{ width: '100%' }}>
                        {this.state.filteredTransactions.length ? <p></p> : <p>Oops! No Transactions entered</p>}
                        {this.state.filteredTransactions.map(data => (
                            this.renderSingleTransaction(data)
                        ))}
                    </List>
                </div>

            </div>
        );
    }

    renderSingleTransaction(transaction) {
        return (
            <div className='transaction'>
                <Grid container spacing={0}>
                    <Grid item xs={7}>
                        {capitalizeFirstLowercaseRest(transaction.title)}
                    </Grid>
                    <Grid item xs={4}>
                        {(transaction.type ? '' : '-')+ '$' + transaction.amount}
                    </Grid>
                    <Grid item xs={1}>
                        <IconButton 
                            size='small'
                            onClick={this.handleEditTransactionOpen.bind(this, transaction.transaction_id)}>
                            <ModeEditIcon />
                        </IconButton>
                        {this.renderEditTransactionDialogBox()}
                    </Grid>
                    <Grid className='secondary-transaction-detail' item xs={7}>
                        <i>{transaction.note}</i>
                    </Grid>
                    <Grid className='secondary-transaction-detail' item xs={4}>
                        {transaction.date.substring(0, 10)}
                    </Grid>
                    <Grid item xs={1}>
                        <IconButton className='delete-button'
                            size='small'
                            onClick={this.handleDeleteTransactionOpen.bind(this, transaction.transaction_id)}>
                                <DeleteIcon />
                        </IconButton>
                    </Grid>
                </Grid>
            </div>
        )
    }
    
    renderEditTransactionDialogBox = () => {
        return (
            <ReusableTransactionDialog
                title={'Edit Transaction'}
                isDialogOpen={this.state.isEditDialogOpen}
                currentTransaction={this.getSelectedTransaction()}
                handleChange={this.handleChange}
                handleClose={this.handleEditClose}
                submitMethod={this.submitEditTransaction}
            />
        )
    }

    renderCreateTransactionDialogBox = () => {
        return (
            <ReusableTransactionDialog
                title={'Add Transaction'}
                isDialogOpen={this.state.isCreateDialogOpen}
                currentTransaction={{}}
                handleChange={this.handleChange}
                handleClose={this.handleCreateClose}
                submitMethod={this.submitCreateTransaction}
            />
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

    renderFilterByCategory() {
        return (
            <FormControl sx={{ m: 1, minWidth: 300 }} className='multiselect-form'>
                <InputLabel id="demo-multiple-chip-label">Filter </InputLabel>
                <Select
                    className='ms-select'
                    labelId="demo-multiple-chip-label"
                    id="demo-multiple-chip"
                    multiple
                    autoWidth
                    value={this.state.filterByCategories}
                    onChange={this.handleFilterCategoryChange}
                    input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
                    renderValue={(selected) => (
                        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                        {selected.map((value) => (
                            <Chip key={value} label={value} className='ms-chip' />
                        ))}
                        </Box>
                    )}
                    MenuProps={{ PaperProps: {
                        style: {
                            maxHeight: 48 * 4.5 + 8,
                            width: 250,
                        },
                        }}}
                    >
                    {CATEGORY_MAP.map((category) => (
                        <MenuItem
                            key={category.id}
                            value={category.value}
                            >
                            {category.value}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>
        )
    }

    renderBasicPie() {
        return (
            <ResponsiveContainer>
            <PieChart width={800} height={800}>
            <Legend
                height={'15%'}
                iconType="circle"
                layout="vertical"
                verticalAlign="middle"
                iconSize={10}
                padding={5}
                // formatter={renderColorfulLegendText}
            />
              <Pie
                data={this.state.chartData}
                cx="50%"
                cy="20%"
                labelLine={false}
                label={renderCustomizedLabel}
                outerRadius={'50%'}
                innerRadius={'35%'}
                fill="#ffffff"
                dataKey="value"
                stroke='none'
              >
                {this.state.chartData.map(item => (
                  <Cell style={{outline: 'none'}} key={`cell-${item.category}`} />
                ))}
              </Pie>
            <Tooltip />
            </PieChart>
          </ResponsiveContainer>
        );
    }
}

// const renderColorfulLegendText = (value: string, entry: any) => {
//     return (
//       <span style={{ color: "#596579", fontWeight: 500, padding: "10px" }}>
//         {value}
//       </span>
//     );
//   };

export default Transactions;