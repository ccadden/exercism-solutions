class Brackets
  def self.paired?(input)
    seen = []

    input.each_char do |char|
      case char
      when '(', '[', '{'
        seen << char
      when ')'
        return false if seen.pop != '('
      when ']'
        return false if seen.pop != '['
      when '}'
        return false if seen.pop != '{'
      end
    end

    seen.length == 0
  end
end
