export const TOKEN = 'token'
export const EMAIL = 'email'
export const ACTIVE_PAGE = 'activePage'
export const IS_SURVEY_COMPLETE = 'IsSurveyComplete'

export const HOME = 'Home'
export const TRANSACTIONS = 'Transactions'
export const BILLS = 'Bills'
export const BUDGETS = 'Budgets'
export const GOALS = 'Goals'

export const FREQUENCY_MAP = [
    {'id': 'once', 'value': 'Once'},
    {'id': 'weekly', 'value': 'Weekly'},
    {'id': 'bi_weekly', 'value': 'Every Two Weeks'},
    {'id': 'monthly', 'value': 'Monthly'},
    {'id': 'bi_monthly', 'value': 'Every Two Months'},
    {'id': 'quarterly', 'value': 'Quarterly'},
    {'id': 'half_yearly','value': 'Every Six Month'},
    {'id': 'yearly', 'value': 'Yearly'}
]

export const CATEGORY_MAP = [
    {'id': 'auto_and_transport', 'value': 'Auto & Transport'},
    {'id': 'bills_and_utilities', 'value': 'Bills & Utilities'},
    {'id': 'education', 'value': 'Education'},
    {'id': 'entertainment', 'value': 'Entertainment'},
    {'id': 'food_and_dining', 'value': 'Food & Dining'},
    {'id': 'health_and_fitness','value': 'Health & Fitness'},
    {'id': 'home','value': 'Home'},
    {'id': 'income','value': 'Income'},
    {'id': 'investments','value': 'Investments'},
    {'id': 'personal_care','value': 'Personal Care'},
    {'id': 'pets','value': 'Pets'},
    {'id': 'shopping','value': 'Shopping'},
    {'id': 'taxes','value': 'Taxes'},
    {'id': 'travel','value': 'Travel'},
    {'id': 'uncategorized', 'value': 'Others'}
]

export const TIME_SEQUENCE_MAP = [
    {'id': '10d', 'value': 'Last 10 days'},
    {'id': '01m', 'value': 'Last 1 month'},
    {'id': '03m', 'value': 'Last 3 months'},
    {'id': '01q', 'value': 'Last 1 quarter'},
    {'id': '03q', 'value': 'Last 3 quarters'},
    {'id': '01y','value': 'Last 1 year'},
    {'id': '03y','value': 'Last 3 years'},
    {'id': '05y','value': 'Last 5 years'},
    {'id': '10y','value': 'Last 10 years'},
]

export const TIME_CONVERSION_MAP = [
    {'id': 'd', 'value':'days'},
    {'id': 'w', 'value':'weeks'},
    {'id': 'm', 'value':'months'},
    {'id': 'q', 'value':'quarters'},
    {'id': 'y', 'value':'years'},
]