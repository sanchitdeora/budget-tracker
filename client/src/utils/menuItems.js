import {dashboard, expenses, transactions, trend} from '../utils/Icons'

export const menuItems = [
    {
        id: 1,
        title: 'Home',
        icon: dashboard,
        link: '/home'
    },
    {
        id: 2,
        title: "Transactions",
        icon: transactions,
        link: "/transactions",
    },
    {
        id: 3,
        title: "Bills",
        icon: trend,
        link: "/bills",
    },
    {
        id: 4,
        title: "Budgets",
        icon: expenses,
        link: "/budgets",
    },
    {
        id: 5,
        title: "Goals",
        icon: trend,
        link: "/goals",
    },
]