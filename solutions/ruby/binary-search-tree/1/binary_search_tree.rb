=begin
Write your code for the 'Binary Search Tree' exercise in this file. Make the tests in
`binary_search_tree_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/binary-search-tree` directory.
=end

class Bst
  attr_reader :data, :left, :right
  
  def initialize(data)
    @data = data
    @left = nil
    @right = nil
  end

  def insert(val)
    if val <= data
      if !left.nil?
        @left.insert(val)
      else
        @left = Bst.new(val)
      end
    else
      if !right.nil?
        @right.insert(val)
      else
        @right = Bst.new(val)
      end
    end
  end

  def each(&block)
    return to_enum(:each) unless block_given?
    
    left&.each(&block)

    yield data

    right&.each(&block)
  end
end