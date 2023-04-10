import React from 'react';
import { FormControl, TextField, DialogActions, DialogContent, DialogTitle, Dialog, InputAdornment, FormGroup } from '@mui/material';
import GoalsBudgetSelect from './GoalsBudgetSelect'

class ReusableGoalDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    };

    handleChange = (event) => {
        const {
          target: { value },
        } = event;
        this.setState(
          typeof value === 'string' ? value.split(',') : value,
        );
      };

    

    render() {
        return(
            <Dialog
                className='goal-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <DialogTitle textAlign={'center'}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <FormGroup>
                        <FormControl className='goal-input-group' sx={{ width: 300 }}>
                            <br></br>
                            <TextField 
                                name='name'
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
                            <FormControl className='goal-input-group' sx={{ width: 300 }}>
                                <TextField 
                                name='current_amount'
                                InputProps={{
                                    startAdornment: 
                                    <InputAdornment disableTypography position="start">
                                        $</InputAdornment>,
                                    inputMode: 'numeric', pattern: '[0-9]*' 
                                }}
                                label="Current Amount"
                                className='goal-input-box'
                                onChange={this.props.handleChange}
                                variant="outlined" />
                                <br></br>
                            </FormControl>
                            :
                            ""
                        }
                        <FormControl className='goal-input-group' sx={{ width: 300 }}>
                            <TextField 
                            name='target_amount'
                            type='number'
                            InputProps={{
                                startAdornment: 
                                <InputAdornment disableTypography position="start">
                                    $</InputAdornment>,
                                inputMode: 'numeric', pattern: '[0-9]*' 
                            }}
                            label="Target Amount"
                            className='goal-input-box'
                            onChange={this.props.handleChange}
                            variant="outlined" />
                        <br></br>
                        </FormControl>
                        <FormControl className='goal-input-group' sx={{ width: 300 }}>
                            <TextField
                                id="date"
                                label="Target Date"
                                type="date"
                                className='goal-input-box'
                                onChange={this.props.handleChange}
                                InputLabelProps={{
                                shrink: true,
                                }}
                            />
                        <br></br>
                        </FormControl>
                        <FormControl className='goal-input-group' sx={{ width: 300 }}>
                            <GoalsBudgetSelect handleBudgetIds={this.props.handleBudgetIds} />
                        </FormControl>
                    </FormGroup>
                </DialogContent>
                <DialogActions>
                    <button
                        type='submit'
                        className='goal-submit-btn'
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type='submit'
                        className='close-goal-submit-btn'
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableGoalDialog;