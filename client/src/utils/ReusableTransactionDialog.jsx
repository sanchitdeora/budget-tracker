import React from 'react';
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'

class ReusableTransactionDialog extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};
	};

    render() {
        return(
            <Dialog
                className='transaction-dialog'
                open={this.props.isDialogOpen}
                fullWidth={true}
                onClose={this.props.handleClose}
            >
                <DialogTitle textAlign={"center"}>{this.props.title}</DialogTitle>
                <DialogContent>
                    <div className="transaction-input-group">
                        
                        <label htmlFor="title">Title</label><br></br>
                        <input
                            type="text"
                            name="title"
                            className="transaction-input-box"
                            placeholder="Title"
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className="transaction-input-group">
                        <label htmlFor="category">Category</label><br></br>
                        <select 
                            name="category"
                            className="transaction-input-box"
                            defaultValue={"DEFAULT"}
                            onChange={this.props.handleChange}
                        >
                            <option value="DEFAULT" disabled>None</option>
                            <option value="auto_and_transport">Auto & Transport</option>
                            <option value="bills_and_utilities">Bills & Utilities</option>
                            <option value="education">Education</option>
                            <option value="entertainment">Entertainment</option>
                            <option value="food_and_dining">Food & Dining</option>
                            <option value="health_and_fitness">Health & Fitness</option>
                            <option value="home">Home</option>
                            <option value="income">Income</option>
                            <option value="investments">Investments</option>
                            <option value="personal_care">Personal Care</option>
                            <option value="pets">Pets</option>
                            <option value="shopping">Shopping</option>
                            <option value="taxes">Taxes</option>
                            <option value="travel">Travel</option>
                            <option value="uncategorized">Others</option>
                        </select>
                    </div>
                    <br></br>
                    <div className="transaction-input-group">
                        <label htmlFor="amount">Amount</label><br></br>
                        <input
                            type="number"
                            name="amount"
                            className="transaction-input-box"
                            placeholder="Amount"
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className="transaction-input-group">
                    <label htmlFor="Debit">
                        <input
                            type="radio"
                            name="type"
                            value="debit"
                            placeholder="Debit"
                            defaultChecked
                            onChange={this.props.handleChange}
                        /> Debit
                    </label>
                    <label htmlFor="Credit">
                        <input
                            type="radio"
                            name="type"
                            value="credit"
                            placeholder="Credit"
                            onChange={this.props.handleChange}
                            /> Credit
                    </label>
                    </div>
                    <br></br>
                    <div className="transaction-input-group">
                        <label htmlFor="date">Date</label><br></br>
                        <input
                            type="date"
                            name="date"
                            className="transaction-input-box"
                            onChange={this.props.handleChange}
                        />
                    </div>
                    <br></br>
                    <div className="transaction-input-group">
                        <label htmlFor="Note">Note</label><br></br>
                        <textarea
                            type="text"
                            name="note"
                            className="transaction-input-box"
                            placeholder="Note"
                            onChange={this.props.handleChange}
                        />
                    </div>
                </DialogContent>
                <DialogActions>
                    <button
                        type="submit"
                        className="transaction-submit-btn"
                        onClick={this.props.submitMethod}>Submit
                    </button>						
                    <button
                        type="submit"
                        className="close-transaction-submit-btn"
                        onClick={this.props.handleClose}>Close
                    </button>
                </DialogActions>
            </Dialog>
        )
    }
}

export default ReusableTransactionDialog;