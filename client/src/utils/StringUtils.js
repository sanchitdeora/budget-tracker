export function capitalizeFirstLowercaseRest(str) {
    var splitStr = str.toLowerCase().split(' ');
    for (var i = 0; i < splitStr.length; i++) {
        splitStr[i] = splitStr[i].charAt(0).toUpperCase() + splitStr[i].substring(1);     
    }
    return splitStr.join(' '); 
};

export function changeDateFormatToMmDdYyyy(str) {
    var splitStr = str.split('-')
    var year = splitStr[0]
    var month = splitStr[1]
    var day = splitStr[2]

    return [month, day, year].join('-')
}

export function transformSnakeCaseToText(str) {
    var splitStr = str.split('_')
    var resultStr = ""
    splitStr.forEach(element => {
        resultStr = resultStr + " " + capitalizeFirstLowercaseRest(element)
    });
    return resultStr
}