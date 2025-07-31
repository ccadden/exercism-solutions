class Hamming
  def self.compute(oldSeq, newSeq)
    if oldSeq.length != newSeq.length
      raise ArgumentError, 'sequences must be same length'
    end
    index = 0
    diff = 0
    oldSeq.each_char do |chr|
      diff += 1 if chr != newSeq[index]
      index += 1
    end
    return diff
  end
end

module BookKeeping
  VERSION = 3 # Where the version number matches the one in the test.
end