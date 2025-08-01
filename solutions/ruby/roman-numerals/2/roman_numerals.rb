class Integer
  def to_roman
    romans = {
      1000 => "M",
      900 => "CM",
      500 => "D",
      400 => "CD",
      100 => "C",
      90 => "XC",
      50 => "L",
      40 => "XL",
      10 => "X",
      9 => "IX",
      5 => "V",
      4 => "IV",
      1 => "I"
    }
    int = self
    numeral = ""
    factors = romans.keys
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