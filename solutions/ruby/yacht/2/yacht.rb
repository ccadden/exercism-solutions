class Yacht
  attr_reader :dice, :category

  def initialize(dice, category)
    @dice = dice.sort
    @category = category
  end

  def score
    case category
    when 'ones' then dice.count(1)
    when 'twos' then dice.count(2) * 2
    when 'threes' then dice.count(3) * 3
    when 'fours' then dice.count(4) * 4
    when 'fives' then dice.count(5) * 5
    when 'sixes' then dice.count(6) * 6
    when 'yacht' then dice.uniq.length == 1 and 50 or 0
    when 'full house' then full_house? and dice.sum or 0
    when 'four of a kind' then (dice.tally.select { |_, v| v >= 4 }.keys.first or 0) * 4
    when 'little straight' then dice.eql?([1, 2, 3, 4, 5]) and 30 or 0
    when 'big straight' then dice.eql?([2, 3, 4, 5, 6]) and 30 or 0
    else dice.sum
    end
  end

  private

  def full_house?
    dice.uniq.length == 2 && !dice.tally.any? { |_, v| v == 4 }
  end
end