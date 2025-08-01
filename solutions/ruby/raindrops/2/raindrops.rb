class Raindrops
  SOUNDS = { 3 => 'Pling', 5 => 'Plang', 7 => 'Plong'}.freeze
  
  def self.convert(number)
    result = SOUNDS.select{ |key, val| (number % key).zero? }.values.join
    result.empty? ? number.to_s : result
  end
end
