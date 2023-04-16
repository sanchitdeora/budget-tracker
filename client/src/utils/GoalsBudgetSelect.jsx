import React, { useState } from 'react';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import { makeStyles, useTheme } from '@material-ui/core/styles';
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

const useStyles = makeStyles((theme) => ({
  formControl: {
    minWidth: 120,
    maxWidth: 300,
  },
  chips: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  chip: {
    margin: 2,
  },
  noLabel: {
    marginTop: theme.spacing(3),
  },
}));

function getStyles(budget, allBudgets, theme) {
  return {
    fontWeight:
      allBudgets.indexOf(budget) === -1
        ? theme.typography.fontWeightRegular
        : theme.typography.fontWeightMedium,
  };
}

export default function GoalsBudgetSelect({handleBudgetIds, currentGoal, allBudgets}) {
  const theme = useTheme();
  const classes = useStyles();
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
    if (allBudgets === undefined || allBudgets.length === 0) {
      return [];
    }
    var budgetNames = [];
    console.log('allBudgets', allBudgets);
    console.log('finding for id: ', budgetIds);

    budgetIds.forEach(id => {
      budgetNames.push(allBudgets.find(item => item.budget_id === id))
    });

    console.log('budget found ',  budgetNames.map(obj => obj.budget_name));
    return budgetNames;
    // return budgetIds;
  }


  return (
    <div>
      <FormControl className={classes.formControl}>
        <InputLabel id="demo-mutiple-chip-label">Budgets</InputLabel>
        <Select
          name="budget_id_list"
          labelId="demo-mutiple-chip-label"
          id="demo-mutiple-chip"
          multiple
          value={budgets}
          onChange={handleChange}
          input={<OutlinedInput label="Budgets" />}
          renderValue={(selected) => (
            <div className={classes.chips}>
              {selected.map((value) => {
                return (
                  <Chip
                    key={value.budget_id}
                    label={value.budget_name}
                    className={classes.chip}
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
              style={getStyles(
                item.budget_id,
                budgets.map((item) => item.budget_id),
                theme
              )}
            >
              {item.budget_name}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </div>
  );
}