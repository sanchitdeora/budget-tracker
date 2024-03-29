import React from 'react';
import { FormControl, TextField, DialogActions, DialogContent, Dialog, InputAdornment, FormGroup } from '@mui/material';
import GoalsBudgetSelect from './GoalsBudgetSelect'
import dayjs from 'dayjs';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';

class ReusableGoalDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};

        console.log("props for goal dialog", this.props);
        console.log("states for goal dialog", this.state);
    };

    // render functions

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
                        <FormControl className='goal-input-group' sx={{ width: 300 }}>
                            <br></br>
                            <TextField 
                                defaultValue={this.props.currentGoal.goal_name}
                                name='goal_name'
                                label="Name"
                                color='primary'
                                className='goal-input-box'
                                onChange={this.props.handleChange}
                                variant="outlined" />
                            <br></br>
                        </FormControl>
                        {
                            this.props.title.includes("Edit") 
                            ?
                            <FormControl className='input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentGoal.current_amount}
                                name='current_amount'
                                InputProps={{
                                    startAdornment: 
                                    <InputAdornment disableTypography position="start">
                                        $</InputAdornment>,
                                    inputMode: 'numeric', pattern: '[0-9]*' 
                                }}
                                label="Current Amount"
                                onChange={this.props.handleChange}
                                variant="outlined" />
                                <br></br>
                            </FormControl>
                            :
                            ""
                        }
                        <FormControl className='input-group' sx={{ width: 300 }}>
                            <TextField 
                            defaultValue={this.props.currentGoal.target_amount}
                            name='target_amount'
                            type='number'
                            InputProps={{
                                startAdornment: 
                                <InputAdornment disableTypography position="start">
                                    $</InputAdornment>,
                                inputMode: 'numeric', pattern: '[0-9]*' 
                            }}
                            label="Target Amount"
                            onChange={this.props.handleChange}
                            variant="outlined" />
                        <br></br>
                        </FormControl>
                        <div className='input-group'>
                            <DatePicker label="Target Date"
                                defaultValue={dayjs(transformDateFormatToYyyyMmDd(this.props.currentGoal.target_date))}
                                name='target_date'
                                id="target_date"
                                className='input-date'
                                onChange={this.props.handleChange}
                                />
                        </div>
                        <br></br>
                        <div>
                            <GoalsBudgetSelect handleBudgetIds={this.props.handleBudgetIds} currentGoal={this.props.currentGoal} allBudgets={this.props.allBudgets} />
                        </div>
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

export default ReusableGoalDialog;