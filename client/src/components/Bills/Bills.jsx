import React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import { capitalizeFirstLowercaseRest, changeDateFormatToMmDdYyyy } from '../../utils/StringUtils';
import './Bills.scss';
import { IconButton } from '@mui/material';
import ReusableBillDialog from '../../utils/ReusableBillDialog';
import axios from 'axios';



class Bills extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allBills: [],
            billId: '',
            title: '',
            category: '',
            amount_due: 0,
            due_date: new Date(),
            frequency: '',
            is_paid: false,
            note: '',
            isCreateDialogOpen: false,
            isEditDialogOpen: false,
        };
        this.getAllBills()
    };

    cleanBillState = () => {
        this.setState({
            billId: '',
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
            value = changeDateFormatToMmDdYyyy(value);
            console.log("Onchange | name: "+name+" value: ", value);
        }
        this.setState({
            [name]: value,
        });
    }
        
    // get bill

    async getAllBills() {
        let res = await axios.get('/api/bills');
        if (res.data.body != null)
        {
            this.setState({
                allBills: res.data.body,
            });
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
        console.log("Due date in state: ", this.state.due_date, " date now: ", due_date, "date in UTC: ")
        let billBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount_due': parseFloat(this.state.amount_due),
            'due_date': due_date,
            'frequency': this.state.frequency,
            'is_paid': this.state.is_paid,
            'note': this.state.note,
        }
        console.log('The create form was submitted with the following data:', billBody,);
        this.postBillRequest(billBody)
        this.handleCreateClose()
    }

    async postBillRequest(billBody) {
        let res = await axios.post('/api/bill', billBody);
        console.log(res);
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
        console.log('Edit id: ', id)
        this.setState({
            billId: id,
            isEditDialogOpen: true
        });
    }

    submitEditBill = () => {
        let due_date = new Date(this.state.due_date);
        console.log("Due date in state: ", this.state.due_date, " date now: ", due_date)
        let billBody = {
            'title': this.state.title,
            'category': this.state.category,
            'amount_due': parseFloat(this.state.amount_due),
            'due_date': due_date,
            'frequency': this.state.frequency,
            'is_paid': this.state.is_paid,
            'note': this.state.note,
        }
        console.log('The edit form was submitted with the following data:', billBody);
        this.putBillRequest(billBody)
        this.handleEditClose()
    }

    async putBillRequest(billBody) {
        let res = await axios.put('/api/bill/'+this.state.billId, billBody);
        console.log(res);
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
        console.log('Delete id: ', id)
        this.deleteBillRequest(id)
    }

    async deleteBillRequest(id) {
        let res = await axios.delete('/api/bill/'+id);
        console.log(res);
        this.getAllBills();
    }

    // bill paid

    handleBillPaid = (id, isPaid) => {
        console.log('Print if paid for id: ', id, ' isPaid? ', isPaid)
        isPaid ? this.putBillIsPaidRequest(id) : this.putBillIsUnpaidRequest(id)
    }

    async putBillIsPaidRequest(id) {
        let res = await axios.put('/api/bill/updateIsPaid/' + id);
        console.log(res);
        this.getAllBills();
    }

    async putBillIsUnpaidRequest(id) {
        let res = await axios.put('/api/bill/updateIsUnpaid/' + id);
        console.log(res);
        this.getAllBills();
    }


    render() {
        return (
            <div className='bills-inner-container'>
                <div className='header'>
                    Bills
                </div>
                <div className='bills-box'>
                    <div className='bills-category-box'>
                        <List sx={{ width: '100%', bgcolor: 'background.paper' }}>
                            {this.state.allBills.length ? <p></p> : <p>Oops! No Bills entered</p>}
                            {this.state.allBills?.map(data => (
                                <div className='bill'>
                                    <ListItem key={data.bill_id} id={data.bill_id} alignItems='flex-start'>
                                        <ListItemText
                                            style={{width: '65%'}}
                                            primary={capitalizeFirstLowercaseRest(data.title)}
                                            secondary={<React.Fragment>
                                                <Typography
                                                    sx={{ display: 'inline' }}
                                                    component='span'
                                                    variant='body2'
                                                    color='text.primary'
                                                    >
                                                    {data.note}
                                                </Typography>
                                            </React.Fragment>} />
                                    <ListItemText
                                            primary={'$' + data.amount_due}
                                            secondary={data.due_date.substring(0, 10)}
                                        />
                                    </ListItem>
                                    <Box
                                        display={'flex'}
                                        justifyContent={'space-between'}
                                        marginRight='15%'
                                        marginLeft='2%'
                                    >
                                        <div>
                                            {data.is_paid ? 
                                            <Button variant='text' onClick={this.handleBillPaid.bind(this, data.bill_id, false)}>
                                                Unpaid
                                            </Button> 
                                            :
                                            <Button variant='text' onClick={this.handleBillPaid.bind(this, data.bill_id, true)}>
                                                Paid
                                            </Button>
                                            }
                                        </div>
                                        <div>
                                            <IconButton onClick={this.handleEditBillOpen.bind(this, data.bill_id)}>
                                                <ModeEditIcon />
                                            </IconButton>
                                            <ReusableBillDialog
                                                title={'Edit Bill'}
                                                isDialogOpen={this.state.isEditDialogOpen}
                                                handleChange={this.handleChange}
                                                handleClose={this.handleEditClose}
                                                submitMethod={this.submitEditBill}
                                                />
                                            <IconButton edge='end'onClick={this.handleDeleteBillOpen.bind(this, data.bill_id)}>
                                                    <DeleteIcon />
                                            </IconButton>
                                        </div>
                                    </Box>
                                    <Divider variant='middle' component='li' />
                                </div>
                            ))}
                        </List>
                    </div>
                </div>
                <div className='create-button'>
                    <IconButton size='large' onClick={this.handleCreateBillOpen}>
                        <AddCircleIcon />
                    </IconButton>
                </div>
                
                <ReusableBillDialog
                    title={'Add Bill'}
                    isDialogOpen={this.state.isCreateDialogOpen}
                    handleChange={this.handleChange}
                    handleClose={this.handleCreateClose}
                    submitMethod={this.submitCreateBill}
                />
            </div>
        );
    }
}

export default Bills;