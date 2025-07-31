//
// This is only a SKELETON file for the 'Binary Search' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const find = (list, value) => {
  let left = 0;
  let right = list.length;

  while (left < right) {
    let mid = Math.floor((left + right) / 2);

    const item = list[mid];

    if (item > value) {
      right = mid;
    } else if (item < value) {
      left = mid + 1;
    } else {
      return mid;
    }
  }

  throw new Error("Value not in array");
};
