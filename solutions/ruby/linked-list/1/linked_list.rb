# Write your code for the 'Linked List' exercise in this file. Make the tests in
# `linked_list_test.rb` pass.
#
# To get started with TDD, see the `README.md` file in your
# `ruby/linked-list` directory.
class Node
  attr_accessor :value, :next, :prev

  def initialize(value)
    @value = value
    @next = nil
    @prev = nil
  end
end

class Deque
  def initialize
    @head = nil
    @tail = nil
    @length = 0
  end

  def push(value)
    new_node = Node.new(value)

    unless @tail.nil?
      @tail.next = new_node
      new_node.prev = @tail
    end

    @head = new_node if @head.nil?

    @length += 1

    @tail = new_node
  end

  def pop
    unless @tail.nil?
      dead_node = @tail
      @tail = @tail.prev

      @length -= 1

      @head = nil if @length.zero?

      return dead_node.value
    end

    nil
  end

  def shift
    unless @head.nil?
      dead_node = @head
      @head = @head.next
      @length -= 1

      @tail = nil if @length.zero?

      return dead_node.value
    end

    nil
  end

  def unshift(value)
    new_node = Node.new(value)

    unless @head.nil?
      @head.prev = new_node
      new_node.next = @head
    end

    @head = new_node

    @length += 1

    @tail = new_node if @tail.nil?
  end
end
