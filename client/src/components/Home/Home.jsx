import React from 'react';
import { getMenuItemsByTitle } from '../../utils/menuItems';
import { HOME } from '../../utils/GlobalConstants';
import './Home.scss';


import { DateRangePicker, SingleDatePicker, DayPickerRangeController } from 'react-dates';
import Box from '@mui/material/Box';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Chip from '@mui/material/Chip';

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            personName: [],
        };
        this.props.setNavBarActive(getMenuItemsByTitle(HOME))
    };

    handleChange = (event) => {
        const {
          target: { value },
        } = event;

        this.setState({ personName: typeof value === 'string' ? value.split(',') : value});
      };
    render() {
        return (
            <div>
                <div className="date-picker">
                <p></p>
                    <DateRangePicker
                        startDate={this.state.startDate}
                        startDateId="your_unique_start_date_id"
                        endDate={this.state.endDate}
                        endDateId="your_unique_end_date_id"
                        onDatesChange={({ startDate, endDate }) => this.setState({ startDate, endDate })}
                        focusedInput={this.state.focusedInput} 
                        onFocusChange={focusedInput => this.setState({ focusedInput })}
                    />
                </div>
                <div className="form-select">
                    <FormControl sx={{ m: 1, width: 300 }} className='multiselect-form-class-mine'>
                        <InputLabel id="demo-multiple-chip-label">Filter </InputLabel>
                        <Select
                            className='ms-select'
                            labelId="demo-multiple-chip-label"
                            id="demo-multiple-chip"
                            multiple
                            value={this.state.personName}
                            onChange={this.handleChange}
                            input={<OutlinedInput id="select-multiple-chip" label="Chip" />}
                            renderValue={(selected) => (
                                <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {selected.map((value) => (
                                    <Chip key={value} label={value} className='ms-chip' />
                                ))}
                                </Box>
                            )}
                            MenuProps={MenuProps}
                            >
                            {names.map((name) => (
                                <MenuItem
                                    key={name}
                                    value={name}
                                    //   style={getStyles(name, personName, theme)}
                                    >
                                    {name}
                                </MenuItem>
                            ))}
                        </Select>
                    </FormControl>
                </div>
            </div>
        );
    }
}

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
    PaperProps: {
      style: {
        maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
        width: 250,
      },
    },
  };

const names = [
    'Oliver Hansen',
    'Van Henry',
    'April Tucker',
    'Ralph Hubbard',
    'Omar Alexander',
    'Carlos Abbott',
    'Miriam Wagner',
    'Bradley Wilkerson',
    'Virginia Andrews',
    'Kelly Snyder',
  ];

export default Home;