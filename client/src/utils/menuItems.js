import {home, expenses, transactions, bill, goal} from '../utils/Icons'
import { BILLS, BUDGETS, GOALS, HOME, TRANSACTIONS } from './GlobalConstants'



export const menuItems = [
    {
        id: 0,
        title: HOME,
        icon: home,
        link: '/home'
    },
    {
        id: 1,
        title: TRANSACTIONS,
        icon: transactions,
        link: '/transactions',
    },
    {
        id: 2,
        title: BILLS,
        icon: bill,
        link: '/bills',
    },
    {
        id: 3,
        title: BUDGETS,
        icon: expenses,
        link: '/budgets',
    },
    {
        id: 4,
        title: GOALS,
        icon: goal,
        link: '/goals',
    },
]

export function getMenuItemsByTitle(title) {
    return menuItems.find(item => item.title === title)
}