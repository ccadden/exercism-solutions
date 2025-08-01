class RailFenceCipher
  def self.encode(message, rails)
    return message unless valid_input?(message, rails)
    
    zig_zag(rails, message.length).
      zip(message.chars).
      sort.
      map(&:last).
      join
  end

  def self.decode(message, rails)
    return message unless valid_input?(message, rails)
    
    zig_zag(rails, message.length).
      sort.
      zip(message.chars).
      sort_by{ |c| c.first.last }.
      map(&:last).
      join
  end

  private

  def self.valid_input?(message, rails)
    !(message.empty? || rails < 2)
  end

  def self.zig_zag(rails, length)
    ((0...rails).to_a + (1...rails - 1).to_a.reverse).
      cycle.
      first(length).
      zip(0...length)
  end
end
