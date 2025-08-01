module Die
  def self.roll(roll: 4, die: 6, remove_lowest: 1)
    roll.times.map { rand(1..die) }.sort.pop(roll - remove_lowest).sum
  end
end

class DndCharacter
  BASE_HITPOINTS = 10
  private_constant :BASE_HITPOINTS

  attr_reader :strength,
    :dexterity,
    :constitution,
    :intelligence,
    :wisdom,
    :charisma,
    :hitpoints

  def initialize
    @strength     = Die.roll
    @dexterity    = Die.roll
    @constitution = Die.roll
    @intelligence = Die.roll
    @wisdom       = Die.roll
    @charisma     = Die.roll

    @hitpoints    = BASE_HITPOINTS + self.class.modifier(constitution)
  end

  def self.modifier(ability_score)
    ability_score / 2 - 5
  end
end