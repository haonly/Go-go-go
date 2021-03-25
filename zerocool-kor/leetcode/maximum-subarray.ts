import Debug from "debug";
const debug = Debug("maximum subarray");
const maxSubArray = (arr: number[]) => {
  debug("maxSubarray started -> ", arr);
  for (let i = 1; i < arr.length; i++) {
    const b = arr[i - 1];
    if (b > 0) {
      debug(`[${i}] arr[i] pre : ${arr[i]}, b:${b}`);
      arr[i] += b;
      debug(`[${i}] arr[i] after : ${arr[i]}`);
    } else {
      debug(`[${i}] skipped.`);
    }
  }
  debug("maxSubarray after -> ", arr);
  return Math.max(...arr);
};
debug("result -> " + maxSubArray([1, -2, 3, 5, -4, 2, 5]));
