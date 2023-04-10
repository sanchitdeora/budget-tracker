/* eslint-disable no-sequences */
/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useEffect } from 'react';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import axios from 'axios';

export default function GoalsBudgetSelect({handleBudgetIds}) {
  const [budgets, setBudgets] = useState([]);
  const [allbudgets, setAllBudgets] = useState([]);
  
  useEffect(() => {
    axios.get('/api/budgets').then((res => {
        console.log('get all budgets in goalBudgetSelect: ', res.data.body)
        var keys = ['budget_id', 'name'];
        var resu = res.data.body.map(obj => keys.reduce((a, c) => (obj[c] ? a[c] = obj[c] : c, a), {}));

        console.log("budgets response in goalBudgetSelect: ", resu)
        if (res.data.body != null) {
            setAllBudgets(
                // res.data.body.map(x => {id: x.goal_id, name: x.name})
                res.data.body.map(obj => keys.reduce((a, c) => (obj[c] ? a[c] = obj[c] : c, a), {}))
            )
        }
        console.log('get all budgets state: ', allbudgets)
    }))
  }, allbudgets);

  const handleChange = (event) => {
    // console.log("entered: ", event)
    const {target: { value }} = event;
    // console.log("reached with value: ", value)
    setBudgets(value);
    // console.log("reached budgets: ", budgets)
    handleBudgetIds(event.target.name, value.map(x => x.budget_id))
  };

  return (
    <div>
        <FormControl className='goal-input-group' sx={{ width: 300 }}>
        <InputLabel id="demo-multiple-name-label">Budgets</InputLabel>
        <Select
          labelId="demo-multiple-name-label"
          id={"demo-multiple-name"}
          name="budget_id_list"
          multiple
          value={budgets}
          onChange={handleChange}
          input={<OutlinedInput label="Budgets" />}
        >
          {allbudgets.map((budget) => (
            <MenuItem
              key={budget.budget_id}
              id={budget.budget_id}
              value={budget}
            >
              {budget.name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </div>
  );
}