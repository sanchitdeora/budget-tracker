import React, { useState } from 'react';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import { Chip } from '@mui/material';

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

export default function GoalsBudgetSelect({handleBudgetIds, currentGoal, allBudgets}) {
  const [budgets, setBudgets] = useState(currentGoal !== undefined ? getBudgetsFromIds(currentGoal.budget_id_list) : []);

  const handleChange = (event) => {
    console.log("entered: ", event.target)
    const {target: { value }} = event;
    console.log("reached with value: ", value)
    setBudgets(value);
    console.log("reached budgets: ", budgets)
    handleBudgetIds(event.target.name, value.map(x => x.budget_id))
  };

  function getBudgetsFromIds(budgetIds) {
    if (budgetIds === undefined || allBudgets === undefined || allBudgets.length === 0) {
      return [];
    }
    var budgetNames = [];
    console.log('allBudgets', allBudgets);
    console.log('finding for id: ', budgetIds);

    let budget
    budgetIds.forEach(id => {
      budget = allBudgets.find(item => item.budget_id === id)
      if (budget !== undefined) {
        // console.log("val to find budget with budget id in map:" , budget)
        budgetNames.push(allBudgets.find(item => item.budget_id === id))
      }
    });

    console.log('budget found ',  budgetNames.map(obj => obj.budget_name));
    return budgetNames;
    // return budgetIds;
  }


  return (
    <div>
      <FormControl className='input-group' sx={{width:'300px'}}>
        <InputLabel id="input-label">Budgets</InputLabel>
        <Select
          name="budget_id_list"
          multiple
          value={budgets}
          onChange={handleChange}
          input={<OutlinedInput label="Budgets" />}
          renderValue={(selected) => (
            <div className='goal-budget-chips'>
              {selected.map((value) => {
                return (
                  <Chip
                    key={value.budget_id}
                    label={value.budget_name}
                    className='goal-budget-chip'
                  />
                );
              })}
            </div>
          )}
          MenuProps={MenuProps}
        >
          {allBudgets.map((item) => (
            <MenuItem
              key={item._id}
              value={item}
            >
              {item.budget_name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </div>
  );
}