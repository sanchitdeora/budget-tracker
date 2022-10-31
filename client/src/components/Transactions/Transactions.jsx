import React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import "./Transactions.scss";
import { IconButton } from '@mui/material';
import ReusableTransactionDialog from '../../utils/ReusableTransactionDialog';
import axios from 'axios';


class Transactions extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			allTransactions: [],
			transactionId: "",
			title: "",
			category: "",
			amount: 0,
			date: new Date(),
			note: "",
			isCreateDialogOpen: false,
			isEditDialogOpen: false,
		};
		this.getAllTransactions()
	};

	capitalizeFirstLowercaseRest = (str) => {
		var splitStr = str.toLowerCase().split(' ');
   		for (var i = 0; i < splitStr.length; i++) {
			splitStr[i] = splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);     
		}
		return splitStr.join(' '); 
	};
	
	handleChange = (event) => {
        let value = event.target.value;
		let name = event.target.name;
		this.setState({
			[name]: value,
		});
	}
		
	// get transaction

	async getAllTransactions() {
		let res = await axios.get("/api/transactions");
		console.log("get all transactions: ", res.data.body)
		if (res.data.body != null)
		{
			this.setState({
				allBills: res.data.body,
			});
		}
	}
	

	// create transaction

	handleCreateTransactionOpen = () => {
		this.setState({
			isCreateDialogOpen: true
		});
	}

	submitCreateTransaction = () => {
		let date = new Date(this.state.date);
		let transactionBody = {
			"title": this.state.title,
			"category": this.state.category,
			"amount": parseFloat(this.state.amount),
			"date": date,
			"note": this.state.note,
		}
		console.log('The create form was submitted with the following data:', transactionBody,);
		this.postTransactionRequest(transactionBody)
		this.handleCreateClose()
	}

	async postTransactionRequest(transactionBody) {
		let res = await axios.post("/api/transaction", transactionBody);
		console.log(res);
		this.getAllTransactions();
	}

	handleCreateClose = () => {
		this.setState({
			title: "",
			category: "",
			amount: 0,
			date: "",
			note: "",
			isCreateDialogOpen: false
		});
	};


	// edit transaction

	handleEditTransactionOpen = (id) => {
		console.log("Edit id: ", id)
		this.setState({
			transactionId: id,
			isEditDialogOpen: true
		});
	}

	submitEditTransaction = () => {
		let date = new Date(this.state.date);
		let transactionBody = {
			"title": this.state.title,
			"category": this.state.category,
			"amount": parseFloat(this.state.amount),
			"date": date,
			"note": this.state.note,
		}
		console.log('The edit form was submitted with the following data:', transactionBody);
		this.putTransactionRequest(transactionBody)
		this.handleEditClose()
	}

	async putTransactionRequest(transactionBody) {
		let res = await axios.put("/api/transaction/"+this.state.transactionId, transactionBody);
		console.log(res);
		this.getAllTransactions();
	}

	handleEditClose = () => {
		this.setState({
			title: "",
			category: "",
			amount: 0,
			date: "",
			note: "",
			isEditDialogOpen: false
		});
	};

	// delete transaction

	handleDeleteTransactionOpen = (id) => {
		console.log("Delete id: ", id)
		this.deleteTransactionRequest(id)
	}

	async deleteTransactionRequest(id) {
		let res = await axios.delete("/api/transaction/"+id);
		console.log(res);
		this.getAllTransactions();
	}


	render() {
		return (
			<div className='transactions-inner-container'>
				<div className="header">
					Transactions
				</div>
				<div className='transactions-box'>
					<div className="transaction-category-box">
						<List sx={{ width: '100%', bgcolor: 'background.paper' }}>
							{this.state.allTransactions.length ? <p></p> : <p>Oops! No Transactions entered</p>}
							{this.state.allTransactions?.map(data => (
								<div className='transaction'>
									<ListItem key={data.transaction_id} id={data.transaction_id} alignItems="flex-start">
										<ListItemText
											primary={this.capitalizeFirstLowercaseRest(data.title)}
											secondary={<React.Fragment>
												<Typography
													sx={{ display: 'inline' }}
													component="span"
													variant="body2"
													color="text.primary"
												>
													{"$"}{data.amount}
												</Typography>
												{" — "}{data.date}
												<br></br>
												<i>{data.note}</i>
											</React.Fragment>}
										/>
									</ListItem>
									<IconButton onClick={this.handleEditTransactionOpen.bind(this, data.transaction_id)}>
											<ModeEditIcon />
									</IconButton>
									<ReusableTransactionDialog
										title={"Edit Transaction"}
										isDialogOpen={this.state.isEditDialogOpen}
										handleChange={this.handleChange}
										handleClose={this.handleEditClose}
										submitMethod={this.submitEditTransaction}
									/>
									<IconButton onClick={this.handleDeleteTransactionOpen.bind(this, data.transaction_id)}>
											<DeleteIcon />
									</IconButton>
									<Divider variant="middle" component="li" />
								</div>
							))}
						</List>
					</div>
				</div>
				<div className='create-button'>
					<IconButton size='large' onClick={this.handleCreateTransactionOpen}>
						<AddCircleIcon />
					</IconButton>
				</div>
				
				<ReusableTransactionDialog
					title={"Add Transaction"}
					isDialogOpen={this.state.isCreateDialogOpen}
					handleChange={this.handleChange}
					handleClose={this.handleCreateClose}
					submitMethod={this.submitCreateTransaction}
				/>
			</div>
		);
	}
}

export default Transactions;