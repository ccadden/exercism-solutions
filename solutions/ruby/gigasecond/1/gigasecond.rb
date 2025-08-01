class Gigasecond
  def self.from(date)
    return Time.at(date.to_i + 1000000000)
  end
end

module BookKeeping
  VERSION = 6 # Where the version number matches the one in the test.
end