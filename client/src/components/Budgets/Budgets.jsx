/* eslint-disable array-callback-return */
import React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import LinearProgress from '@mui/material/LinearProgress';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import { transformSnakeCaseToText } from '../../utils/StringUtils';
import './Budgets.scss';
import { IconButton } from '@mui/material';
import ReusableBudgetDialog from '../../utils/ReusableBudgetDialog';
import axios from 'axios';

class Budgets extends React.Component {
	constructor(props) {
		super(props);
        this.INCOME = "Income";
        this.SPENDING = "Spending"
		this.state = {
			budgetTest: [{
				budget_id: 'BG_1234',
				name: 'Monthly Expenditure',
				income_map: {
					"salary": 3000,
				},
				spending_limit_map: {
					"bills_and_utilities": 1000,
					"taxes": 500,
					"food_and_dining": 500,
					"uncategorized": 500
				},
				goal_amount_map: {},
				frequency: 'monthly',
				savings: 500
			}],

			allBudgets: [],
			budget_id: '',
			name: '',
			income_map: {},
			spending_limit_map: {},
			goal_amount_map: {},
			frequency: '',
			savings: 0,
			isCreateDialogOpen: false,
			isEditDialogOpen: false,
		};
		console.log(this.state.allBudgets.length ? 'true' : 'false')
		this.getAllBudgets()
	};

	cleanBudgetState = () => {
		this.setState({
			budget_id: '',
			name: '',
			income_map: {},
			spending_limit_map: {},
			goal_amount_map: {},
			frequency: '',
			savings: 0,
		})
	}


	handleChange = (event) => {
        let value = event.target.value;
		let name = event.target.name;
        
        // for handling income map
        if (name.startsWith(this.INCOME)) {
            let map=this.state.income_map;
            if(name === this.INCOME) {
                map[value] = 0;
            }
            else if(name.startsWith(this.INCOME + "-")) {
                map[name.substring(this.INCOME.length + 1)] = parseFloat(value)
            }
            name="income_map"
            value = map
        }

        // for handling spending map
        if (name.startsWith(this.SPENDING)) {
            let map=this.state.spending_limit_map;
            if(name === this.SPENDING) {
                map[value] = 0;
            }
            else if(name.startsWith(this.SPENDING + "-")) {
                
                map[name.substring(this.SPENDING.length + 1)] = parseFloat(value)
            }
            name="spending_limit_map"
            value = map
        }

		this.setState({
			[name]: value,
		});
	}

    handleRemoveFromMap = (item) => {
        let name;
        let newMap;
        if(item.value.startsWith(this.INCOME + "-")) {
            name="income_map"
            item.value = item.value.substring(this.INCOME.length + 1);
            newMap = this.state.income_map

        }
        if(item.value.startsWith(this.SPENDING + "-")) {
            name="spending_limit_map"
            item.value = item.value.substring(this.SPENDING.length + 1);
            newMap = this.state.spending_limit_map
        }


        Object.keys(newMap).map(element => {
            if(element === item.value)
                delete newMap[element]
        });

        this.setState({
            [name]: newMap
        })
    }

	// get budget

	async getAllBudgets() {
		let res = await axios.get('/api/budgets');
		console.log('get all budgets: ', res.data.body)
		if (res.data.body != null)
		{
			this.setState({
				allBudgets: res.data.body,
			});
		} else {
			this.setState({
				allBudgets: [],
			});
		}
	}

	// create budget

	handleCreateBudgetOpen = () => {
		this.setState({
			isCreateDialogOpen: true
		});
	}

	submitCreateBudget = () => {
		let budgetBody = {
			'name': this.state.name,
			'income_map': this.state.income_map,
			'spending_limit_map': this.state.spending_limit_map,
			'goal_amount_map': this.state.goal_amount_map,
			'frequency': this.state.frequency,
			'savings': parseFloat(this.state.savings),
		}
		console.log('The create form was submitted with the following data:', budgetBody);
		this.postBudgetRequest(budgetBody)
		this.handleCreateClose()
	}

	async postBudgetRequest(budgetBody) {
		let res = await axios.post('/api/budget', budgetBody);
		console.log(res);
		this.getAllBudgets();
	}

	handleCreateClose = () => {
		this.cleanBudgetState()
		this.setState({
			isCreateDialogOpen: false
		});
	};

	// edit budget

	handleEditBudgetOpen = (id) => {
		console.log('Edit id: ', id)
		this.setState({
			budget_id: id,
			isEditDialogOpen: true
		});
	}

	submitEditBudget = () => {
		let budgetBody = {
			'name': this.state.name,
			'income_map': this.state.income_map,
			'spending_limit_map': this.state.spending_limit_map,
			'goal_amount_map': this.state.goal_amount_map,
			'frequency': this.state.frequency,
			'savings': parseFloat(this.state.savings),
		}
		console.log('The edit form was submitted with the following data:', budgetBody);
		this.putBudgetRequest(budgetBody)
		this.handleEditClose()
	}

	async putBudgetRequest(budgetBody) {
		let res = await axios.put('/api/budget/'+this.state.budget_id, budgetBody);
		console.log(res);
		this.getAllBudgets();
	}

	handleEditClose = () => {
		this.cleanBudgetState()
		this.setState({
			isEditDialogOpen: false
		});
	};

	// delete budget

	handleDeleteBudget = (id) => {
		console.log('Delete id: ', id)
		this.deleteBudgetRequest(id)
	}

	async deleteBudgetRequest(id) {
		let res = await axios.delete('/api/budget/'+id);
		console.log(res);
		this.getAllBudgets();
	}

	render() {
		return (
			<div className='budgets-inner-container'>
				<div className='header'>
					Budgets
				</div>
                <div className='create-budget-button'>
                    <Button size='large' style={{color: '#00897b'}} onClick={this.handleCreateBudgetOpen} startIcon={<AddCircleIcon />} >
                        <strong>Create a new Budget</strong>
                    </Button>
				</div>

				<div className='budgets-box'>
                    <List sx={{ width: '100%', bgcolor: 'background.paper' }}>
                        {this.state.allBudgets.length ? <p></p> : <h1>Create a New Budget</h1>}
                        {this.state.allBudgets?.map(data => (
                            <div className='budget'>
                                <ListItem key={data.budget_id} id={data.budget_id} style={{display: 'block'}}>
                                    
                                    <div className='budgets-title-box'>										
                                        <h2>{data.name}</h2>
                                        <Box
                                            display={'flex'}
                                            justifyContent={'space-between'}
                                            marginLeft='auto'
                                        >
                                            <div className="" style={{marginLeft: 'auto', marginRight: '5%'}}>
                                                <IconButton onClick={this.handleEditBudgetOpen.bind(this, data.budget_id)}>
                                                    <ModeEditIcon />
                                                </IconButton>
                                                <ReusableBudgetDialog
                                                    title={'Edit Budget'}
                                                    isDialogOpen={this.state.isEditDialogOpen}
                                                    handleChange={this.handleChange}
                                                    handleClose={this.handleEditClose}
                                                    submitMethod={this.submitEditBudget}
                                                />
                                                <IconButton edge='end'onClick={this.handleDeleteBudget.bind(this, data.budget_id)}>
                                                    <DeleteIcon />
                                                </IconButton>
                                            </div>
                                        </Box>
                                    </div>


                                        <Divider variant='middle' />
                                    <h3>Income</h3>

                                    {Object.keys(data.income_map).map(key => (
                                        <div className='budgets-category-box' key={key} >
                                            <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                                <ListItemText
                                                    style={{width: '90%'}}
                                                    primary={transformSnakeCaseToText(key)}
                                                />
                                                <ListItemText
                                                    primary={'$ ' + data.income_map[key]}
                                                />
                                            </div>
                                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                                <Box sx={{ width: '100%', mr: 1 }}>
                                                    <LinearProgress variant="determinate" color="primary" value={Math.round(1500/data.income_map[key] * 100)}/>
                                                </Box>
                                                <Box sx={{ minWidth: 35 }}>
                                                    <Typography variant="body2" color="text.secondary">{`${(Math.round(1500/data.income_map[key] * 100))}%`}</Typography>
                                                </Box>
                                            </Box>
                                        </div>
                                    ))}

                                    <br></br>
                                    <br></br>

                                    <Divider variant='middle' />
                                    <h3>Spending</h3>

                                    {Object.keys(data.spending_limit_map).map(key => (
                                        <div className='budgets-category-box' key={key}>
                                            <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                                <ListItemText
                                                    style={{width: '90%'}}
                                                    primary={transformSnakeCaseToText(key)}
                                                />
                                                <ListItemText
                                                    primary={'$ ' + data.spending_limit_map[key]}
                                                />
                                            </div>
                                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                                <Box sx={{ width: '100%', mr: 1 }}>
                                                    <LinearProgress variant="determinate" color="primary" value={Math.round(50/data.spending_limit_map[key] * 100)}/>
                                                </Box>
                                                <Box sx={{ minWidth: 35 }}>
                                                    <Typography variant="body2" color="text.secondary">{`${(Math.round(50/data.spending_limit_map[key] * 100))}%`}</Typography>
                                                </Box>
                                            </Box>
                                        </div>
                                    ))}  

                                    {/* <h3>Goals</h3>
                                        <Divider variant='middle' />

                                    {Object.keys(data.goal_amount_map).map(key => (
                                        <div className='budgets-category-box' key={key}>
                                            <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                                <ListItemText
                                                    style={{width: '80%'}}
                                                    primary={capitalizeFirstLowercaseRest(key)}
                                                />
                                                <ListItemText
                                                    primary={'$ ' + data.goal_amount_map[key]}
                                                />
                                            </div>
                                        </div>
                                    ))} */}

                                    <br></br>
                                    <br></br>

                                    <Divider variant='middle' />
                                    <h3>Savings</h3>

                                    <div className='budgets-category-box'>
                                        <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                            <ListItemText
                                                style={{width: '90%'}}
                                                primary="Savings"
                                            />
                                            <ListItemText
                                                primary={'$ ' + data.savings}
                                            />
                                        </div>
                                        <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                                <Box sx={{ width: '100%', mr: 1 }}>
                                                    <LinearProgress variant="determinate" color="primary" value={Math.round(150/data.savings * 100)}/>
                                                </Box>
                                                <Box sx={{ minWidth: 35 }}>
                                                    <Typography variant="body2" color="text.secondary">{`${(Math.round(150/data.savings * 100))}%`}</Typography>
                                                </Box>
                                        </Box>
                                    </div>

                                    <br></br>
                                    <br></br>

                                    
                                </ListItem>
                            </div>
                        ))}
                    </List>
				</div>

				<ReusableBudgetDialog
					title={'Add New Budget'}
					isDialogOpen={this.state.isCreateDialogOpen}
					handleChange={this.handleChange}
					handleClose={this.handleCreateClose}
                    handleRemoveFromMap={this.handleRemoveFromMap}
					submitMethod={this.submitCreateBudget}
				/>
			</div>
		);
	}
}

export default Budgets;