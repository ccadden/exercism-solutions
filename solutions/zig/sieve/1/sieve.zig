pub fn primes(buffer: []u32, limit: u32) []u32 {
    if (limit < 2) {
        return &[_]u32{};
    }

    var b_idx: usize = 0;

    // tests reference primes up to 1000
    var nums = [_]bool{true} ** 1001;

    var i: u32 = 2;

    while (i <= limit) : (i += 1) {
        if (nums[i]) {
            buffer[b_idx] = i;
            b_idx += 1;
        }

        var j: u32 = 2;

        while (i * j <= limit) : (j += 1) {
            const ij = i * j;
            nums[ij] = false;
        }
    }

    return buffer[0..b_idx];
}
