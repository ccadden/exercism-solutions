class Grains
  def self.square(num)
    if num < 1 || num > 64
      raise(ArgumentError)
    end
    2 ** (num - 1)
  end

  def self.total
    # Geometric series convergence 
    (1 - 2**64) / (1 - 2)
  end
end

module BookKeeping
  VERSION = 1
end