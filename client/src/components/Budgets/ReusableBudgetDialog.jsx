import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import BudgetMapInput from './BudgetMapInput';
import { CATEGORY_MAP, FREQUENCY_MAP } from '../../utils/GlobalConstants';
import axios from 'axios';
import { EXPENSES, GOALS, INCOMES } from './BudgetConstants';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';

class ReusableBudgetDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {allGoals:[]};

        console.log("props for budget dialog", this.props);
        console.log("states for budget dialog", this.state);
        this.getAllGoals()
    };

    // getGoalIds

    async getAllGoals() {
        let res = await axios.get('/api/goals');
        console.log('get all goals: ', res.data.body)
        if (res.data.body != null)
        {
            this.setState({
                allGoals: res.data.body.map(obj => {
                    return {
                        'id': obj.goal_id,
                        'value': obj.goal_name
                    }
                }),
            });
        } else {
            this.setState({
                allGoals: [],
            });
        }

    }

    handleIncomeChange = (event, amountMap) => {
        console.log("show reusable-handleIncomeChange amountMap: ", amountMap)
        this.props.handleInputChange(event, {type: INCOMES, data: amountMap});
    }

    handleExpenseChange = (event, amountMap) => {
        console.log("show reusable-handleExpenseChange amountMap: ", amountMap)
        this.props.handleInputChange(event, {type: EXPENSES, data: amountMap});

    }

    handleGoalChange = (event, amountMap) => {
        console.log("show reusable-handleGoalChange amountMap: ", amountMap)
        this.props.handleInputChange(event, {type: GOALS, data: amountMap});

    }

    render() {
        return(
            <Dialog
                className='budget-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <DialogTitle textAlign={'center'}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <div className='budget-input-group'>
                        
                        <label htmlFor='budget_name'>Name</label><br></br>
                        <input
                            defaultValue={this.props.currentBudget.budget_name}
                            type='text'
                            name='budget_name'
                            className='budget-input-box'
                            placeholder='Name'
                            onChange={this.props.handleInputChange}
                            />
                    </div>
                    <br></br>
                    <div className='budget-input-group'>
                        <BudgetMapInput
                            name={INCOMES}
                            optionsList={CATEGORY_MAP}
                            handleChange={this.handleIncomeChange}
                            currentDataMap={this.props.currentBudget.income_map ? this.props.currentBudget.income_map : []}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <BudgetMapInput 
                            name={EXPENSES}
                            optionsList={CATEGORY_MAP}
                            handleChange={this.handleExpenseChange}
                            currentDataMap={this.props.currentBudget.expense_map ? this.props.currentBudget.expense_map : []}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <BudgetMapInput 
                            name={GOALS}
                            optionsList={this.state.allGoals}
                            handleChange={this.handleGoalChange}
                            currentDataMap={this.props.currentBudget.goal_map ? this.props.currentBudget.goal_map : []}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <label htmlFor='date'>Date</label><br></br>
                        <input
                            defaultValue={transformDateFormatToYyyyMmDd(this.props.currentBudget.creation_date)}
                            type='date'
                            name='creation_time'
                            className='budget-input-box'
                            onChange={this.props.handleInputChange}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <label htmlFor='frequency'>Frequency</label><br></br>
                        <select
                            name='frequency'
                            className='budget-input-box'
                            defaultValue={this.props.currentBudget.frequency}
                            onChange={this.props.handleInputChange}
                        >
                            {FREQUENCY_MAP.map(freq => (
                                <option value={freq.id}>{freq.value}</option>
                            ))} 
                        </select>
                    </div>
                </DialogContent>
                <DialogActions>
                    <button
                        type='submit'
                        className='budget-submit-btn'
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type='submit'
                        className='close-budget-submit-btn'
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableBudgetDialog;