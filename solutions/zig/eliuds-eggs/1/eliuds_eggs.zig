pub fn eggCount(number: usize) usize {
    var count: usize = 0;
    var n = number;
    while (n != 0) {
        if (n % 2 != 0) {
            count += 1;
        }

        n /= 2;
    }

    return count;
}
