/* eslint-disable array-callback-return */
import React from 'react';
import Divider from '@mui/material/Divider';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import Box from '@mui/material/Box';
import LinearProgress from '@mui/material/LinearProgress';
import './GoalDetail.scss';
import { IconButton } from '@mui/material';
import ReusableGoalDialog from './ReusableGoalDialog';
import ArrowBackIosNew from '@mui/icons-material/ArrowBackIosNew';
import { List } from '@material-ui/core';
import { capitalizeFirstLowercaseRest, transformDateFormatToMmDdYyyy, getRemainingDays } from '../../utils/StringUtils';

class GoalDetail extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            goal_id: '',
            goal_budget_map: {},

            isEditDialogOpen: false,            
        };

        console.log("props for goal", this.props);
        console.log("states for goal", this.state);
        // this.props.goal.budget_id_list.forEach(this.getGoalBudgetsByIds);
        this.getGoalBudgetsByIds(this.props.goal.budget_id_list)
    };

    
    // get budgets for budget id list

    getGoalBudgetsByIds(budget_id_list) {
        console.log("get budget list for goal: ", budget_id_list)
        budget_id_list.forEach(budgetId => this.getBudgetById(budgetId));
    }

    getBudgetById(budget_id) {
        var budget = this.props.allBudgets.find(budget => budget.budget_id === budget_id)
        console.log('get budget: ', budget)
        if (budget !== undefined || budget != null) {
            let goal_budget_map = this.state.goal_budget_map
            goal_budget_map[budget_id] = budget
            this.setState({
                goal_budget_map: goal_budget_map
            })
        }
    }

    // state handlers

    cleanGoalState = () => {
        this.setState({
            goal_id: '',
            isEditDialogOpen: false,            
        })
    }

    handleEditGoalOpen = (id) => {
        console.log('Edit goal id for goalDetail: ', id)
        this.setState({
            goal_id: id,
            isEditDialogOpen: true
        });
    }

    handleEditClose = () => {
        this.cleanGoalState()
        this.setState({
            isEditDialogOpen: false
        });
    };
    
    submitMethod = () => {
        this.props.submitMethod()
        this.handleEditClose()
    }


    //utils

    fetchGoalCurrentAmountFromBudget = (budget, goalId) => {
        var goal = budget.goal_map.find(goal => goal.id === goalId)
        if (goal !== undefined && goal.current_amount > 0) {
            return goal.current_amount
        }
        return 0
    }

    
    // render functions
    
    renderBudgetDetails = (title, goalBudgetMap) =>
    {
        return (
            <div>
                <div className='goal-detail-budget-map-container'>
                {/* <Divider variant='middle' /> */}
                    <h3>{title}</h3>
                    <div>
                        <List sx={{ width: '100%' }}>
                            {/* {goalBudgetMap.length ? <p></p> : <p>Oops! No Budgets found</p>} */}
                            {Object.keys(goalBudgetMap).map(key => (
                                <div className='goal-budgets'>
                                    <div className='goal-budget-detail' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                        <div>
                                            {capitalizeFirstLowercaseRest(goalBudgetMap[key].budget_name)}
                                        </div>
                                        <div>
                                            {'$ ' + this.fetchGoalCurrentAmountFromBudget(goalBudgetMap[key], this.props.goal.goal_id)}
                                        </div>
                                    </div>
                                    <div className='goal-budget-detail' style={{display: 'flex', marginTop: '0px', paddingTop: '5px', paddingBottom: '5px'}}>
                                        <div id='secondary'>
                                            {transformDateFormatToMmDdYyyy(goalBudgetMap[key].creation_time) + " — " + transformDateFormatToMmDdYyyy(goalBudgetMap[key].expiration_time)}
                                        </div>
                                        <div id='secondary'>
                                            {goalBudgetMap[key].is_closed? "Closed" : "Open"}
                                        </div>
                                    </div>
                                    <Divider variant='fullWidth' component='li' />
                                </div>
                            ))}
                        </List>
                    </div>
                <br></br>
                </div>
                <br></br>
            </div>
        )
    }

    render() {
        return (
            <div className='goal-detail-container'>
                <IconButton style={{padding: '2%'}} onClick={this.props.handleGoalClose}>
                    <ArrowBackIosNew />
                </IconButton>

                {this.props.goal ? <h2 className='header goal-detail-header'>{this.props.goal.goal_name}</h2> : <h3>Add navigate back or error</h3>}

                <IconButton style={{marginRight: '2%', padding: '2%'}} onClick={this.handleEditGoalOpen.bind(this, this.props.goal.goal_id)}>
                    <ModeEditIcon />
                </IconButton>
                <ReusableGoalDialog
                    title={'Edit Goal'}
                    currentGoal={this.props.goal}
                    allBudgets={this.props.allBudgets}
                    isDialogOpen={this.state.isEditDialogOpen}
                    handleChange={this.props.handleChange}
                    handleClose={this.handleEditClose}
                    handleBudgetIds={this.props.handleBudgetIds}
                    submitMethod={this.submitMethod}
                />

                {/* <div className='goal'> */}
                    <div className='goal-left-panel'>
                        
                        <div className='goal-detail-info-cards goal-detail-info-cards-dark'>
                            <h3>Target Amount</h3>
                            <br/>
                            {`$ ` + this.props.goal.target_amount}
                        </div>
                        
                        <div className='goal-detail-info-cards'>
                            <h3>Remaining Amount</h3>
                            <br/>
                            {`$ ` + (this.props.goal.target_amount - this.props.goal.current_amount)}
                        </div>
                    </div>

                    <div className='goal-center-panel'>
                        <div className='goal-detail-item-box'>
                            <div className='goal-detail-item'>
                                <div>
                                    {`$ ` + this.props.goal.current_amount}
                                </div>
                                <div>
                                    {`$ ` + this.props.goal.target_amount}
                                </div>
                            </div>
                            <div className='goal-detail-item'>
                                <Box sx={{ width: '100%', mr: 1.5, mt: 1.2 }}>
                                    <LinearProgress variant="determinate" color="primary" value={Math.round(this.props.goal.current_amount/this.props.goal.target_amount * 100)}/>
                                </Box>
                                <Box sx={{ minWidth: 35 }} id='complete-percentage'>
                                    {`${(Math.round(this.props.goal.current_amount/this.props.goal.target_amount * 100))}%`}
                                </Box>
                            </div>
                        </div>
                        <div className='goal-detail-budgets goal-detail-info-cards-dark'>
                            {this.renderBudgetDetails("Goal Budgets", this.state.goal_budget_map)}
                        </div>
                    </div>
                    
                    <div className='goal-right-panel'>
                        <div className='goal-detail-info-cards goal-detail-info-cards-dark'>
                            <h3>Target Date</h3>
                            <br/>
                            {transformDateFormatToMmDdYyyy(this.props.goal.target_date.substring(0, 10))}
                        </div>
                        <div className='goal-detail-info-cards'>
                            <h3>Days Remaining</h3>
                            <br/>
                            {getRemainingDays(this.props.goal.target_date)}
                        </div>
                    </div>
            </div>
        );
    }
}

export default GoalDetail;