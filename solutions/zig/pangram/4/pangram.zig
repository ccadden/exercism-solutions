pub fn isPangram(str: []const u8) bool {
    var bits: u26 = 0;

    for (str) |c| {
        var mask: u26 = undefined;
        switch (c) {
            'a'...'z' => {
                mask = @as(u26, 1) << @intCast(c - 'a');
            },
            'A'...'Z' => {
                mask = @as(u26, 1) << @intCast(c - 'A');
            },
            else => continue,
        }

        bits |= mask;
    }

    return bits == 0b11111111111111111111111111;
}
