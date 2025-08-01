const std = @import("std");

pub fn isIsogram(str: []const u8) bool {
    var seen: u26 = 0;

    for (str) |c| {
        if (!std.ascii.isAlphabetic(c)) {
            continue;
        }

        const offset: u5 = @intCast(std.ascii.toLower(c) - 'a');
        const bit: u26 = @as(u26, 1) << offset;

        if ((seen & bit) >> offset == 1) {
            return false;
        } else {
            seen |= bit;
        }
    }

    return true;
}
