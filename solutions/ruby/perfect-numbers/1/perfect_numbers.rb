class PerfectNumber
  def self.classify(number)
    raise RuntimeError if number < 1

    sum = (1..number/2).sum do |n|
      (number % n).zero? ? n : 0
    end

    case
      when sum < number then 'deficient'
      when sum > number then 'abundant'
      else 'perfect'
    end
  end
end