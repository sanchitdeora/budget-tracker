import { CATEGORY_MAP, FREQUENCY_MAP } from "./GlobalConstants";

export function capitalizeFirstLowercaseRest(str) {
    var splitStr = str.toLowerCase().split(' ');
    for (var i = 0; i < splitStr.length; i++) {
        splitStr[i] = splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);     
    }
    return splitStr.join(' '); 
};

export function transformSnakeCaseToText(str) {
    var splitStr = str.split('_')
    var resultStr = ""
    splitStr.forEach(element => {
        resultStr = resultStr + " " + capitalizeFirstLowercaseRest(element)
    });
    return resultStr
}

// frequency/category map utils
export function findCategoryById(id) {
    return CATEGORY_MAP.find(x => x.id === id).value
}

export function findFrequencyById(id) {
    return FREQUENCY_MAP.find(x => x.id === id).value
}


// date/time utils

function splitDate(str) {
    var splitStr = str.split('-')
    var year = splitStr[0]
    var month = splitStr[1]
    var day = splitStr[2].substring(0,2)

    return [month, day, year]
}

export function transformDateFormatToYyyyMmDd(str) {
    if (str === undefined) {
        return ''
    }
    return str.substring(0, 10)
}


export function transformDateFormatToMmDdYyyy(str) {
    return splitDate(str).join('-')
}

export function transformDateFormatToMmmDdYyyy(str) {
    var date = splitDate(str)
    date[1] = date[1][0] === '0' ? date[1].substring(1, 2) : date[1]
    
    return getShortMonthName(str) + " " + date[1] + ", " + date[2]
}

export function getFullMonthName(stringDate) {
    const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    return months[new Date(stringDate).getMonth()];
}

export function getShortMonthName(stringDate) {
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];
    return months[new Date(stringDate).getMonth()];
}

export function getYear(stringDate) {
    return new Date(stringDate).getFullYear();
}

export function getRemainingDays(stringDate) {
    let remainingDays = Math.round((new Date(stringDate).getTime() - new Date().getTime()) / (1000 * 60 * 60 * 24))
    return remainingDays > 0 ? remainingDays : 0; 
}
