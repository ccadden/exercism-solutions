class Raindrops
  def self.convert(number)
    noise = ''
    noise << 'Pling' if number % 3 == 0
    noise << 'Plang' if number % 5 == 0
    noise << 'Plong' if number % 7 == 0
    return number.to_s if noise == ''
    return noise
  end
end

module BookKeeping
  VERSION = 3 # Where the version number matches the one in the test.
end