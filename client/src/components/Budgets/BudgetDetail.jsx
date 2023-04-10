import React from 'react';
import Divider from '@mui/material/Divider';
import LinearProgress from '@mui/material/LinearProgress';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import ArrowBackIosNewIcon from '@mui/icons-material/ArrowBackIosNew';
import ListItemText from '@mui/material/ListItemText';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import './BudgetDetail.scss';
import { IconButton } from '@mui/material';
import ReusableBudgetDialog from '../../utils/ReusableBudgetDialog';
import { capitalizeFirstLowercaseRest, changeDateFormatToMmDdYyyy } from '../../utils/StringUtils';

class BudgetDetail extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            budget_id: '',
            isEditDialogOpen: false,
        };
        console.log("props for budget detail", this.props);
        console.log("states for budget detail", this.state);
    };

    cleanBudgetState = () => {
        this.setState({
            budget_id: '',
            // isEditDialogOpen: false,
        })
    }

    handleEditBudgetOpen = (id) => {
        console.log('Edit budgte id for budgetDetail: ', id)
        this.setState({
            budget_id: id,
            isEditDialogOpen: true
        });
    }

    submitMethod = () => {
        this.props.submitMethod()
        this.handleEditClose()
    }

    handleEditClose = () => {
        this.cleanBudgetState()
        this.setState({
            isEditDialogOpen: false
        });
    };
    
    // render functions

    renderBudgetMaps = (title, dataMap) =>
    {
        return (
            <div>
                    <div className='budget-detail-map-box'>
                {/* <Divider variant='middle' /> */}
                    <h3>{title}</h3>

                    {Object.keys(dataMap).map(key => (
                        <div className='budgets-category-box' key={key} >
                            <div className='budgets-list-item-text-container' style={{display: 'flex'}}>
                                <ListItemText
                                    style={{width: '90%'}}
                                    primary={dataMap[key].name}
                                />
                                <ListItemText
                                    primary={'$ ' + dataMap[key].amount}
                                />
                            </div>
                            <Box sx={{ display: 'flex', alignItems: 'center' }}>
                                <Box sx={{ width: '100%', mr: 1 }}>
                                    <LinearProgress variant="determinate" color="primary" value={Math.round(500/dataMap[key].amount * 100)}/>
                                </Box>
                                <Box sx={{ minWidth: 35 }}>
                                    <Typography variant="body2" color="text.secondary">{`${(Math.round(500/dataMap[key].amount * 100))}%`}</Typography>
                                </Box>
                            </Box>
                        </div>
                    ))}

                    <br></br>
                    </div>
                    <br></br>
            </div>
        )

    }

    render() {
        return (
            <div className='budgets-inner-container'>
                <div className='header'>
                    Budgets
                </div>

                <div className='budgets-box'>
                    {this.props.budget ? <p></p> : <h3>Add redirect back or error</h3>}
                    <div className='budget'>                            
                        <Box
                            display={'flex'}
                            justifyContent={'space-between'}
                            marginLeft='auto'
                            alignItems='center'
                            >
                            <IconButton style={{height: 'auto'}} onClick={this.props.handleBudgetClose}>
                                <ArrowBackIosNewIcon />
                            </IconButton>
                            
                            <h2>{this.props.budget.budget_name}</h2>

                            <IconButton style={{marginRight: '2%'}} onClick={this.handleEditBudgetOpen.bind(this, this.props.budget.budget_id)}>
                                <ModeEditIcon />
                            </IconButton>
                            <ReusableBudgetDialog
                                title={'Edit Budget'}
                                isDialogOpen={this.state.isEditDialogOpen}
                                handleInputChange={this.props.handleInputChange}
                                handleClose={this.handleEditClose}
                                currentBudget={this.props.budget}
                                submitMethod={this.submitMethod}
                            />
                        </Box>
                        <Box
                            display={'flex'}
                            justifyContent={'space-between'}
                            marginLeft='auto'
                            alignItems='center'
                        >
                            <div className='budget-detail-other-box'>{changeDateFormatToMmDdYyyy(this.props.budget.creation_time)} — {changeDateFormatToMmDdYyyy(this.props.budget.expiration_time)}</div>
                            <div className='budget-detail-other-box'>Frequency: {capitalizeFirstLowercaseRest(this.props.budget.frequency)}</div>
                        </Box>


                        {this.renderBudgetMaps('Income', this.props.budget.income_map)}
                        {this.renderBudgetMaps('Expense', this.props.budget.expense_map)}
                        {this.renderBudgetMaps('Goals', this.props.budget.goal_map)}
                        {this.renderBudgetMaps('Savings', [{name: "Savings", amount : this.props.budget.savings}])}
                    </div>
                </div>
            </div>
        );
    }
}

export default BudgetDetail;