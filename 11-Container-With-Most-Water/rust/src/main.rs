pub fn max_area(height: Vec<i32>) -> i32 {
    let (mut left, mut right) = (0, height.len() - 1);
    let mut result: i32 = 0;

    while left < right {
        let width: i32 = (right - left).try_into().unwrap();
        let mut minHeight = height[left];
        
        if minHeight > height[right] {
            minHeight = height[right]
        }

        let area: i32 = minHeight * width;

        if area > result {
            result = area;
        }

        if height[right] <= height[left] {
            right -= 1
        } else {
            left += 1
        }
    }

    result
}

fn main() {
    let height = vec![1, 8, 6, 2, 5, 4, 8, 3, 7];
    print!("{}", max_area(height))
}
