class SumOfMultiples
  attr_reader :level, :base_values
  
  def initialize(*base_values)
    @base_values = base_values
  end

  def to(level)
    unique_multiples = Set.new
    
    base_values.each do |value|
      next if value.zero?
      
      level.times do |n|
        unique_multiples.add(n) if (n % value).zero?
      end
    end

    unique_multiples.sum
  end
end