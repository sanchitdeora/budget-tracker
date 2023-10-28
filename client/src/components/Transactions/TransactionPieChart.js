import { findCategoryById } from "../../utils/StringUtils";

export function getTxChartData(txData) {        
    return aggregateData(txData.map(item => ({name: findCategoryById(item.category), value: Math.abs(item.amount)})))
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
