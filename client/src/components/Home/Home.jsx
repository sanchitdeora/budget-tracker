import React from 'react';
import moment from 'moment';
import axios from 'axios';

import { getTxChartData, getLegendsDiffById } from '../Charts/ChartUtils';
import { capitalizeFirstLowercaseRest, findCategoryById, getTimeConversionById } from '../../utils/StringUtils';
import { getMenuItemsByTitle } from '../../utils/menuItems';
import { HOME, TIME_SEQUENCE_MAP } from '../../utils/GlobalConstants';

import { AreaChart, Area, BarChart, Bar, Rectangle, Legend, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

import { Card, List, FormControl, InputLabel, Select, MenuItem } from '@mui/material';

import './Home.scss';

class Home extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allTransactions: [],
            filteredTransactions: [],

            filterByCategories: [],

            filterByDate: {
                startDate: moment().subtract(1, 'month'),
                endDate: moment(),
                id: '01m',
            },

            areaChartData: [],
            barChartDate: [],

            cardInfo: {
                totalIncome: 0,
                totalExpense: 0,
                netIncome: 0,
            }
        };

        this.props.setNavBarActive(getMenuItemsByTitle(HOME))

        this.getAllTransactions()
    };

    async getAllTransactions() {
        let res = await axios.get('/api/transactions');
        console.log('get all transactions: ', res.data.body)
        
        if (res.data.body != null) {
            // sort transactions by date
            let sortedTransactions = res.data.body.sort((p1, p2) => (p1.date < p2.date) ? 1 : (p1.date > p2.date) ? -1 : 0);
            
            
            let filteredTransactions = this.internalCalculateFilteredTransactions(this.state.filterByDate, sortedTransactions);

            this.setState({
                allTransactions: sortedTransactions,
                filteredTransactions: filteredTransactions,
            });

            this.setCardInfo(filteredTransactions);

            this.setChartData(filteredTransactions);

        } else {
            this.setState({
                allTransactions: [],
            });
        }

    }

    calculateCardInfo = (filteredTransactions) => {
        let cardIf = {totalIncome: 0, totalExpense: 0, netIncome: 0}

        cardIf.totalIncome = filteredTransactions.filter(x => x.type).reduce((n, {amount}) => n + amount, 0); 
        cardIf.totalExpense = filteredTransactions.filter(x => !x.type).reduce((n, {amount}) => n + amount, 0);
        cardIf.netIncome = cardIf.totalIncome - cardIf.totalExpense;

        return cardIf;
    }

    setCardInfo = (filteredTransactions) => {
        this.setState({
            cardInfo: this.calculateCardInfo(filteredTransactions)
        })
    }

    setChartData = (data) => {
        this.setState({
            areaChartData: this.createAreaDataSet(data),
            barChartData: this.createBarDataSet(data)
        });
    }

    calculateFilteredTransactions = (event) => {
        console.log("filter range: ", event.target.value);
        let id = event.target.value;

        console.log("post conversion: ", id.substring(2))
        
        console.log("testing dates momentssss end: ", moment())
        console.log("testing dates momentssss start: ", moment().subtract(parseInt(id.substring(0,2)), getTimeConversionById(id.substring(2))))
        
        let filterDate = {
            startDate: moment().subtract(parseInt(id.substring(0,2)), getTimeConversionById(id.substring(2))),
            endDate: moment(),
            id: event.target.value,
        }
        
        let transactions = this.internalCalculateFilteredTransactions(filterDate, this.state.allTransactions)

        
        this.setCardInfo(transactions);
        this.setState({
            filteredTransactions: transactions,
            filterByDate: filterDate
        })

        this.setChartData(transactions);
    }

    internalCalculateFilteredTransactions = (filterDate, transactions) => {
        console.log("internal filter range: ", filterDate, transactions);

        var startDate = filterDate.startDate === null ? moment(0) : filterDate.startDate
        var endDate = filterDate.endDate === null ? moment() : filterDate.endDate
        
        transactions = transactions.filter(
            x => startDate.isBefore(moment(x.date)) && endDate.isAfter(moment(x.date))
        )

        return transactions
    }

    createAreaDataSet = (allTx) => {
    
        let data = [];
        let legend_diff = getLegendsDiffById(this.state.filterByDate.id)
    
        let startDate = this.state.filterByDate.startDate;
        let endDate = startDate.clone().add(parseInt(legend_diff.diff.substring(0,2)), getTimeConversionById(legend_diff.diff.substring(2)))
        // console.log("pre-run start:" , startDate, endDate)

        let i = 1;
        let label = capitalizeFirstLowercaseRest(getTimeConversionById(legend_diff.diff.substring(2)));

        while(this.state.filterByDate.endDate.isAfter(startDate)) {
            // console.log("run start:" , startDate, endDate, i)

            let item = this.createAreaChartDataItem(label + i, startDate, endDate, allTx); 
            data.push(item);
            
            startDate = endDate.clone();
            endDate = startDate.clone().add(parseInt(legend_diff.diff.substring(0,2)), getTimeConversionById(legend_diff.diff.substring(2)))
            i ++;
        }

        console.log("created dataset", data, legend_diff)

        return data
    }

    createAreaChartDataItem = (label, startDate, endDate, tx) => {
        let filteredTx = this.internalCalculateFilteredTransactions({startDate: startDate, endDate: endDate}, tx)
        let cardIf = this.calculateCardInfo(filteredTx)

        return {name: label, Income: cardIf.totalIncome, Expense: cardIf.totalExpense}
    }

    createBarDataSet = (allTx) => {
        console.log("createBarDataSet", getTxChartData(allTx))
        
        return getTxChartData(allTx)
    }


    render() {
        let netLossStyle = {
            border: '1px solid #fa2f2f',
            color: '#fa2f2f !important',
            fontWeight: 'bolder',
        };

        return (
            <div className='home-container'>
                <h2 className='header home-header'>
                    {HOME}
                </h2>
                <div className='filter-button-container'>
                    {this.renderFilterButtons()}
                </div>
                <Card className='display-cards'>
                    <div className='display-amount-lg'>
                        {`$ ` + this.state.cardInfo.totalIncome}
                    </div>
                    <div className='display-label'>
                        Total Income
                    </div>
                </Card>
                <div className='display-graph'>
                    {/* Income vs Expense */}
                    {this.renderAreaChart(`Income vs Expense`)}
                </div>
                <div className='display-graph'>
                    {/* Expense Breakdown */}
                    {this.renderBarChart(`Expense Breakdown`)}
                </div>
                <div></div>
                <Card className='display-cards'>
                    <div className='display-amount-lg' id='center-display'>
                    {`$ ` + this.state.cardInfo.totalExpense}
                    </div>
                    <div className='display-label'>
                        Total Expense
                    </div>
                </Card>
                <div></div>
                <Card className='display-cards' sx={this.state.cardInfo.netIncome < 0 ? netLossStyle : ''}>
                    <div className='display-amount-lg'>
                    {this.state.cardInfo.netIncome > 0 ? `$ ` + this.state.cardInfo.netIncome : `-$ ` + Math.abs(this.state.cardInfo.netIncome)}
                    </div>
                    <div className='display-label'>
                        Net Income
                    </div>
                </Card>

                <div></div>
                <div className='latest-transactions'>
                    {this.renderTransactionDetails('Latest Transactions', this.state.filteredTransactions)}
                </div>
            </div>
        );
    }

    renderFilterButtons = () => {
        return (
            <div className='home-filter-by-date'>
                <FormControl className='input-group' sx={{ width: 300 }}>
                    <InputLabel id="input-label">Filter</InputLabel>
                    <Select
                        name="filterDateRange"
                        labelId="input-label"
                        defaultValue={''}
                        label="Filter"
                        onChange={this.calculateFilteredTransactions}
                        variant="outlined" 
                    >
                        {TIME_SEQUENCE_MAP.map(seq => (
                            <MenuItem key={seq.id} value={seq.id}> {seq.value} </MenuItem>
                        ))}
                    </Select>
                </FormControl>
                <div className='display-timerange-value'>
                    {this.state.filterByDate.startDate.format('MMMM Do YYYY') + ' - ' + this.state.filterByDate.endDate.format('MMMM Do YYYY')}
                </div>
            </div>
        )
    }
    
    renderTransactionDetails = (title, transactionMap) =>
    {
        return (
            <div>
                <h3 className='header latest-tx-header'>{title}</h3>
                <List sx={{ width: '100%' }}>
                    {transactionMap.length ? <p></p> : <p>Oops! No Transactions entered</p>}
                    {transactionMap.slice(0, 10).map(data => (
                        <div className='latest-transaction'>
                            <div className='latest-transaction-info' id='left'>
                                {capitalizeFirstLowercaseRest(data.title)}
                            </div>
                            <div className='latest-transaction-info'>
                                {findCategoryById(data.category)}
                            </div>
                            <div className='latest-transaction-info'>
                                {data.date.substring(0, 10)}
                            </div>
                            <div className='latest-transaction-info' id='right'>
                                {(data.type ? ' ' : '-')+ '$ ' + data.amount}
                            </div>
                        </div>
                    ))}
                </List>
            </div>
        )
    }

    renderAreaChart = (title) => {
        return (
            <div className='areaChart-container'>
                {title}
                    <ResponsiveContainer>
                        <AreaChart
                            className='areaChartIE'
                            data={this.state.areaChartData}
                            margin={{
                                top: 10,
                                right: 30,
                                left: 0,
                                bottom: 0,
                            }}
                            >
                            <XAxis dataKey="name" />
                            <YAxis />
                            <Tooltip />
                            <Legend />
                            <Area type="monotone" dataKey="Income" stroke="#2dfb86" fill="#2dfb86" fillOpacity={0.8} />
                            <Area type="monotone" dataKey="Expense" stroke="#fe1616" fill="#fe1616" fillOpacity={0.2} />
                        </AreaChart>
                    </ResponsiveContainer>
            </div>
        );
    }

    renderBarChart = (title) => {        
        return (
            <div className='areaChart-container'>
                {title}
                <ResponsiveContainer width="100%" height="100%">
                    <BarChart
                        className='areaChartIE'
                        data={this.state.barChartData}
                        margin={{
                            top: 5,
                            right: 30,
                            left: 20,
                            bottom: 5,
                        }}
                        >
                        <CartesianGrid strokeDasharray="3 3" />
                        <XAxis dataKey="name" />
                        <YAxis />
                        <Tooltip />
                        <Legend />
                        <Bar dataKey="value" fill="#2dfb86" activeBar={<Rectangle fill="pink" stroke="blue" />} />
                    </BarChart>
                </ResponsiveContainer>
            </div>
        );
    }

}

export default Home;