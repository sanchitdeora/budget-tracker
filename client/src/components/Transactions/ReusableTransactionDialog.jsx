import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import { CATEGORY_MAP } from '../../utils/GlobalConstants';
import { FormControl, FormControlLabel, FormGroup, FormLabel, InputAdornment, InputLabel, MenuItem, Radio, RadioGroup, Select, TextField } from '@mui/material';
import dayjs from 'dayjs';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';
import '../../utils/Dialog.scss';

class ReusableTransactionDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            transactionType: props.currentTransaction !== undefined && props.currentTransaction.type !== undefined ? props.currentTransaction.type : false,
            transactionCategory: Object.keys(props.currentTransaction).length > 0 ? props.currentTransaction.category : '',
        };

        console.log("heereee: ", props.currentTransaction !== undefined && props.currentTransaction.type !== undefined ? props.currentTransaction.type : false)
        console.log("heereee1: ", props.currentTransaction)
    };

    handleChangeTransactionType = (event) => {
        event.target.name = "type"
        console.log('Name: ' + event.target.name + ' value: ' + event.target.value)
        let val = event.target.value === 'credit' ? true : false 
        this.setState({transactionType: val});
        this.props.handleChange(event)
    }

    render() {
        return(
            <Dialog
                className='transaction-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                hideBackdrop={false}
                onClose={this.props.handleClose}
            >
                <h3 className='header dialog-header'>{this.props.title}</h3>
                <DialogContent className='dialog-body'>
                    <FormGroup>
                        <br></br>
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentTransaction.title}
                                name='title'
                                label='Title'
                                // color='primary'
                                onChange={this.props.handleChange}
                                variant="outlined"
                            />
                        </FormControl>
                        <br></br>
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <InputLabel id="transaction-input-label">Category</InputLabel>
                            <Select
                                labelId="transaction-input-label"
                                id="demo-simple-select"
                                defaultValue={this.state.transactionCategory}
                                // value={this.props.currentTransaction.category}
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
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentTransaction.amount}
                                name='amount'
                                type='number'
                                label='Amount'
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
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <FormLabel id="transaction-input-label">Transaction Type</FormLabel>
                            <RadioGroup
                                row
                                defaultValue={this.props.currentTransaction.type}
                                name="type"
                                variant="outlined" 
                                />
                                <FormControlLabel id="transaction-input-label" value="debit" label="Debit" control={<Radio  
                                    checked={this.state.transactionType === false} 
                                    onChange={this.handleChangeTransactionType}
                                    />} />
                                <FormControlLabel id="transaction-input-label" value="credit" label="Credit" control={<Radio  
                                    checked={this.state.transactionType === true}
                                    onChange={this.handleChangeTransactionType}
                                />} />
                        </FormControl>
                        <br></br>
                        <div className='transaction-input-group'>
                            <DatePicker label="Date"
                                defaultValue={dayjs(transformDateFormatToYyyyMmDd(this.props.currentTransaction.date))}
                                orientation='landscape'
                                className='transaction-input-date'
                                name='date'
                                onChange={this.props.handleChange}
                                />
                        </div>
                        <br></br>
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentTransaction.note}
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

export default ReusableTransactionDialog;