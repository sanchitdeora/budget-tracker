import React from "react"; 
import axios from 'axios';
import { Card, CardContent, CardActions, IconButton } from '@mui/material';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import DeleteIcon from '@mui/icons-material/Delete';

import './BudgetCards.scss';
import '../../utils/FilterButton.scss';
import BudgetDetail from "./BudgetDetail";
import ReusableBudgetDialog from './ReusableBudgetDialog';
import { getFullMonthName, getYear } from "../../utils/StringUtils";
import { EXPENSES, GOALS, INCOMES } from "./BudgetConstants";
import FilterButton from "../../utils/FilterButton";
import { BUDGETS } from '../../utils/GlobalConstants';
import { getMenuItemsByTitle } from '../../utils/menuItems';

class BudgetCards extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            allBudgets: [],

            filteredBudgets: [],
            filterCategories: [],
            is_closed_budget_displayed: false,

            budget_id: '',
            budget_name: '',
            income_map: [],
            expense_map: [],
            goal_map: [],
            frequency: '',
            creation_time: new Date(),

            isBudgetOpen: false,
            isCreateDialogOpen: false,
            isEditDialogOpen: false,
        };
        console.log("states for budget card", this.state);
        this.props.setNavBarActive(getMenuItemsByTitle(BUDGETS))
        this.getAllBudgets()
    };


    // get budget

    async getAllBudgets() {
        let res = await axios.get('/api/budgets');
        console.log('get all budget: ', res.data.body)
        var filteredBudgets = [];
        if (res.data.body != null) {
            filteredBudgets = res.data.body.filter(item => !item.is_closed); 
            this.setState({
                allBudgets: res.data.body,
                filteredBudgets: filteredBudgets
            });
        } else {
            this.setState({
                allBudgets: [],
            });
        }

        console.log("Start Filtering by Date", filteredBudgets)
        this.setState({
            filterCategories: ['All', ...new Set(filteredBudgets.map(item => this.getFilterDate(item.creation_time)))]
        });
    }


    // create budgets

    submitCreateBudget = () => {
        let budgetBody = {
            'budget_name': this.state.budget_name,
            'income_map': this.state.income_map,
            'expense_map': this.state.expense_map,
            'goal_map': this.state.goal_map,
            'frequency': this.state.frequency,
            'creation_time': new Date(this.state.creation_time + "T00:00:00-05:00"),
        }
        console.log('The create budgets form was submitted with the following data:', budgetBody);
        this.postBudgetRequest(budgetBody)
        this.handleCreateBudgetClose()
    }

    async postBudgetRequest(budgetBody) {
        let res = await axios.post('/api/budget', budgetBody);
        console.log("Post budgets response", res);
        this.getAllBudgets();
    }

    handleCreateBudgetClose = () => {
        this.cleanBudgetState()
        this.setState({
            isCreateDialogOpen: false
        });
    };

    
    // edit budget

    submitEditBudget = () => {
        let budgetBody = {
            'budget_name': this.state.budget_name,
            'income_map': this.state.income_map,
            'expense_map': this.state.expense_map,
            'goal_map': this.state.goal_map,
            'frequency': this.state.frequency,
            'creation_time': new Date(this.state.creation_time + "T00:00:00-05:00"),
        }
        console.log('The budgets edit form was submitted with the following data:', budgetBody);
        this.putBudgetRequest(budgetBody)
    }

    async putBudgetRequest(budgetBody) {
        let res = await axios.put('/api/budget/'+this.state.budget_id, budgetBody);
        console.log("Put budgets response", res);
        this.getAllBudgets();
    }

    // delete budgets

    handleDeleteBudget = (id) => {
        console.log('Delete budget id: ', id)
        this.deleteBudgetRequest(id)
    }

    async deleteBudgetRequest(id) {
        let res = await axios.delete('/api/budget/'+id);
        console.log("Delete budgets response", res);
        this.getAllBudgets();
    }

    handleInputMapChange = (amountMap) => {
        console.log("getting in main handleInputMapChange, amountMap:", amountMap)
        var name;
        // for handling input map
        if (amountMap !== null && amountMap.type === INCOMES) {
            name="income_map"
        }

        if (amountMap !== null && amountMap.type === EXPENSES) {
            name="expense_map"
        }

        if (amountMap !== null && amountMap.type === GOALS) {
            name="goal_map"
        }
        console.log("setting in main handleInputMapChange, name: ", name, "value:", amountMap.data)
        this.setState({
            [name]: amountMap.data,
        });
    }
    
    handleInputChange = (event, amountMap) => {
        if(amountMap !== undefined) {
            this.handleInputMapChange(amountMap);
            return;
        }

        let value = event.target.value;
        let name = event.target.name;

        console.log("setting in main handleInputChange, name: ", name, "value:", value)
        this.setState({
            [name]: value,
        });
    }

    // utils

    totalAmountFromMap = (data) => {
        if (data === undefined) {
            return "$" + 0;
        }
        // console.log("adding this data", data)
        return "$" + (data.reduce(function (total, item) {
            return total + item.amount;
          }, 0));
    }

    getFilterDate = (stringDate) => {
        console.log("Creation time for budget: ", stringDate)
        return getFullMonthName(stringDate) + " " + getYear(stringDate)
    }

    filterBudgetsByDate = (filterCategory) =>{

        if(filterCategory === 'All') {
            var filteredBudgets
            if (this.state.is_closed_budget_displayed) {
                filteredBudgets = this.filterIsClosed(this.state.allBudgets)
            } else {
                filteredBudgets = this.filterIsOpen(this.state.allBudgets)
            }
            this.setState({
                filteredBudgets: filteredBudgets,
            });
            return;
        }

        const filteredData = this.state.allBudgets.filter(item => this.getFilterDate(item.creation_time) === filterCategory);
        this.setState({
            filteredBudgets: filteredData,
        });
    }

    filterBudgetsClosed = (e) => {
        console.log('budget open: ', e)
        let filteredBudgets
        if (e.target.value === 'closed') {
            filteredBudgets = this.filterIsClosed(this.state.allBudgets)
            this.setState({
                filteredBudgets: filteredBudgets,
                filterCategories: ['All', ...new Set(filteredBudgets.map(item => this.getFilterDate(item.creation_time)))],
                is_closed_budget_displayed: true,
            })
        } else if (e.target.value === 'open') {
            filteredBudgets = this.filterIsOpen(this.state.allBudgets)
            this.setState({
                filteredBudgets: filteredBudgets,
                filterCategories: ['All', ...new Set(filteredBudgets.map(item => this.getFilterDate(item.creation_time)))],
                is_closed_budget_displayed: false,
            })
        }
    }

    filterIsClosed = (budgets) => {
        return budgets.filter(item => item.is_closed)
    }

    filterIsOpen = (budgets) => {
        return budgets.filter(item => !item.is_closed)
    }

    // state handlers

    handleBudgetOpen(budgetId) {
        this.setState({isBudgetOpen: true})
        this.setBudgetState(budgetId)
        console.log('budget open: ', this.state.isBudgetOpen)
        console.log('budget id: ', this.state.budget_id)
    }

    handleBudgetClose() {
        this.cleanBudgetState()
        this.setState({isBudgetOpen: false, budget_id: ''})
        console.log('budget open: ', this.state.isBudgetOpen)
        console.log('budget id: ', this.state.budget_id)
    }

    handleCreateBudgetOpen = () => {
        this.setState({
            isCreateDialogOpen: true
        });
        console.log('isCreateDialogOpen: ', this.state.isCreateDialogOpen)
    }

    cleanBudgetState = () => {
        this.setState({
            budget_id: '',
            budget_name: '',
            income_map: [],
            expense_map: [],
            goal_map: [],
            frequency: '',
            creation_time: new Date(),
        })
    }

    setBudgetState = (budgetId) => {
        let budget = this.state.allBudgets.find(item => item.budget_id === budgetId)
        this.setState({
            budget_id: budget.budget_id,
            budget_name: budget.budget_name,
            income_map: budget.income_map,
            expense_map: budget.expense_map,
            goal_map: budget.goal_map,
            frequency: budget.frequency,
            savings: budget.savings,
            creation_time: budget.creation_time,
        })
    }

    // render functions

    renderFilterDateBoxes() {
        console.log('filter categories: ', this.state.filterCategories)
        console.log('filter budgets: ', this.state.filteredBudgets)
        return(
            <FilterButton key="filter" button={this.state.filterCategories} filter={this.filterBudgetsByDate} />
        )
    }

    renderFilterClosedBoxes() {
        console.log('filter categories: ', this.state.filterCategories)
        console.log('filter budgets: ', this.state.filteredBudgets)
        return(
            <div>
                <label id="filter-closed">Filter Budgets</label>
                <br></br>
                <select
                    name="filter-closed"
                    value={this.state.is_closed_budget_displayed?"closed":"open"}
                    label=""
                    onChange={this.filterBudgetsClosed}
                >
                    <option defaultValue value={"open"}>Open</option>
                    <option value={"closed"}>Closed</option>
                </select>
            </div>
        )

    }

    renderOpenBudgetDialog = () => {
        return (
            <BudgetDetail 
                budget={this.state.allBudgets.find(item => item.budget_id === this.state.budget_id)}
                handleBudgetClose={this.handleBudgetClose.bind(this)}
                handleInputChange={this.handleInputChange}
                submitMethod={this.submitEditBudget}
            >
            </BudgetDetail>
        )
    }

    renderCreateBudgetDialogBox = () => {
        return (
            <ReusableBudgetDialog
                title={'Add New Budget'}
                isDialogOpen={this.state.isCreateDialogOpen}
                handleInputChange={this.handleInputChange}
                handleClose={this.handleCreateBudgetClose}
                currentBudget={{}}
                submitMethod={this.submitCreateBudget}
            />
        )
    }

    renderBudgetCards() {
        return (
            <div className="budget-cards-container">
                <h2 className='header'>
                        {BUDGETS}
                </h2>
                <div className='budget-filter-by-date'>
                    {this.renderFilterDateBoxes()}
                </div>
                <div className='budget-filter-by-close'>
                    {this.renderFilterClosedBoxes()}
                </div>
                
                <div className='budget-cards'>
                    {this.state.filteredBudgets.length ? <p></p> : <h3>Create a New Budget</h3>}
                    <div className='create-budget-card-button'>
                        <IconButton size='large' onClick={this.handleCreateBudgetOpen}>
                            <AddCircleIcon />
                        </IconButton>
                    </div>
                    <div className="budget-cards-box">
                        {this.state.filteredBudgets?.map(data => (
                            <Card className="budget-item">
                                <div onClick={() => this.handleBudgetOpen(data.budget_id)}>
                                    <h4> {data.budget_name}</h4>
                                    <CardContent style={{verticalAlign: 'middle'}}>
                                        <div>
                                            Total Income: {this.totalAmountFromMap(data.income_map)}
                                        </div>
                                        <div>
                                            Target Expense: {this.totalAmountFromMap(data.expense_map)}
                                        </div>
                                        <div>
                                            Target for Goals: {this.totalAmountFromMap(data.goal_map)}
                                        </div>
                                        <div>
                                            Target Savings: {"$" + data.savings}
                                        </div>
                                    </CardContent>
                                </div>
                                <CardActions                                 
                                    style={{display: 'flex', flexDirection: 'row-reverse', marginRight: '5%'}}
                                    // justifyContent={'space-between'}
                                >
                                    <IconButton className='delete-button' edge='end' 
                                        onClick={this.handleDeleteBudget.bind(this, data.budget_id)}>
                                        <DeleteIcon />
                                    </IconButton>
                                </CardActions>
                            </Card>
                        ))}
                    </div>
                </div>
            </div>
        )
    }

    render() {
        return (
            <div className="budget-cards-outer-container">
                {this.state.isBudgetOpen ? this.renderOpenBudgetDialog() : this.renderBudgetCards()}
                {this.renderCreateBudgetDialogBox()}
             </div>
        )
    }
}

export default BudgetCards;
