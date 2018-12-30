function logItem(val: any) {
    var node = document.createElement("li");
    var textnode = document.createTextNode(val);
    node.appendChild(textnode);
    document.getElementById("list").appendChild(node);
}

/**
 * Given an array, find the int that appears an odd number of times.
 * There will always be only one integer that appears an odd number of times.
 * @param xs
 */
export const findOdd = (xs: number[]): number => {
    const map = xs.reduce((acc, val) => acc.set(val, 1 + (acc.get(val) || 0)), new Map());
    const sorted = [...map.entries()].sort();
    const oddNumberArr = sorted.filter(x => x[1] % 2 !== 0);
    return oddNumberArr.length !== 0 ? oddNumberArr[0][0] : 0;
};


