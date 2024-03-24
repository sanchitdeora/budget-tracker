import React from "react";
import { Card, CardActions, CardContent, IconButton } from "@mui/material";
import AddCircleIcon from '@mui/icons-material/AddCircle';
import DeleteIcon from '@mui/icons-material/Delete';
import axios from "axios";
import ReusableGoalDialog from "./ReusableGoalDialog";
import FilterButton from "../../utils/FilterButton";
import { getFullMonthName, getYear, transformDateFormatToMmDdYyyy } from "../../utils/StringUtils";
import './GoalCards.scss';
import '../../utils/FilterButton.scss';

import GoalDetail from "./GoalDetail";
import { GOALS } from '../../utils/GlobalConstants';
import { getMenuItemsByTitle } from '../../utils/menuItems';



class GoalCards extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            allGoals: [],

            filteredGoals: [],
            filterCategories: [],
                        
            goal_id: '',
            goal_name: '',
            current_amount: 0,
            target_amount: 0,
            target_date: '',
            budget_id_list: [],
            
            isGoalDetailsOpen: false,
            isCreateDialogOpen: false,
            isEditDialogOpen: false,

            allActiveBudgets: []
        };
        console.log("states for goal card", this.state);
        this.props.setNavBarActive(getMenuItemsByTitle(GOALS))

        this.getAllGoals();
        this.getAllBudgets();
    }

    // get budget

    async getAllBudgets() {
        let res = await axios.get('/api/budgets');
        console.log('get all budget: ', res.data.body)
        if (res.data.body != null)
        {
            this.setState({
                allActiveBudgets: res.data.body,
            });
        } else {
            this.setState({
                allActiveBudgets: [],
            });
        }
    }
    

    // get goal

    async getAllGoals() {
        let res = await axios.get('/api/goals');
        console.log('get all goals: ', res.data.body)
        var filteredGoals = [];
        if (res.data.body != null) {
            filteredGoals = res.data.body
            this.setState({
                allGoals: res.data.body,
                filteredGoals: filteredGoals,
            });
        } else {
            this.setState({
                allGoals: [],
            });
        }

        this.setState({
            filterCategories: ['All', ...new Set(filteredGoals.map(item => this.getFilterDate(item.target_date)))]
        });
    }

    // create goal

    submitCreateGoal = () => {
        let target_date = new Date(this.state.target_date);
        let goalBody = {
            'goal_name': this.state.goal_name,
            'current_amount': parseFloat(this.state.current_amount),
            'target_amount': parseFloat(this.state.target_amount),
            'target_date': target_date,
            'budget_id_list': this.state.budget_id_list,
        }
        console.log('The create goal form was submitted with the following data:', goalBody);
        this.postGoalRequest(goalBody)
        this.handleCreateGoalClose()
    }

    async postGoalRequest(goalBody) {
        let res = await axios.post('/api/goal', goalBody);
        console.log("post goal response", res);
        this.getAllGoals();
    }

    // edit goal

    submitEditGoal = () => {
        let target_date = new Date(this.state.target_date);
        console.log("Goal - Target Date in state: ", this.state.target_date, " date now: ", target_date)
        let goalBody = {
            'goal_id': this.state.goal_id,
            'goal_name': this.state.goal_name,
            'current_amount': parseFloat(this.state.current_amount),
            'target_amount': parseFloat(this.state.target_amount),
            'target_date': target_date,
            'budget_id_list': this.state.budget_id_list,
        }
        console.log('The goal edit form was submitted with the following data:', goalBody);
        this.putGoalRequest(goalBody)
    }

    async putGoalRequest(goalBody) {
        let res = await axios.put('/api/goal/'+this.state.goal_id, goalBody);
        console.log("put goal response", res);
        this.getAllGoals();
    }


    // delete goal

    handleDeleteGoal = (id) => {
        console.log('Delete goal id: ', id)
        this.deleteGoalRequest(id)
    }

    async deleteGoalRequest(id) {
        let res = await axios.delete('/api/goal/'+id);
        console.log("delete goal response", res);
        this.getAllGoals();
    }


    // util functions

    getFilterDate = (stringDate) => {
        return getFullMonthName(stringDate) + " " + getYear(stringDate)
    }

    filterGoals = (filterCategory) =>{
        console.log("Control in filter goals: ", filterCategory)
        if(filterCategory === 'All') {
            this.setState({
                filteredGoals: this.state.allGoals,
            });
            return;
        }

        const filteredData = this.state.allGoals.filter(item => this.getFilterDate(item.target_date) === filterCategory);
        this.setState({
            filteredGoals: filteredData,
        });
    }


    // state handlers

    handleChange = (event) => {
        console.log("goal form change event name: " + event.target.name, "event value: " + event.target.value)
        let value = event.target.value;
        let name = event.target.name;
        if (name === 'target_date') {
            value = transformDateFormatToMmDdYyyy(value);
        }
        this.setState({
            [name]: value,
        });
    }

    handleCreateGoalClose = () => {
        this.cleanGoalState()
        this.setState({
            isCreateDialogOpen: false
        });
    };

    handleEditGoalOpen = (id) => {
        console.log('Edit goal id: ', id)
        this.setState({
            goal_id: id,
            isEditDialogOpen: true
        });
    }

    handleGoalDetailsOpen(goalId) {
        this.setState({isGoalDetailsOpen: true})
        this.setGoalState(goalId)
        console.log('goal open: ', this.state.isGoalDetailsOpen)
        console.log('goal id: ', this.state.goal_id)
    }

    handleGoalClose() {
        this.cleanGoalState()
        this.setState({isGoalDetailsOpen: false, goal_id: ''})
        console.log('goal open: ', this.state.isGoalDetailsOpen)
        console.log('goal id: ', this.state.goal_id)
    }

    handleCreateGoalOpen = () => {
        this.setState({
            isCreateDialogOpen: true
        });
        console.log('isCreateDialogOpen: ', this.state.isCreateDialogOpen)
    }

    cleanGoalState = () => {
        this.setState({
            goal_id: '',
            goal_name: '',
            current_amount: 0,
            target_amount: 0,
            target_date: '',
            budget_id_list: [],
        })
    }

    setGoalState = (goalId) => {
        let goal = this.state.allGoals.find(item => item.goal_id === goalId)
        this.setState({
            goal_id: goal.goal_id,
            goal_name: goal.goal_name,
            current_amount: goal.current_amount,
            target_amount: goal.target_amount,
            target_date: goal.target_date,
            budget_id_list: goal.budget_id_list,
        })
    }

    handleBudgetIds = (name, value) => {
        this.setState({
            [name]: value,
        });
    }

    // render functions

    renderFilterBoxes() {
        console.log('filter categories: ', this.state.filterCategories)
        console.log('filter goals: ', this.state.filteredGoals)
        return(
            <FilterButton button={this.state.filterCategories} filter={this.filterGoals} />
        )

    }

    renderCreateGoalDialogBox = () => {
        return (
                <ReusableGoalDialog
                    title={'Add New Goal'}
                    currentGoal={{}}
                    allBudgets={this.state.allActiveBudgets.filter(item => !item.is_closed)}
                    isDialogOpen={this.state.isCreateDialogOpen}
                    handleChange={this.handleChange}
                    handleClose={this.handleCreateGoalClose}
                    handleBudgetIds={this.handleBudgetIds}
                    submitMethod={this.submitCreateGoal}
                />
        )
    }

    renderOpenGoalDialog = () => {
        return (
            <GoalDetail 
                goal={this.state.allGoals.find(item => item.goal_id === this.state.goal_id)}
                allBudgets={this.state.allActiveBudgets.filter(item => !item.is_closed)}
                handleGoalClose={this.handleGoalClose.bind(this)}
                handleChange={this.handleChange}
                handleBudgetIds={this.handleBudgetIds}
                submitMethod={this.submitEditGoal}
            >
            </GoalDetail>
        )
    }

    renderGoalCards() {
        return (
            <div className="goal-cards-container">
                <h2 className='header'>
                    {GOALS}
                </h2>

                <div className='goal-filter-by-date'>
                    {this.renderFilterBoxes()}
                </div>

                <div className='goal-cards'>
                    {this.state.filteredGoals.length ? <p></p> : <h3>Create a New Goal</h3>}
                    <div className='create-goal-card-button'>
                        <IconButton size='large' onClick={this.handleCreateGoalOpen}>
                            <AddCircleIcon />
                        </IconButton>
                    </div>
                    <div className="goal-cards-box">
                        {this.state.filteredGoals?.map(data => (
                            <Card className="goal-item">
                                <div onClick={() => this.handleGoalDetailsOpen(data.goal_id)}>
                                    <h4> {data.goal_name}</h4>
                                    <CardContent style={{verticalAlign: 'middle'}}>
                                        <div>
                                            Current Amount: {data.current_amount}
                                        </div>
                                        <div>
                                            Target Amount: {data.target_amount}
                                        </div>
                                        <div>
                                            Target Date: {data.target_date}
                                        </div>
                                    </CardContent>
                                </div>
                                <CardActions                                 
                                    style={{display: 'flex', flexDirection: 'row-reverse', marginRight: '5%'}}
                                    // justifyContent={'space-between'}
                                >
                                    <IconButton edge='end'onClick={this.handleDeleteGoal.bind(this, data.goal_id)}>
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
            <div className='goal-cards-outer-container'>
                {this.state.isGoalDetailsOpen ? this.renderOpenGoalDialog() : this.renderGoalCards()}
                {this.renderCreateGoalDialogBox()}
             </div>
        )
    }

}

export default GoalCards;