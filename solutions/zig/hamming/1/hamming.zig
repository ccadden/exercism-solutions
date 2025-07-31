pub const DnaError = error{
    EmptyDnaStrands,
    UnequalDnaStrands,
};

pub fn compute(first: []const u8, second: []const u8) DnaError!usize {
    if (first.len == 0) return DnaError.EmptyDnaStrands;
    if (second.len == 0) return DnaError.EmptyDnaStrands;
    if (first.len != second.len) return DnaError.UnequalDnaStrands;

    var diff: usize = 0;

    for (first, second) |c1, c2| {
        if (c1 != c2) diff += 1;
    }

    return diff;
}
