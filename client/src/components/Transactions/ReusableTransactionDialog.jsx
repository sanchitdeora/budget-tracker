import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import { CATEGORY_MAP } from '../../utils/GlobalConstants';
import { transformDateFormatToYyyyMmDd } from '../../utils/StringUtils';
import { FormControl, FormControlLabel, FormGroup, FormLabel, InputAdornment, InputLabel, MenuItem, OutlinedInput, Radio, RadioGroup, Select, TextField } from '@mui/material';

class ReusableTransactionDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            transactionType: props.currentTransaction !== undefined && props.currentTransaction.type !== undefined ? props.currentTransaction.type : false,
        };

        console.log("heereee: ", props.currentTransaction !== undefined && props.currentTransaction.type !== undefined ? props.currentTransaction.type : false)
        console.log("heereee1: ", props.currentTransaction)
    };

    handleChangeCurrentTransaction = (event) => {
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
                <DialogTitle textAlign={'center'}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <FormGroup>
                        <br></br>
                        <FormControl className='transaction-input-group' sx={{ width: 300 }}>
                            <TextField 
                                defaultValue={this.props.currentTransaction.title}
                                name='title'
                                label='Title'
                                // color='primary'
                                className='transaction-input-box'
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
                                className='transaction-input-box'
                                defaultValue={this.props.currentTransaction.category}
                                // value={this.props.currentTransaction.category}
                                label="Category"
                                onChange={this.props.handleChange}
                                variant="outlined" 
                            >
                                {CATEGORY_MAP.map(category => (
                                    <MenuItem value={category.id}> {category.value} </MenuItem>
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
                                className='transaction-input-box'
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
                                name="row-radio-buttons-group"
                                // className='transaction-input-box'
                                variant="outlined" 
                                />
                                <FormControlLabel value="debit" label="Debit" control={<Radio  
                                    checked={this.state.transactionType === false} 
                                    onChange={this.handleChangeCurrentTransaction}
                                    />} />
                                <FormControlLabel value="credit" label="Credit" control={<Radio  
                                    checked={this.state.transactionType === true}
                                    onChange={this.handleChangeCurrentTransaction}
                                />} />
                        </FormControl>
                        <br></br>
                        <div className='transaction-input-group'>
                            <InputLabel id="transaction-input-label">Date</InputLabel>
                            <input
                                defaultValue={transformDateFormatToYyyyMmDd(this.props.currentTransaction.date)}
                                type='date'
                                name='date'
                                className='transaction-input-box'
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
                                className='transaction-input-box'
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
                        className='transaction-submit-btn'
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type='submit'
                        className='close-transaction-submit-btn'
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableTransactionDialog;