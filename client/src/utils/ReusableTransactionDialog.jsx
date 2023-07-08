import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import { CATEGORY_MAP } from './GlobalConstants';
import { transformDateFormatToYyyyMmDd } from './StringUtils';

class ReusableTransactionDialog extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
        };

        console.log(props);
    };

    handleChangeCurrentTransaction() {

    }

    render() {
        return(
            <Dialog
                className='transaction-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <DialogTitle textAlign={'center'}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <div className='transaction-input-group'>
                        
                        <label htmlFor='title'>Title</label><br></br>
                        <input
                            defaultValue={this.props.currentTransaction.title}
                            type='text'
                            name='title'
                            className='transaction-input-box'
                            placeholder='Title'
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className='transaction-input-group'>
                        <label htmlFor='category'>Category</label><br></br>
                        <select 
                            name='category'
                            className='transaction-input-box'
                            defaultValue={this.props.currentTransaction.category}
                            onChange={this.props.handleChange}
                        >
                            {CATEGORY_MAP.map(freq => (
                                <option value={freq.id}>{freq.value}</option>
                            ))} 
                        </select>
                    </div>
                    <br></br>
                    <div className='transaction-input-group'>
                        <label htmlFor='amount'>Amount</label><br></br>
                        <input
                            defaultValue={this.props.currentTransaction.amount}
                            type='number'
                            name='amount'
                            className='transaction-input-box'
                            placeholder='Amount'
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className='transaction-input-group'>
                    <label htmlFor='Debit'>
                        <input
                            type='radio'
                            name='type'
                            value='debit'
                            placeholder='Debit'
                            checked={this.props.currentTransaction.type === false}
                            onChange={this.props.handleChange}
                        /> Debit
                    </label>
                    <label htmlFor='Credit'>
                        <input
                            type='radio'
                            name='type'
                            value='credit'
                            placeholder='Credit'
                            checked={this.props.currentTransaction.type === true}
                            onChange={this.props.handleChange}
                            /> Credit
                    </label>
                    </div>
                    <br></br>
                    <div className='transaction-input-group'>
                        <label htmlFor='date'>Date</label><br></br>
                        <input
                            defaultValue={transformDateFormatToYyyyMmDd(this.props.currentTransaction.date)}
                            type='date'
                            name='date'
                            className='transaction-input-box'
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className='transaction-input-group'>
                        <label htmlFor='Note'>Note</label><br></br>
                        <textarea
                            defaultValue={this.props.currentTransaction.note}
                            type='text'
                            name='note'
                            className='transaction-input-box'
                            placeholder='Note'
                            onChange={this.props.handleChange}
                        />
                    </div>
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