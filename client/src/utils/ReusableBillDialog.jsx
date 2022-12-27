import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'

class ReusableBillDialog extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};
	};

    render() {
        return(
            <Dialog
                className='bill-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <DialogTitle textAlign={'center'}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <div className='bill-input-group'>
                        
                        <label htmlFor='title'>Title</label><br></br>
                        <input
                            type='text'
                            name='title'
                            className='bill-input-box'
                            placeholder='Title'
                            onChange={this.props.handleChange}
                            />
                    </div>
                    <br></br>
                    <div className='bill-input-group'>
                        <label htmlFor='category'>Category</label><br></br>
                        <select 
                            name='category'
                            className='bill-input-box'
                            defaultValue={'DEFAULT'}
                            onChange={this.props.handleChange}
                        >
                            <option value='DEFAULT' disabled>None</option>
                            <option value='bills_and_utilities'>Bills & Utilities</option>
                            <option value='education'>Education</option>
                            <option value='entertainment'>Entertainment</option>
                            <option value='loans'>Loans</option>
                            <option value='Medical'>Medical</option>
                            <option value='Rent'>Rent</option>
                            <option value='uncategorized'>Others</option>
                        </select>
                    </div>
                    <br></br>
                    <div className='bill-input-group'>
                        <label htmlFor='amount'>Amount Due</label><br></br>
                        <input
                            type='number'
                            name='amount_due'
                            className='bill-input-box'
                            placeholder='Amount Due'
                            onChange={this.props.handleChange}
                            />
                    </div>
                    <br></br>
                    <div className='bill-input-group'>
                        <label htmlFor='date'>Date Due</label><br></br>
                        <input
                            type='datetime'
                            name='due_date'
                            className='bill-input-box'
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className='bill-input-group'>
                        <label htmlFor='frequency'>Frequency</label><br></br>
                        <select 
                            name='frequency'
                            className='bill-input-box'
                            defaultValue={'once'}
                            onChange={this.props.handleChange}
                        >
                            <option value='once'>Once</option>
                            <option value='weekly'>Weekly</option>
                            <option value='bi-weekly'>Every Two Weeks</option>
                            <option value='monthly'>Monthly</option>
                            <option value='bi_monthly'>Every Two Months</option>
                            <option value='quaterly'>Quaterly</option>
                            <option value='half_yearly'>Every Six Months</option>
                            <option value='yearly'>Yearly</option>
                        </select>
                    </div>
                    <br></br>
                    <div className='bill-input-group'>
                        <label htmlFor='Note'>Note</label><br></br>
                        <textarea
                            name='note'
                            className='bill-input-box'
                            placeholder='Add notes here'
                            onChange={this.props.handleChange}
                            />
                    </div>
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