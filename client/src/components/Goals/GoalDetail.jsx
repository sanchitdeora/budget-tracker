/* eslint-disable array-callback-return */
import React from 'react';
import Divider from '@mui/material/Divider';
import ListItemText from '@mui/material/ListItemText';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import LinearProgress from '@mui/material/LinearProgress';
import './GoalDetail.scss';
import { IconButton } from '@mui/material';
import ReusableGoalDialog from '../../utils/ReusableGoalDialog';
import ArrowBackIosNew from '@mui/icons-material/ArrowBackIosNew';

class GoalDetail extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            goal_id: '',
            isEditDialogOpen: false,            
        };

        console.log("props for goal", this.props);
        console.log("states for goal", this.state);
    };

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

    
    // render functions

    render() {
        return (
            <div className='goals-inner-container'>

                <div className='goals-box'>
                    {this.props.goal ? <p></p> : <h3>Add redirect back or error</h3>}
                    <div className='goal'>
                    <Box
                            className='goal-detail-header-box'
                            display={'flex'}
                            justifyContent={'space-between'}
                            marginLeft='auto'
                            alignItems='center'
                            >
                            <IconButton style={{padding: '2%'}} onClick={this.props.handleGoalClose}>
                                <ArrowBackIosNew />
                            </IconButton>

                            <h2>{this.props.goal.goal_name}</h2>

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
                        </Box>
                        <br></br>
                        <Divider variant='middle' />
                        <br></br>

                        <Box sx={{ flexGrow: 1 }}>
                            <Grid container spacing={2}
                                sx={{padding: '2px 10px'}}
                            >
                                <Grid item xs={10}>
                                    <ListItemText
                                        primary={`$` + this.props.goal.current_amount}
                                    />
                                </Grid>
                                <Grid item xs={2}>
                                    <ListItemText
                                        sx={{textAlign: 'right'}}
                                        primary={`$` + this.props.goal.target_amount}
                                        />
                                </Grid>
                                <Grid item xs={12}>
                                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                        <Box sx={{ width: '100%', mr: 1 }}>
                                            <LinearProgress variant="determinate" color="primary" value={Math.round(this.props.goal.current_amount/this.props.goal.target_amount * 100)}/>
                                        </Box>
                                        <Box sx={{ minWidth: 35 }}>
                                            <Typography variant="body2" color="text.secondary">{`${(Math.round(this.props.goal.current_amount/this.props.goal.target_amount * 100))}%`}</Typography>
                                        </Box>
                                    </Box>
                                </Grid>
                                <Grid item xs={10}>
                                    <ListItemText
                                        primary={`Target Date: `}
                                        secondary={`Add more information; Eg: about budgets`}
                                    />
                                </Grid>
                                <Grid item xs={2}>
                                    <ListItemText
                                        sx={{textAlign: 'right'}}
                                        primary={this.props.goal.target_date.substring(0, 10)}
                                    />
                                </Grid>
                                <Grid item xs={10}>
                                    <div></div>
                                </Grid>
                            </Grid>
                        </Box>
                    </div>
                </div>
            </div>
        );
    }
}

export default GoalDetail;