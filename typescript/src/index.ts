function logItem(val: any) {
  const node = document.createElement('li');
  const textnode = document.createTextNode(val);
  node.appendChild(textnode);
  document.getElementById('list').appendChild(node);
}

/**
 * Given an array, find the int that appears an odd number of times.
 * There will always be only one integer that appears an odd number of times.
 * @param xs
 */
export const findOdd = (xs: number[]): number => {
  const map = xs.reduce((acc, val) => acc.set(val, 1 + (acc.get(val) || 0)), new Map());
  const sorted = [...map.entries()].sort();
  const oddNumberArr = sorted.filter((x) => x[1] % 2 !== 0);
  return oddNumberArr.length !== 0 ? oddNumberArr[0][0] : 0;
};

/**
 * ### Can you find the needle in the haystack?
 * Write a function findNeedle() that takes an array full of junk but containing one "needle"
 * After your function finds the needle it should return a message (as a string) that says:
 * "found the needle at position " plus the index it found the needle, so:
 * ```typescript
 * findNeedle(['hay', 'junk', 'hay', 'hay', 'moreJunk', 'needle', 'randomJunk']
 * ```
 * should return
 * "found the needle at position 5"
 * @param haystack
 */
export function findNeedle(haystack: any[]): string {
  return `found the needle at position ${haystack.indexOf('needle')}`;
}

/**
 * get longest string
 * var strings = ["4999", "longestString", "23", "notLongest"];
 * getLongestString(strings); // output: "longestString"
 * @param array
 */
export function getLongestString(array: string[]): string {
  const length: Array<{ length: number, value: string; }> = [];
  array.forEach((x) => length.push({length: x.length, value: x}));
  return length
      .sort((a, b) => a.length - b.length)
      .reverse()[0].value;
}

/**
 * Recursive sum
 * recursiveSumAll(5); // output: 15  (1 + 2 + 3 + 4 + 5)
 * recursiveSumAll(9); // output: 45  (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9)
 * @param num
 */
export function recursiveSumAll(num: number): number {
  if (num > 0) {
    return recursiveSumAll(num - 1) + num;
  }
  return num;
}

/**
 * var foo = [2, 5, 9, 13, 16, 19, 20, 21];
 * var bar = [1, 5, 7, 8, 12, 22, 29];
 *
 * getLeastKNumber(foo, bar, 4); // output: [1, 2, 5, 5]
 * getLeastKNumber(foo, bar, 6); // output: [1, 2, 5, 5, 7, 8]
 * getLeastKNumber(foo, bar, 7); // output: [1, 2, 5, 5, 7, 8, 9]
 */

function getLeastKNumber(firstSet: number[], secondSet: number[], limit: number): number[] {
  const mergeSet = firstSet.concat(secondSet);
  mergeSet.sort((a, b) => a - b);
  return mergeSet.slice(0, limit);
}

