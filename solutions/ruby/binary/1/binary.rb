class Binary
  def self.to_decimal(binary_string)
    raise ArgumentError unless /\A[01]+\z/.match?(binary_string)
    # reverse the string and then add to a stack that power of two
    string_array = binary_string.reverse.chars
    stack = 0
    string_array.each_index do |index|
      stack = stack + string_array[index].to_i * 2 ** index
    end
    stack
  end
end

module BookKeeping
  VERSION = 3 # Where the version number matches the one in the test.
end