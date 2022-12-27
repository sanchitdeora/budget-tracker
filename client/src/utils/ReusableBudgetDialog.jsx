import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import BudgetMapInput from './BudgetMapInput';

class ReusableBudgetDialog extends React.Component {
	constructor(props) {
		super(props);
		this.state = {};
	};

    handleIncomeChange = (event) => {
        if (!event.target.name.startsWith("Income")) {
            event.target.name = "Income-" + event.target.name;
        }
        this.props.handleChange(event);
    }

    handleSpendingChange = (event) => {
        if (!event.target.name.startsWith("Spending")) {
            event.target.name = "Spending-" + event.target.name;
        }
        this.props.handleChange(event);
    }

    handleRemoveFromIncomeMap = (item) => {
        item.value = "Income-" + item.value;
        this.props.handleRemoveFromMap(item);
    }

    handleRemoveFromSpendingMap = (item) => {
        item.value = "Spending-" + item.value;
        this.props.handleRemoveFromMap(item);
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
                        
                        <label htmlFor='name'>Name</label><br></br>
                        <input
                            type='text'
                            name='name'
                            className='budget-input-box'
                            placeholder='Name'
                            onChange={this.props.handleChange}
                            />
                    </div>
                    <br></br>
                    <div className='budget-input-group'>
                        <BudgetMapInput
                            name="Income"
                            handleChange={this.handleIncomeChange}
                            handleRemoveFromMap={this.handleRemoveFromIncomeMap}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <BudgetMapInput 
                            name="Spendings"
                            handleChange={this.handleSpendingChange}
                            handleRemoveFromMap={this.handleRemoveFromSpendingMap}
                        />
                    </div>

                    <br></br>
                    <div className='budget-input-group'>
                        <label htmlFor='savings'>Savings</label><br></br>
                        <input
                            type='number'
                            name='savings'
                            className='budget-input-box'
                            placeholder='Savings'
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className='budget-input-group'>
                        <label htmlFor='frequency'>Frequency</label><br></br>
                        <select 
                            name='frequency'
                            className='budget-input-box'
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