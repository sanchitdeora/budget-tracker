/* eslint-disable array-callback-return */
import React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import LinearProgress from '@mui/material/LinearProgress';
import './Goals.scss';
import { IconButton } from '@mui/material';
import ReusableGoalDialog from '../../utils/ReusableGoalDialog';
import axios from 'axios';
import { capitalizeFirstLowercaseRest, transformDateFormatToMmDdYyyy } from '../../utils/StringUtils';

class Goals extends React.Component {
    constructor(props) {
        super(props);
        this.INCOME = "Income";
        this.SPENDING = "Spending"
        this.state = {
            goalTest: [{
                goal_id: "G-993064e3-8fec-437b-9484-48cab67d5dcc",
                name: "Vegas",
                current_amount: 1000,
                target_amount: 5000,
                target_date: "2019-12-31T19:00:00-05:00",
                budget_id_list: [
                    "BG-3c8f4fb3-627b-42ec-b170-d22f5dae6904"
                ]
            }],

            allGoals: [],
            goal_id: '',
            name: '',
            current_amount: 0,
            target_amount: 0,
            target_date: '',
            budget_id_list: [],
            isCreateDialogOpen: false,
            isEditDialogOpen: false,
        };
        this.getAllGoals()
    };

    cleanGoalState = () => {
        this.setState({
            goal_id: '',
            name: '',
            current_amount: 0,
            target_amount: 0,
            target_date: '',
            budget_id_list: [],
        })
    }

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

    handleBudgetIds = (name, value) => {
        this.setState({
            [name]: value,
        });
    }

    // get goal

    async getAllGoals() {
        let res = await axios.get('/api/goals');
        console.log('get all goals: ', res.data.body)
        if (res.data.body != null)
        {
            this.setState({
                allGoals: res.data.body,
            });
        } else {
            this.setState({
                allGoals: [],
            });
        }
    }

    // create goal

    handleCreateGoalOpen = () => {
        this.setState({
            isCreateDialogOpen: true
        });
    }

    submitCreateGoal = () => {
        let target_date = new Date(this.state.target_date);
        let goalBody = {
            'name': this.state.name,
            'current_amount': parseFloat(this.state.current_amount),
            'target_amount': parseFloat(this.state.target_amount),
            'target_date': target_date,
            'budget_id_list': this.state.budget_id_list,
        }
        console.log('The create goal form was submitted with the following data:', goalBody);
        this.postGoalRequest(goalBody)
        this.handleCreateClose()
    }

    async postGoalRequest(goalBody) {
        let res = await axios.post('/api/goal', goalBody);
        console.log("post goal response", res);
        this.getAllGoals();
    }

    handleCreateClose = () => {
        this.cleanGoalState()
        this.setState({
            isCreateDialogOpen: false
        });
    };

    // edit goal

    handleEditGoalOpen = (id) => {
        console.log('Edit goal id: ', id)
        this.setState({
            goal_id: id,
            isEditDialogOpen: true
        });
    }

    submitEditGoal = () => {
        let target_date = new Date(this.state.target_date);
        console.log("Goal - Target Date in state: ", this.state.target_date, " date now: ", target_date)
        let goalBody = {
            'name': this.state.name,
            'current_amount': parseFloat(this.state.current_amount),
            'target_amount': parseFloat(this.state.target_amount),
            'target_date': target_date,
            'budget_id_list': this.state.budget_id_list,
        }
        console.log('The goal edit form was submitted with the following data:', goalBody);
        this.putGoalRequest(goalBody)
        this.handleEditClose()
    }

    async putGoalRequest(goalBody) {
        let res = await axios.put('/api/goal/'+this.state.goal_id, goalBody);
        console.log("put goal response", res);
        this.getAllGoals();
    }

    handleEditClose = () => {
        this.cleanGoalState()
        this.setState({
            isEditDialogOpen: false
        });
    };

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

    render() {
        return (
            <div className='goals-inner-container'>
                <div className='header'>
                    Goals
                </div>
                <div className='create-goal-button'>
                    <Button size='large' style={{color: '#00897b'}} onClick={this.handleCreateGoalOpen} startIcon={<AddCircleIcon />} >
                        <strong>Create a new Goal</strong>
                    </Button>
                </div>

                <div className='goals-box'>
                    <List sx={{ width: '100%'}}>
                        {this.state.allGoals.length ? <p></p> : <h1>Create a New Goal</h1>}
                        {this.state.allGoals?.map(data => (
                            <div className='goal'>
                                <ListItem key={data.goal_id} id={data.goal_id}>
                                    <Box sx={{ flexGrow: 1 }}>
                                        <Grid container spacing={2}
                                            sx={{padding: '2px 10px'}}
                                        >
                                            <Grid item xs={10}>
                                                <ListItemText
                                                    primary={capitalizeFirstLowercaseRest(data.name)}
                                                />
                                            </Grid>
                                            <Grid item xs={2}>
                                                <ListItemText
                                                    sx={{textAlign: 'right'}}
                                                    primary={data.target_date.substring(0, 10)}
                                                />
                                            </Grid>
                                            <Grid item xs={12}>
                                                <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                                    <Box sx={{ width: '100%', mr: 1 }}>
                                                        <LinearProgress variant="determinate" color="primary" value={Math.round(data.current_amount/data.target_amount * 100)}/>
                                                    </Box>
                                                    <Box sx={{ minWidth: 35 }}>
                                                        <Typography variant="body2" color="text.secondary">{`${(Math.round(data.current_amount/data.target_amount * 100))}%`}</Typography>
                                                    </Box>
                                                </Box>
                                            </Grid>
                                            <Grid item xs={10}>
                                                <ListItemText
                                                    primary={`$` + data.current_amount}
                                                />
                                            </Grid>
                                            <Grid item xs={2}>
                                                <ListItemText
                                                    sx={{textAlign: 'right'}}
                                                    primary={`$` + data.target_amount}
                                                    />
                                            </Grid>
                                            <Grid item xs={10}>
                                                <div></div>
                                            </Grid>
                                            <Grid item xs={2}>
                                                <Box sx={{display: 'flex', justifyContent: 'flex-end'}}>
                                                        <IconButton onClick={this.handleEditGoalOpen.bind(this, data.goal_id)}>
                                                            <ModeEditIcon />
                                                        </IconButton>
                                                        <ReusableGoalDialog
                                                            title={'Edit Goal'}
                                                            isDialogOpen={this.state.isEditDialogOpen}
                                                            handleChange={this.handleChange}
                                                            handleClose={this.handleEditClose}
                                                            handleBudgetIds={this.handleBudgetIds}
                                                            submitMethod={this.submitEditGoal}
                                                        />
                                                        <IconButton edge='end'onClick={this.handleDeleteGoal.bind(this, data.goal_id)}>
                                                            <DeleteIcon />
                                                        </IconButton>
                                                    </Box>
                                            </Grid>
                                        </Grid>
                                    </Box>



                                    {/* <div className='goals-title-box'>
                                        <ListItemText
                                            style={{width: '65%'}}
                                            primary={capitalizeFirstLowercaseRest(data.name)}
                                        />
                                        <ListItemText
                                            primary={data.target_date.substring(0, 10)}
                                        />
                                    </div>
                                        <Box
                                            display={'flex'}
                                            justifyContent={'space-between'}
                                            marginLeft='auto'
                                        >
                                            <div className="" style={{marginLeft: 'auto', marginRight: '5%'}}>
                                                <IconButton onClick={this.handleEditGoalOpen.bind(this, data.goal_id)}>
                                                    <ModeEditIcon />
                                                </IconButton>
                                                <ReusableGoalDialog
                                                    title={'Edit Goal'}
                                                    isDialogOpen={this.state.isEditDialogOpen}
                                                    handleChange={this.handleChange}
                                                    handleClose={this.handleEditClose}
                                                    submitMethod={this.submitEditGoal}
                                                />
                                                <IconButton edge='end'onClick={this.handleDeleteGoal.bind(this, data.goal_id)}>
                                                    <DeleteIcon />
                                                </IconButton>
                                            </div>
                                        </Box>
                                     */}



                                </ListItem>
                            </div>
                        ))}
                    </List>
                </div>

                <ReusableGoalDialog
                    title={'Add New Goal'}
                    isDialogOpen={this.state.isCreateDialogOpen}
                    handleChange={this.handleChange}
                    handleClose={this.handleCreateClose}
                    handleBudgetIds={this.handleBudgetIds}
                    submitMethod={this.submitCreateGoal}
                />
            </div>
        );
    }
}

export default Goals;