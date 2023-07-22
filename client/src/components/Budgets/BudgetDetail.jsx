import React from 'react';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import LinearProgress from '@mui/material/LinearProgress';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { FormControl, TextField, DialogActions, DialogContent, DialogTitle, Dialog, InputAdornment, FormGroup } from '@mui/material';


import './BudgetDetail.scss';
import { IconButton, Button } from '@mui/material';
import ReusableBudgetDialog from '../../utils/ReusableBudgetDialog';
import { capitalizeFirstLowercaseRest, transformDateFormatToMmmDdYyyy } from '../../utils/StringUtils';
import axios from 'axios';

class BudgetDetail extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            budget_id: '',
            goal_id: '',
            current_goal_amount: 0,
            transactions: [],
            isEditDialogOpen: false,
            isAddGoalAmountDialogOpen: false,
        };

        this.getTransactionByDate(this.props.budget)
        console.log("props for budget detail", this.props);
        console.log("states for budget detail", this.state);
    };

    // get transactions by date

    async getTransactionByDate(budget) {
        let startEpoch = Date.parse(budget.creation_time)
        let endEpoch = Date.parse(budget.expiration_time)

        console.log("start and end: ", startEpoch, endEpoch)

        let res = await axios.get('/api/transactions/startTime/' + startEpoch + '/endTime/' + endEpoch);
        console.log('get transactions by time range: ', res.data.body)
        if (res.data.body != null)
        {
            this.setState({
                transactions: res.data.body,
            });
        }
    }

    // state handlers

    cleanBudgetState = () => {
        this.setState({
            budget_id: '',
            goal_id: '',
            isEditDialogOpen: false,
            isAddGoalAmountDialogOpen: false
        })
    }

    handleEditBudgetOpen = (id) => {
        console.log('Edit budget id for budgetDetail: ', id)
        this.setState({
            budget_id: id,
            isEditDialogOpen: true
        });
    }
    
    handleEditClose = () => {
        this.cleanBudgetState()
        this.setState({
            isEditDialogOpen: false
        });
    };

    handleAddGoalAmountOpen = (id) => {
        console.log('Open add Goal Amount in budgetDetail');
        this.setState({
            goal_id: id,
            isAddGoalAmountDialogOpen: true
        });
    }
    
    handleAddGoalAmountClose = () => {
        console.log('Close add Goal Amount in budgetDetail');
        this.cleanBudgetState()
        this.setState({
            isAddGoalAmountDialogOpen: false
        });
    };
    
    handleGoalAmountChange = (event) => {
        console.log('Handle Goal Amount in budgetDetail: ', event.target.value);
        this.setState({
            current_goal_amount: parseFloat(event.target.value)
        });
    }
    
    submitEditMethod = () => {
        this.props.submitMethod()
        this.handleEditClose()
    }

    submitGoalAmountMethod = () => {
        console.log('Submit add Goal Amount in budgetDetail: ', this.state.goal_id);
        var goal_map = this.props.budget.goal_map

        var objIndex = goal_map.findIndex(x => x.id === this.state.goal_id)
        goal_map[objIndex].current_amount += this.state.current_goal_amount
        console.log("Found: ", goal_map)

        let budgetBody = {
            'budget_name': this.props.budget.budget_name,
            'income_map': this.props.budget.income_map,
            'expense_map': this.props.budget.expense_map,
            'goal_map': this.props.budget.goal_map,
            'frequency': this.props.budget.frequency,
            'savings': parseFloat(this.props.budget.savings),
        }
        console.log('The budgets edit form was submitted with the following data:', budgetBody);
        this.putBudgetRequest(budgetBody)
        this.handleAddGoalAmountClose()
    }

    async putBudgetRequest(budgetBody) {
        let res = await axios.put('/api/budget/'+this.props.budget.budget_id, budgetBody);
        console.log("Put budgets response", res);
        // this.props.getAllBudgets();
    }

    // utils

    getTimeRange = (creation_time, expiration_time) => {
        console.log('Get time range creation time: ', creation_time, 'expiration time: ', expiration_time)
        if (new Date(expiration_time).getTime() < new Date(creation_time).getTime()) {
            return ("From: " + transformDateFormatToMmmDdYyyy(creation_time))
        }
        return (transformDateFormatToMmmDdYyyy(creation_time) + " - " + transformDateFormatToMmmDdYyyy(expiration_time))
    }

    getProgressPercentage = (currAmount, totalAmount) => {
        // console.log('Get progress percentage: ', currAmount, 'total amount: ', totalAmount)
        currAmount = currAmount && currAmount > 0 ? currAmount : 0;

        if (totalAmount <= 0 || currAmount >= totalAmount) {
            return 100;
        }

        return Math.round(currAmount / totalAmount * 100);
    }

    
    calculateTotalSavings = () => {
        let incomeAmount = this.calculateBudgetMap(this.props.budget.income_map)
        let expenseAmount = this.calculateBudgetMap(this.props.budget.expense_map)
        let goalAmount = this.calculateBudgetMap(this.props.budget.goal_map)
        return [incomeAmount, expenseAmount, goalAmount, (incomeAmount - expenseAmount - goalAmount)]
    }

    calculateBudgetMap(budget_map) {
        let amount = 0
        budget_map.forEach(element => {
            // console.log(element)
            amount += element.current_amount
        });
        return amount
    }



    // render functions

    render() {
        return (
            <div className='budgets-inner-container'>

                <div className='budgets-box'>
                    {this.props.budget ? <p></p> : <h3>Add redirect back or error</h3>}
                    <div className='budget'>                            
                        <Box
                            className='budget-detail-header-box'
                            display={'flex'}
                            justifyContent={'space-between'}
                            marginLeft='auto'
                            alignItems='center'
                            >
                            <IconButton style={{padding: '2%'}} onClick={this.props.handleBudgetClose}>
                                <ArrowBackIosNewIcon />
                            </IconButton>
                            
                            <h2>{this.props.budget.budget_name}</h2>

                            <IconButton style={{marginRight: '2%', padding: '2%'}} onClick={this.handleEditBudgetOpen.bind(this, this.props.budget.budget_id)}>
                                <ModeEditIcon />
                            </IconButton>
                            <ReusableBudgetDialog
                                title={'Edit Budget'}
                                isDialogOpen={this.state.isEditDialogOpen}
                                handleInputChange={this.props.handleInputChange}
                                handleClose={this.handleEditClose}
                                currentBudget={this.props.budget}
                                submitMethod={this.submitEditMethod}
                            />
                        </Box>
                        <Divider variant='middle' />
                        <Box
                            marginTop={'3%'}
                            display={'flex'}
                            justifyContent={'space-between'}
                            marginLeft='auto'
                            alignItems='center'
                        >
                            <div className='budget-detail-other-box'>{this.getTimeRange(this.props.budget.creation_time, this.props.budget.expiration_time)}</div>
                            <div className='budget-detail-other-box'>Frequency: {capitalizeFirstLowercaseRest(this.props.budget.frequency)}</div>
                        </Box>


                        {this.renderBudgetMaps('Income', this.props.budget.income_map, false)}
                        {this.renderBudgetMaps('Expense', this.props.budget.expense_map, false)}
                        {this.renderBudgetMaps('Goals', this.props.budget.goal_map, true)}

                        {/* update savings */}
                        {this.renderTotalSavings()}

                        {this.renderTransactionDetails('Transactions', this.state.transactions)}
                    </div>
                </div>
            </div>
        );
    }

    renderBudgetMaps = (title, dataMap, goalFlag) =>
    {
        return (
            <div>
                <div className='budget-detail-map-box'>
                {/* <Divider variant='middle' /> */}
                    <h3>{title}</h3>

                    {Object.keys(dataMap).map(key => (
                        <div className='budgets-category-box' key={key} >
                            <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                <ListItemText
                                    style={{width: '80%'}}
                                    primary={dataMap[key].name}
                                />
                                {goalFlag ?
                                <div>
                                    <Button size='large' onClick={this.handleAddGoalAmountOpen.bind(this, dataMap[key].id)}>
                                        Add Amount
                                    </Button>
                                    {this.renderAddGoalAmount(dataMap[key])}
                                </div>
                                :""}
                            </div>
                            <div className='budgets-list-item-text-container' style={{display: 'flex', marginTop: '0px'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={'$ ' + (dataMap[key].current_amount ? dataMap[key].current_amount : 0)}
                                />
                                <ListItemText
                                    style={{textAlign: 'right'}}
                                    primary={'$ ' + dataMap[key].amount}
                                />
                            </div>
                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                <Box sx={{ width: '100%', mr: 1 }}>
                                    <LinearProgress variant="determinate" color="primary" value={this.getProgressPercentage(dataMap[key].current_amount, dataMap[key].amount)}/>
                                </Box>
                                <Box sx={{ minWidth: 35 }}>
                                    <Typography variant="body2" color="text.secondary">{`${this.getProgressPercentage(dataMap[key].current_amount, dataMap[key].amount)}%`}</Typography>
                                </Box>
                            </Box>
                        </div>
                    ))}

                    <br></br>
                    </div>
                    <br></br>
            </div>
        )

    }
    
    renderAddGoalAmount = (goal) => {
        return (
        <Dialog
            className='budget-dialog'
            open={this.state.isAddGoalAmountDialogOpen}
            fullWidth={true}
            onClose={this.handleAddGoalAmountClose}
        >
            <DialogTitle textAlign={'center'}>{`Add Goal Amount`}</DialogTitle>
            <DialogContent>
                <FormGroup>
                    <FormControl className='goal-input-group' sx={{ width: 300 }}>
                        <br></br>
                        <TextField 
                        defaultValue={0}
                        name='current_amount'
                        type='number'
                        InputProps={{
                            startAdornment: 
                            <InputAdornment disableTypography position="start">
                                $</InputAdornment>,
                            inputMode: 'numeric', pattern: '[0-9]*' 
                        }}
                        label="Add Amount to Goal"
                        className='goal-input-box'
                        onChange={this.handleGoalAmountChange}
                        variant="outlined" />
                    </FormControl>
                </FormGroup>
            </DialogContent>
            <DialogActions>
                <button
                    type='submit'
                    className='budget-submit-btn'
                    onClick={this.submitGoalAmountMethod}>Submit
                </button>						
                <button
                    type='submit'
                    className='close-budget-submit-btn'
                    onClick={this.handleAddGoalAmountClose}>Close
                </button>
            </DialogActions>
        </Dialog>
    )}

    renderTotalSavings = () => {
        let amount = this.calculateTotalSavings()
        console.log(amount)
        return(
            <div>
                <div className='budget-detail-map-box'>
                    <h3>{`Savings`}</h3>
                    <div className='savings'>
                        <List sx={{ width: '100%' }}>
                            <div className='budgets-list-item-text-container' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={capitalizeFirstLowercaseRest(`Total Incomes`)}
                                />
                                <ListItemText
                                    style={{textAlign: 'right'}}
                                    primary={'$ ' + amount[0]}
                                />
                            </div>
                            <Divider variant='middle'/>
                            <div className='budgets-list-item-text-container' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={capitalizeFirstLowercaseRest(`Total Expenses`)}
                                />
                                <ListItemText
                                    style={{textAlign: 'right'}}
                                    primary={'$ ' + amount[1]}
                                />
                            </div>
                            <div className='budgets-list-item-text-container' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={capitalizeFirstLowercaseRest(`Total Goals`)}
                                />
                                <ListItemText
                                    style={{textAlign: 'right'}}
                                    primary={'$ ' + amount[2]}
                                />
                            </div>
                            <Divider variant='fullWidth'/>
                            <div className='budgets-list-item-text-container' style={{display: 'flex', marginTop: '0px', paddingTop: '8px', paddingBottom: '8px'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={capitalizeFirstLowercaseRest(`Total Savings`)}
                                />
                                <ListItemText
                                    style={{textAlign: 'right'}}
                                    primary={(amount[3] > 0 ? ' ' : '-')+ '$ ' + Math.abs(amount[3])}
                                />
                            </div>
                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                <Box sx={{ width: '100%', mr: 1 }}>
                                    <LinearProgress variant="determinate" color="primary" value={((amount[3] / this.props.budget.savings) * 100) > 0 ? ((amount[3] / this.props.budget.savings) * 100) : 0}/>
                                </Box>
                                <Box sx={{ minWidth: 35 }}>
                                    <Typography variant="body2" color="text.secondary">{`${(amount[3] / this.props.budget.savings) * 100}%`}</Typography>
                                </Box>
                            </Box>
                        </List>
                    </div>
                <br></br>
                </div>
                <br></br>
            </div>
        )
    }

    renderTransactionDetails = (title, transactionMap) =>
    {
        return (
            <div>
                <div className='budget-detail-map-box'>
                {/* <Divider variant='middle' /> */}
                    <h3>{title}</h3>
                    <div className='transactions'>
                        <List sx={{ width: '100%' }}>
                            {transactionMap.length ? <p></p> : <p>Oops! No Transactions entered</p>}
                            {transactionMap.map(data => (
                                <div>
                                    <div className='transaction-detail' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                        <ListItemText
                                            style={{width: '80%'}}
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
                                            style={{textAlign: 'right'}}
                                            primary={(data.type ? ' ' : '-')+ '$ ' + data.amount}
                                            secondary={data.date.substring(0, 10)}
                                            />
                                        </div>
                                        {/* <Box
                                            display={'flex'}
                                            justifyContent={'flex-end'}
                                            marginRight='5%'
                                            >
                                        </Box> */}
                                        <Divider variant='fullWidth' component='li' />
                                </div>
                            ))}
                        </List>
                    </div>
                <br></br>
                </div>
                <br></br>
            </div>
        )
    }

}

export default BudgetDetail;