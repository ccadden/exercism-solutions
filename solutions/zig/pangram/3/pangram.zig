pub fn isPangram(str: []const u8) bool {
    var bits: u26 = 0;

    for (str) |c| {
        var mask: u26 = 0;
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

    return bits == ~(@as(u26, 0) & 0);
}
