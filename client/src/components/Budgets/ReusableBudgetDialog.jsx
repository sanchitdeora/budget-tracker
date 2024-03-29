import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import BudgetMapInput from './BudgetMapInput';
import { CATEGORY_MAP, FREQUENCY_MAP } from '../../utils/GlobalConstants';
import axios from 'axios';
import { FormControl, FormGroup, InputLabel, MenuItem, Select, TextField } from '@mui/material';
import dayjs from 'dayjs';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';
import '../../utils/Dialog.scss';
import { EXPENSES, GOALS, INCOMES } from './BudgetConstants';

class ReusableBudgetDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allGoals: []
        }
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
    
    getTransactionFrequency = () => {
        let frequency = Object.keys(this.props.currentBudget).length > 0 ? this.props.currentBudget.frequency : '';
        console.log("For Edit transaction: ", frequency);
        
        return frequency;
    }

    render() {
        return(
            <Dialog
                className='dialog-container'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <h3 className='header dialog-header'>{this.props.title}</h3>
                <DialogContent className='dialog-body'>
                    <FormGroup>
                        <br></br>
                        <FormControl className='input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentBudget.budget_name}
                                name='budget_name'
                                label='Name'
                                // color='primary'
                                onChange={this.props.handleInputChange}
                                variant="outlined"
                            />
                        </FormControl>
                        <br></br>
                        <div className='budget-input-map-group'>
                            <BudgetMapInput
                                name={INCOMES}
                                optionsList={CATEGORY_MAP}
                                handleChange={this.handleIncomeChange}
                                currentDataMap={this.props.currentBudget.income_map ? this.props.currentBudget.income_map : []}
                            />
                        </div>

                        <br></br>
                        <div className='budget-input-map-group'>
                            <BudgetMapInput 
                                name={EXPENSES}
                                optionsList={CATEGORY_MAP}
                                handleChange={this.handleExpenseChange}
                                currentDataMap={this.props.currentBudget.expense_map ? this.props.currentBudget.expense_map : []}
                            />
                        </div>

                        <br></br>
                        <div className='budget-input-map-group'>
                            <BudgetMapInput 
                                name={GOALS}
                                optionsList={this.state.allGoals}
                                handleChange={this.handleGoalChange}
                                currentDataMap={this.props.currentBudget.goal_map ? this.props.currentBudget.goal_map : []}
                            />
                        </div>
                        <br></br>
                        <div className='input-group'>
                            <DatePicker label="Date"
                                defaultValue={dayjs(transformDateFormatToYyyyMmDd(this.props.currentBudget.creation_time))}
                                className='input-date'
                                name='creation_time'
                                onChange={this.props.handleInputChange}
                                />
                        </div>
                        <br></br>
                        <FormControl className='input-group' sx={{ width: 300 }}>
                            <InputLabel id="input-label">Frequency</InputLabel>
                            <Select
                                name='frequency'
                                labelId="input-label"
                                defaultValue={this.getTransactionFrequency}
                                label="Frequency"
                                onChange={this.props.handleInputChange}
                                variant="outlined" 
                                >
                                {FREQUENCY_MAP.map(frequency => (
                                    <MenuItem key={frequency.id} value={frequency.id}> {frequency.value} </MenuItem>
                                ))} 
                            </Select>
                        </FormControl>
                    </FormGroup>
                </DialogContent>
                <DialogActions className='dialog-footer'>
                    <button
                        type='submit'
                        className='dialog-submit-btn'
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type='submit'
                        className='close-dialog-submit-btn'
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableBudgetDialog;