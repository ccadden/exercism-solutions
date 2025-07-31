const mem = @import("std").mem;

pub fn toRna(allocator: mem.Allocator, dna: []const u8) mem.Allocator.Error![]const u8 {
    const rna = try allocator.alloc(u8, dna.len);
    errdefer allocator.free(rna);

    for (dna, 0..) |n, idx| {
        switch (n) {
            'G' => rna[idx] = 'C',
            'C' => rna[idx] = 'G',
            'T' => rna[idx] = 'A',
            'A' => rna[idx] = 'U',
            else => continue,
        }
    }

    return rna;
}
