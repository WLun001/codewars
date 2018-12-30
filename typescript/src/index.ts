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


/**
 * ### Can you find the needle in the haystack?
 Write a function findNeedle() that takes an array full of junk but containing one "needle"
 After your function finds the needle it should return a message (as a string) that says:
 "found the needle at position " plus the index it found the needle, so:
 ```typescript
 findNeedle(['hay', 'junk', 'hay', 'hay', 'moreJunk', 'needle', 'randomJunk']
 ```
 should return
 "found the needle at position 5"
 * @param haystack
 */
export function findNeedle(haystack: any[]): string {
    return `found the needle at position ${haystack.indexOf('needle')}`
}

