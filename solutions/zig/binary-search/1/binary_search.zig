// Take a look at the tests, you might have to change the function arguments

pub fn binarySearch(comptime T: type, target: T, items: []const T) ?usize {
    var left: usize = 0;
    var right: usize = items.len;

    while (left < right) {
        const mid: usize = (left + right) / 2;

        if (items[mid] > target) {
            right = mid;
        } else if (items[mid] < target) {
            left = mid + 1;
        } else {
            return mid;
        }
    }

    return null;
}
