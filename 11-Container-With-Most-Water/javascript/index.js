var maxArea = function (height) {
  let l = 0;
  let r = height.length - 1;
  let maxArea = 0;

  while (l < r) {
    const width = r - l;
    const minHeight = height[r] < height[l] ? height[r] : height[l];
    const area = width * minHeight;

    if (area > maxArea) {
      maxArea = area
    }

    if (height[r] <= height[l]) {
      r--
    } else {
      l++
    }
  }

  return maxArea
};

console.log(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
console.log(maxArea([1, 1]))
console.log(maxArea([8, 7, 2, 1]))