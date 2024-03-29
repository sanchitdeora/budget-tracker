import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import { FormControl, FormGroup, InputAdornment, InputLabel, MenuItem, Select, TextField } from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import dayjs from 'dayjs';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';
import { CATEGORY_MAP, FREQUENCY_MAP } from '../../utils/GlobalConstants';

class ReusableBillDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            billCategory: Object.keys(props.currentBill).length > 0 ? props.currentBill.category : '',
        };

        console.log("props for reusable bill", this.props);
    };

    getBillCategory = () => {
        let category = Object.keys(this.props.currentBill).length > 0 ? this.props.currentBill.category : '';
        console.log("For Edit bill category: ", category);
        
        return category;
    }

    getBillFrequency = () => {
        let frequency = Object.keys(this.props.currentBill).length > 0 ? this.props.currentBill.frequency : '';
        console.log("For Edit bill frequency: ", frequency);
        
        return frequency;
    }

    render() {
        return(
            <Dialog
                className='bill-dialog'
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
                            defaultValue={this.props.currentBill.title}
                            name='title'
                            label='Title'
                            // color='primary'
                            onChange={this.props.handleChange}
                            variant="outlined"
                        />
                    </FormControl>
                    <br></br>
                    <FormControl className='input-group' sx={{ width: 300 }}>
                        <InputLabel id="input-label">Category</InputLabel>
                        <Select
                            labelId="input-label"
                            defaultValue={this.getBillCategory()}
                            label="Category"
                            onChange={this.props.handleChange}
                            variant="outlined" 
                        >
                            {CATEGORY_MAP.map(category => (
                                <MenuItem key={category.id} value={category.id}> {category.value} </MenuItem>
                            ))} 
                        </Select>
                    </FormControl>
                    <br></br>
                    <FormControl className='input-group' sx={{ width: 300 }}>
                        <TextField 
                            defaultValue={this.props.currentBill.amount_due}
                            name='amount_due'
                            type='number'
                            label='Amount Due'
                            InputProps={{
                                startAdornment: 
                                <InputAdornment disableTypography position="start">
                                    $</InputAdornment>,
                                inputMode: 'numeric', pattern: '[0-9]*' 
                            }}
                            onChange={this.props.handleChange}
                            variant="outlined" 
                        />
                    </FormControl>
                    <br></br>
                    <div className='input-group'>
                        <DatePicker label="Due Date"
                            defaultValue={dayjs(transformDateFormatToYyyyMmDd(this.props.currentBill.due_date))}
                            className='input-date'
                            name='due_date'
                            onChange={this.props.handleChange}
                            />
                    </div>
                    <br></br>
                    <FormControl className='input-group' sx={{ width: 300 }}>
                        <InputLabel id="input-label">Frequency</InputLabel>
                        <Select
                            labelId="input-label"
                            defaultValue={this.getBillFrequency()}
                            label="Frequency"
                            onChange={this.props.handleChange}
                            variant="outlined" 
                        >
                            {FREQUENCY_MAP.map(frequency => (
                                <MenuItem key={frequency.id} value={frequency.id}> {frequency.value} </MenuItem>
                            ))} 
                        </Select>
                    </FormControl>
                    <br></br>
                    <FormControl className='input-group ' sx={{ width: 300 }}>
                        <TextField 
                            defaultValue={this.props.currentBill.note}
                            name='note'
                            label='Note'
                            multiline
                            maxRows={2}
                            // className='transaction-input-box'
                            onChange={this.props.handleChange}
                            variant="outlined" 
                        />
                    </FormControl>
                    <br></br>
                </FormGroup>
                </DialogContent>
                <DialogActions>
                    <button
                        type='submit'
                        className='bill-submit-btn'
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type='submit'
                        className='close-bill-submit-btn'
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableBillDialog;