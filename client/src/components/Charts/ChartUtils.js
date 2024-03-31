import { findCategoryById } from "../../utils/StringUtils";

export function getTxChartData(txData) {        
    return aggregateData(txData.map(item => ({name: findCategoryById(item.category), value: Math.abs(item.amount)})))
}

export function getBillsChartData(txData) {        
    return aggregateData(txData.map(item => ({name: findCategoryById(item.category), value: Math.abs(item.amount_due)})))
}

function aggregateData(data) {
    let a = data.reduce((acc, x) => {
        if(acc.find(y => y.name === x.name)) return acc.concat([]);
        const value = data.filter(y => y.name === x.name).map(y => y.value).reduce((a, b) => a + b, 0);
        return acc.concat([{
            ...x,
            value,
        }])}, []
    );
        
    let b = a.sort((a, b) => (a.value >= b.value)? -1:1).map((item, index) => ({...item, fill: COLORS[index % COLORS.length]}));

    return b;
}

const COLORS = ["#4290b5", "#6cb9fe", "#b4f9cd", "#00bf64", "#2d7b60", "#6ea499", "#1a2d44", "#e14b31", "#c23728"]

const LEGENDS_DIFF = [
    {id: '10d', diff: '01d', count: '10' },
    {id: '01m', diff: '01w', count: '04' },
    {id: '03m', diff: '01w', count: '12' },
    {id: '01q', diff: '01m', count: '04' },
    {id: '03q', diff: '01m', count: '12' },
    {id: '01y', diff: '01m', count: '12' },
    {id: '03y', diff: '01q', count: '12' },
    {id: '05y', diff: '01y', count: '05' },
    {id: '10y', diff: '01y', count: '10' },
]

export const TIME_CONVERSION_MAP = [
    {'id': 'd', 'value':'days'},
    {'id': 'm', 'value':'months'},
    {'id': 'q', 'value':'quarters'},
    {'id': 'y', 'value':'years'},
]

export function getLegendsDiffById(id) {
    return LEGENDS_DIFF.find(x => x.id === id)
}