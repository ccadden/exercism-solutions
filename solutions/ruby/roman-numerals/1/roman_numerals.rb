class Integer
  def to_roman
    romans = {
      1 => "I",
      4 => "IV",
      5 => "V",
      9 => "IX",
      10 => "X",
      40 => "XL",
      50 => "L",
      90 => "XC",
      100 => "C",
      400 => "CD",
      500 => "D",
      900 => "CM",
      1000 => "M"
    }
    int = self
    numeral = ""
    factors = romans.keys.sort { |x, y| y <=> x }
    factors.each do |factor|
      next if int < factor
      multiple = int / factor
      remainder = int % factor
      numeral << romans[factor] * multiple
      int = remainder
    end
    numeral
  end
end

module BookKeeping
  VERSION = 2
end