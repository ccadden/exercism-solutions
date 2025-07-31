/// Writes a reversed copy of `s` to `buffer`.
pub fn reverse(buffer: []u8, s: []const u8) []u8 {
    if (s.len == 0) {
        return "";
    }

    var i = s.len - 1;

    while (i >= 0) {
        buffer[i] = s[s.len - 1 - i];
        if (i == 0) break;
        i -= 1;
    }

    return buffer;
}
