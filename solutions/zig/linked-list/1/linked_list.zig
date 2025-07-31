pub fn LinkedList(comptime T: type) type {
    return struct {
        pub const Node = struct {
            prev: ?*Node = null,
            next: ?*Node = null,
            data: T,
        };

        first: ?*Node = null,
        last: ?*Node = null,
        len: usize = 0,

        pub fn push(self: *@This(), node: *Node) void {
            if (self.last) |last| {
                last.next = node;
                node.prev = last;
            }

            if (self.first == null) {
                self.first = node;
            }

            self.last = node;
            self.len += 1;
        }

        pub fn pop(self: *@This()) ?*Node {
            if (self.last) |last| {
                self.last = last.prev;
                self.len -= 1;

                return last;
            }

            return null;
        }

        pub fn shift(self: *@This()) ?*Node {
            if (self.first) |first| {
                self.first = first.next;
                self.len -= 1;

                return first;
            }

            return null;
        }

        pub fn unshift(self: *@This(), node: *Node) void {
            if (self.first) |first| {
                first.prev = node;
                node.next = first;
            }

            self.first = node;
            self.len += 1;
        }

        pub fn delete(self: *@This(), node: *Node) void {
            var current = self.first;

            while (current) |currentNode| {
                if (currentNode == node) {
                    if (currentNode.next) |next| {
                        next.prev = currentNode.prev;
                    }
                    if (currentNode.prev) |prev| {
                        prev.next = currentNode.next;
                    }
                    if (self.first == currentNode) {
                        self.first = currentNode.next;
                    }
                    if (self.last == current) {
                        self.last = currentNode.prev;
                    }

                    self.len -= 1;

                    return;
                }

                current = currentNode.next;
            }
        }
    };
}
