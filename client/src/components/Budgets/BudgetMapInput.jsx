/* eslint-disable array-callback-return */
import { IconButton, Grid, FormLabel } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import React from "react";
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import TextField from '@mui/material/TextField';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import InputAdornment from '@mui/material/InputAdornment';
import '../../utils/Dialog.scss';

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
            <div className="input-group">
                <FormControl className='input-group' sx={{ width: 300 }}>
                    <InputLabel id="input-label">{this.props.name}</InputLabel>
                    <Select
                        labelId="input-label"
                        name={this.props.name}
                        value=''
                        onChange={this.addValue}
                        variant="outlined"
                        input={<OutlinedInput label={this.props.name} />}
                    >
                        {this.props.optionsList.map(option => (
                            <MenuItem key={option.id} value={option.id} id={option.id}> {option.value} </MenuItem>
                        ))} 
                    </Select>
                </FormControl>
                <div className="input-group-inner">
                    {this.state.amountMap.map((item) => (
                        <Grid container spacing={2}>
                            <Grid item xs={3} style={{display: 'flex', alignItems: 'center', justifyContent: 'center'}}>
                                <FormLabel htmlFor={item.name}>{item.name}</FormLabel>
                            </Grid>
                            <Grid item xs={7}>
                                <FormControl className='input-group' sx={{ width: 300 }}>
                                <TextField 
                                    value={item.amount}
                                    type='number'
                                    name={item.name}
                                    label='Amount'
                                    id={item.id}
                                    className='budget-input-map-box'
                                    InputProps={{
                                        startAdornment: 
                                        <InputAdornment disableTypography position="start">
                                            $</InputAdornment>,
                                        inputMode: 'numeric', pattern: '[0-9]*' 
                                    }}
                                    onChange={this.addAmount}
                                    variant="outlined" 
                                />
                            </FormControl>
                            </Grid>
                            <Grid item xs={2} style={{display: 'flex', alignItems: 'right', justifyContent: 'right'}}>
                                <IconButton id={item.id} onClick={this.removeItem.bind(this, item)}><DeleteIcon /></IconButton>
                            </Grid>
                        </Grid>
                    ))}
                </div>
            </div>
        );
    };
  
}

export default BudgetMapInput;