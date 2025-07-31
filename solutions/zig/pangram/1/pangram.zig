const ascii = @import("std").ascii;

pub fn isPangram(str: []const u8) bool {
    var buffer: [1000]u8 = undefined;

    _ = ascii.upperString(&buffer, str);
    const letters: []const u8 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

    for (letters) |c| {
        if (!stringContains(&buffer, c)) {
            return false;
        }
    }

    return true;
}

fn stringContains(str: []const u8, char: u8) bool {
    for (str) |c| {
        if (c == char) {
            return true;
        }
    }

    return false;
}
