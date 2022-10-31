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
import "./Bills.scss";
import { IconButton } from '@mui/material';
import ReusableTransactionDialog from '../../utils/ReusableBillDialog';
import axios from 'axios';



class Bills extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			allBills: [],
			billId: "",
			title: "",
			category: "",
			amount_due: 0,
			date_due: new Date(),
			how_often: "",
			is_paid: false,
			note: "",
			isCreateDialogOpen: false,
			isEditDialogOpen: false,
		};
		console.log(this.state.allBills.length ? "true" : "false")
		this.getAllBills()
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
		
	// get bill

	async getAllBills() {
		let res = await axios.get("/api/bills");
		console.log("get all bills: ", res.data.body)
		if (res.data.body != null)
		{
			this.setState({
				allBills: res.data.body,
			});
		}
	}
	

	// create bill

	handleCreateBillOpen = () => {
		this.setState({
			isCreateDialogOpen: true
		});
	}

	submitCreateBill = () => {
		let date = new Date(this.state.date_due);
		let billBody = {
			"title": this.state.title,
			"category": this.state.category,
			"amount_due": parseFloat(this.state.amount_due),
			"date_due": date,
			"how_often": this.state.how_often,
			"is_paid": this.state.is_paid,
			"note": this.state.note,
		}
		console.log('The create form was submitted with the following data:', billBody,);
		this.postBillRequest(billBody)
		this.handleCreateClose()
	}

	async postBillRequest(billBody) {
		let res = await axios.post("/api/bill", billBody);
		console.log(res);
		this.getAllBills();
	}

	handleCreateClose = () => {
		this.setState({
			isCreateDialogOpen: false
		});
	};


	// edit bill

	handleEditBillOpen = (id) => {
		console.log("Edit id: ", id)
		this.setState({
			billId: id,
			isEditDialogOpen: true
		});
	}

	submitEditBill = () => {
		let date = new Date(this.state.date);
		let billBody = {
			"title": this.state.title,
			"category": this.state.category,
			"amount_due": parseFloat(this.state.amount_due),
			"date_due": date,
			"note": this.state.note,
		}
		console.log('The edit form was submitted with the following data:', billBody);
		this.putBillRequest(billBody)
		this.handleEditClose()
	}

	async putBillRequest(billBody) {
		let res = await axios.put("/api/bill/"+this.state.billId, billBody);
		console.log(res);
		this.getAllBills();
	}

	handleEditClose = () => {
		this.setState({
			isEditDialogOpen: false
		});
	};

	// delete bill

	handleDeleteBillOpen = (id) => {
		console.log("Delete id: ", id)
		this.deleteBillRequest(id)
	}

	async deleteBillRequest(id) {
		let res = await axios.delete("/api/bill/"+id);
		console.log(res);
		this.getAllBills();
	}


	render() {
		return (
			<div className='bills-inner-container'>
				<div className="header">
					Bills
				</div>
				<div className='bills-box'>
					<div className="bills-category-box">
						<List sx={{ width: '100%', bgcolor: 'background.paper' }}>
							{this.state.allBills.length ? <p></p> : <p>Oops! No Bills entered</p>}
							{this.state.allBills?.map(data => (
								<div className='bill'>
									<ListItem key={data.bill_id} id={data.bill_id} alignItems="flex-start">
										<ListItemText
											primary={this.capitalizeFirstLowercaseRest(data.title)}
											secondary={<React.Fragment>
												<Typography
													sx={{ display: 'inline' }}
													component="span"
													variant="body2"
													color="text.primary"
													>
													{"($"}{data.amount}{")"}
												</Typography>
												{" — "}{data.date}
												<br></br>
												{data.note}
											</React.Fragment>} />
									</ListItem>
									<IconButton edge="end" onClick={this.handleEditBillOpen.bind(this, data.bill_id)}>
											<ModeEditIcon />
									</IconButton>
									<ReusableTransactionDialog
										title={"Edit Bill"}
										isDialogOpen={this.state.isEditDialogOpen}
										handleChange={this.handleChange}
										handleClose={this.handleEditClose}
										submitMethod={this.submitEditBill}
									/>
									<IconButton edge="end"onClick={this.handleDeleteBillOpen.bind(this, data.bill_id)}>
											<DeleteIcon />
									</IconButton>
									<Divider variant="middle" component="li" />
								</div>
							))}
						</List>
					</div>
				</div>
				<div className='create-button'>
					<IconButton size='large' onClick={this.handleCreateBillOpen}>
						<AddCircleIcon />
					</IconButton>
				</div>
				
				<ReusableTransactionDialog
					title={"Add Bill"}
					isDialogOpen={this.state.isCreateDialogOpen}
					handleChange={this.handleChange}
					handleClose={this.handleCreateClose}
					submitMethod={this.submitCreateBill}
				/>
			</div>
		);
	}
}

export default Bills;