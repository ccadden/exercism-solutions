const std = @import("std");

pub const TriangleError = error{
    Invalid,
};

pub const Triangle = struct {
    a: f64,
    b: f64,
    c: f64,

    pub fn init(a: f64, b: f64, c: f64) TriangleError!Triangle {
        if (a == 0 or b == 0 or c == 0) {
            return TriangleError.Invalid;
        }

        var sides = [_]f64{ a, b, c };

        std.mem.sort(f64, &sides, {}, comptime std.sort.asc(f64));

        if (sides[2] > sides[0] + sides[1]) {
            return TriangleError.Invalid;
        }

        return Triangle{ .a = a, .b = b, .c = c };
    }

    pub fn isEquilateral(self: Triangle) bool {
        return self.a == self.b and self.b == self.c;
    }

    pub fn isIsosceles(self: Triangle) bool {
        return self.a == self.b or self.a == self.c or self.c == self.b;
    }

    pub fn isScalene(self: Triangle) bool {
        return self.a != self.b and self.b != self.c and self.c != self.a;
    }
};
