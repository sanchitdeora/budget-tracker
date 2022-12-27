/* eslint-disable array-callback-return */
import { IconButton, Grid } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import React from "react";
import { transformSnakeCaseToText } from "./StringUtils";

class BudgetMapInput extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            amountMap: []
        };
    }

    addValue = (e) => {
        console.log("here with: ", e.target.value)
        var addFlag = true
        var result = this.state.amountMap
        result.map(x => {
            if(x.value === e.target.value) {
                addFlag = false
            }
        });
        if (addFlag) {
            result = [...result, {
                value: e.target.value,
                amount: 0
            }]
        }
        this.setState({
            amountMap: result
        });
        this.props.handleChange(e)
    };

    addAmount = e => {
        e.preventDefault();

        const id = e.target.id;

        this.state.amountMap.map(x => {
            if(x.value === id) {
                x.amount = e.target.value
            }
        })
        this.setState({
            amountMap: this.state.amountMap
        });
        this.props.handleChange(e)
    };

    removeItem = (item, event) => {
        this.setState({amountMap: this.state.amountMap.filter(function(o) { 
            return o.value !== item.value 
        })});
        this.props.handleRemoveFromMap(item, event)
    };

    render() {
        return (
            <div className="budget-map-input-group">
                <div>
                    <label htmlFor='category'>{this.props.name}</label><br></br>
                    {this.state.amountMap.map((item) => (
                        <div>
                            <Grid container spacing={2}>
                                <Grid item xs={4} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
                                    <label htmlFor='name'>{transformSnakeCaseToText(item.value)}</label>
                                </Grid>
                                <Grid item xs={6}>
                                <div class="currency-wrap">
	                                <span class="currency-code">$</span>
                                    <input
                                        type='number'
                                        name={item.value}
                                        className='budget-input-box'
                                        id={item.value}
                                        onChange={this.addAmount}
                                        style={{paddingLeft: "25px"}}
                                    /></div>
                                </Grid>
                                <Grid item xs={2} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
                                    <IconButton id={item.value} onClick={this.removeItem.bind(this, item)}><DeleteIcon /></IconButton>
                                </Grid>
                            </Grid>
                            <br></br>
                        </div>
                    ))}
                </div>
                <div className='budget-input-group'>
                    <select 
                        name={this.props.name}
                        className='budget-input-box'
                        defaultValue='DEFAULT'
                        onChange={this.addValue}
                    >
                        <option value='DEFAULT' disabled>None</option>
                        <option value='auto_and_transport'>Auto & Transport</option>
                        <option value='bills_and_utilities'>Bills & Utilities</option>
                        <option value='education'>Education</option>
                        <option value='entertainment'>Entertainment</option>
                        <option value='food_and_dining'>Food & Dining</option>
                        <option value='health_and_fitness'>Health & Fitness</option>
                        <option value='home'>Home</option>
                        <option value='income'>Income</option>
                        <option value='investments'>Investments</option>
                        <option value='personal_care'>Personal Care</option>
                        <option value='pets'>Pets</option>
                        <option value='shopping'>Shopping</option>
                        <option value='taxes'>Taxes</option>
                        <option value='travel'>Travel</option>
                        <option value='uncategorized'>Others</option>
                    </select>
                </div>
            </div>
        );
    };
  
}

export default BudgetMapInput;