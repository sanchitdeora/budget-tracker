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
import { capitalizeFirstLowercaseRest, transformDateFormatToMmDdYyyy } from '../../utils/StringUtils';
import './Transactions.scss';
import { IconButton } from '@mui/material';
import ReusableTransactionDialog from './ReusableTransactionDialog';
import axios from 'axios';


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
        };
        this.getAllTransactions()
    };

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

    getEmptyTransaction = () => {
        return {
            transaction_id: '',
            title: '',
            category: '',
            amount: 0,
            date: new Date(),
            type: false,
            account: '',
            note: '',
        }
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
            });
        } else {
            this.setState({
                allTransactions: [],
            });
        }
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
        console.log('The transaction edit form was submitted with the following data:', transactionBody);
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


    // render functions

    renderEditTransactionDialogBox = () => {
        return (
            <ReusableTransactionDialog
                title={'Edit Transaction'}
                isDialogOpen={this.state.isEditDialogOpen}
                currentTransaction={this.state.allTransactions.find(item => item.transaction_id === this.state.transaction_id)}
                // currentTransaction={{}}
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
                // currentTransaction={this.getEmptyTransaction}
                currentTransaction={{}}
                handleChange={this.handleChange}
                handleClose={this.handleCreateClose}
                submitMethod={this.submitCreateTransaction}
            />
        )
    }

    render() {
        return (
            <div className='transactions-inner-container'>
                <div className='header'>
                    Transactions
                </div>
                <div className='transactions-box'>
                    <div className='transaction-category-box'>
                        <List sx={{ width: '100%', bgcolor: 'background.paper' }}>
                            {this.state.allTransactions.length ? <p></p> : <p>Oops! No Transactions entered</p>}
                            {this.state.allTransactions.map(data => (
                                <div className='transaction'>
                                    <ListItem key={data.transaction_id} id={data.transaction_id} alignItems='flex-start'>
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
                                                    <i>{data.note}</i>
                                                </Typography>

                                            </React.Fragment>}
                                        />
                                        <ListItemText
                                            primary={(data.type ? '' : '-')+ '$' + data.amount}
                                            secondary={data.date.substring(0, 10)}
                                        />
                                    </ListItem>
                                    <Box
                                        display={'flex'}
                                        justifyContent={'flex-end'}
                                        marginRight='5%'
                                    >
                                        <IconButton 
                                            onClick={this.handleEditTransactionOpen.bind(this, data.transaction_id)}>
                                            <ModeEditIcon />
                                        </IconButton>

                                        <div className='fun'>{this.state.transaction_id}</div>

                                        {this.state.transaction_id === data.transaction_id ? this.renderEditTransactionDialogBox() : ""}
                                        <IconButton 
                                            onClick={this.handleDeleteTransactionOpen.bind(this, data.transaction_id)}>
                                                <DeleteIcon />
                                        </IconButton>
                                    </Box>
                                    <Divider variant='middle' component='li' />
                                </div>
                            ))}
                        </List>
                    </div>
                </div>
                <div className='create-button'>
                    <IconButton size='large' onClick={this.handleCreateTransactionOpen}>
                        <AddCircleIcon />
                    </IconButton>
                </div>
                
                {this.renderCreateTransactionDialogBox()}
            </div>
        );
    }
}

export default Transactions;