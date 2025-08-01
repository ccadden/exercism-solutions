class Squares
  def initialize(num)
    @num = num
  end

  def square_of_sum
    sum = 0
    (1..@num).each do |x|
      sum = sum + x
    end
    sum * sum
  end

  def sum_of_squares
    sum = 0
    (1..@num).each do |x|
      sum = sum + ( x * x )
    end
    sum
  end

  def difference
    square_of_sum - sum_of_squares 
  end
end

module BookKeeping
  VERSION = 4
end