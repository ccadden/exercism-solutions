class Element
  attr_reader :datum
  
  def initialize(datum)
    @datum = datum
    @next = nil
  end

  def next
    @next
  end

  def next=(el)
      @next = el
  end
end

class SimpleLinkedList
  attr_reader :list
  def initialize(list=[])
    @list = make_list(list)
  end

  def push(el)
    @list.unshift(el)
    self
  end

  def pop()
    @list.shift()
  end

  def to_a
    list.map(&:datum)
  end

  def reverse!
    @list.reverse!
    self
  end

  private
  
  def make_list(list)
    if list.kind_of?(Range)
      list.to_a.reverse.map { |n| Element.new(n) }
    elsif list.find { |n| n.kind_of?(Numeric) }
      list.reverse.map{ |n| Element.new(n) }
    else
      list.reverse
    end
  end
end