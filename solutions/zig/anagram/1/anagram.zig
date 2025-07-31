const std = @import("std");
const mem = std.mem;

/// Returns the set of strings in `candidates` that are anagrams of `word`.
/// Caller owns the returned memory.
pub fn detectAnagrams(
    allocator: mem.Allocator,
    word: []const u8,
    candidates: []const []const u8,
) !std.BufSet {
    var result_set = std.BufSet.init(allocator);
    errdefer result_set.deinit();

    const sorted_word = try std.ascii.allocLowerString(allocator, word);
    defer allocator.free(sorted_word);

    mem.sort(u8, sorted_word, {}, std.sort.asc(u8));

    for (candidates) |c| {
        if (word.len != c.len) {
            continue;
        }

        if (std.ascii.eqlIgnoreCase(word, c)) {
            continue;
        }

        const sorted_c = try std.ascii.allocLowerString(allocator, c);
        defer allocator.free(sorted_c);

        mem.sort(u8, sorted_c, {}, std.sort.asc(u8));

        if (std.mem.eql(u8, sorted_word, sorted_c)) {
            try result_set.insert(c);
        }
    }

    return result_set;
}
