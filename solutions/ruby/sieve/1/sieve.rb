class Sieve
  def initialize(limit)
    @limit = limit
  end

  def primes
    return [] if @limit < 2
    primes = []
    multiples = []
    (2..@limit).each do |x|
      next if multiples.include?(x)
      primes << x
      iter = x
      while iter <= @limit do
        multiples << iter
        iter = iter + x
      end
    end
    primes
  end
end

module BookKeeping
  VERSION = 1
end