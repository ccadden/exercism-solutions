const std = @import("std");

pub fn response(s: []const u8) []const u8 {
    if (silence(s)) return "Fine. Be that way!";

    var flags: u2 = 0;
    if (question(s)) flags |= 1 << 1;
    if (screamCase(s)) flags |= 1;

    return switch (flags) {
        0b10 => "Sure.",
        0b11 => "Calm down, I know what I'm doing!",
        0b01 => "Whoa, chill out!",
        else => "Whatever.",
    };
}

fn question(s: []const u8) bool {
    var i = s.len - 1;

    while (i > 0) : (i -= 1) {
        if (std.ascii.isWhitespace(s[i])) continue;
        if (s[i] == '?') {
            return true;
        } else {
            return false;
        }
    }

    return true;
}

fn silence(s: []const u8) bool {
    for (s) |c| {
        if (!std.ascii.isWhitespace(c)) {
            return false;
        }
    }

    return true;
}

fn screamCase(s: []const u8) bool {
    var hasLetter = false;
    for (s) |c| {
        if (std.ascii.isWhitespace(c) or !std.ascii.isAlphabetic(c)) continue;
        if (!std.ascii.isUpper(c)) {
            return false;
        } else {
            hasLetter = true;
        }
    }

    return true and hasLetter;
}
