/* eslint-disable array-callback-return */
import { IconButton, Grid } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import React from "react";
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';

class BudgetMapInput extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            amountMap: this.props.currentDataMap
        };

        console.log("props for budget map input", this.props);
        console.log("states for budget map input", this.state);
    }

    addValue = (e) => {
        var addFlag = true
        var result = this.state.amountMap
        var val = e.target.value
        var name = this.props.optionsList.find(o => o.id === val).value

        result.map(x => {
            if(x.id === val) {
                addFlag = false
            }
        });
        if (addFlag) {
            result = [...result, {
                id: val,
                amount: 0,
                name: name
            }]
        }
        this.setState({
            amountMap: result
        });
        console.log("show amountMap in addValue: ", result)
        this.props.handleChange(e, result)
    };

    addAmount = e => {
        e.preventDefault();

        var result = this.state.amountMap
        const id = e.target.id;
        console.log("targetId in addAmount: ", id)

        result.map(x => {
            if(x.id === id) {
                console.log("event in addAmount: ", e)
                x.amount = parseFloat(e.target.value)
                console.log("after update event in addAmount: ", x)
            }
        })

        console.log("show amountMap in addAmount: ", {type: "income", data: result})

        this.setState({amountMap: result});

        this.props.handleChange(e, result)
    };

    removeItem = (item, e) => {
        console.log("item removal: ", item);
        // e.target.reset();

        var result = this.state.amountMap

        result = result.filter(o => {
            console.log("item o: ", o, "matching with: ", item);
            console.log("is matching?: ", o.id === item.id); 
            return o.id !== item.id
        })

        this.setState({amountMap: result});

        console.log("map after removal: ", result);
        this.props.handleChange(e, result)
    };

    render() {
        return (
            <div className="budget-map-input-group">
                <div>
                    {this.state.amountMap.map((item) => (
                        <div>
                            <Grid container spacing={3}>
                                <Grid item xs={4} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
                                    <label htmlFor={item.name}>{item.name}</label>
                                </Grid>
                                <Grid item xs={6}>
                                    <div className="currency-wrap">
                                        <span className="currency-code">$</span>
                                        <input
                                            value={item.amount}
                                            type='number'
                                            name={item.name}
                                            className='budget-input-box'
                                            id={item.id}
                                            onChange={this.addAmount}
                                            style={{paddingLeft: "25px"}}
                                        />
                                    </div>
                                </Grid>
                                <Grid item xs={2} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
                                    <IconButton id={item.id} onClick={this.removeItem.bind(this, item)}><DeleteIcon /></IconButton>
                                </Grid>
                            </Grid>
                            <br></br>
                        </div>
                    ))}
                </div>
                <div className='budget-input-group'>
                    <FormControl className='goal-input-group' sx={{ width: 300 }}>
                        <InputLabel id="demo-multiple-name-label">{this.props.name}</InputLabel>
                        <Select
                        labelId="demo-multiple-name-label"
                        id={"demo-multiple-name"}
                        name={this.props.name}
                        value=''
                        onChange={this.addValue}
                        input={<OutlinedInput label={this.props.name} />}
                        >
                            {this.props.optionsList.map((category) => (
                            <MenuItem
                                key={category.id}
                                id={category.id}
                                value={category.id}
                            >
                                {category.value}
                            </MenuItem>
                            ))}
                        </Select>
                    </FormControl>
                </div>
            </div>
        );
    };
  
}

export default BudgetMapInput;