class Prime
  def self.nth(num)
    if num < 1
      raise(ArgumentError)
    end
    # initialize 2 as first prime
    primes = [2]
    p = 3
    while primes.length < num do
      factors = primes.any? { |prime| p % prime == 0 }
      primes << p if !factors
      p = p + 2
    end
    primes[-1]
  end
end

module BookKeeping
  VERSION = 1
end