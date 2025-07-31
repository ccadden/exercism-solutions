class Brackets
  def self.paired?(input)
    seen = ''
    input.each_char do |char|
      seen += 'p' if char == '('
      seen += 's' if char == '['
      seen += 'c' if char == '{'

      if char == ')'
        return false unless seen[-1] == 'p'

        seen = seen[0..-2]
      end

      if char == ']'
        return false unless seen[-1] == 's'

        seen = seen[0..-2]
      end

      next unless char == '}'
      return false unless seen[-1] == 'c'

      seen = seen[0..-2]
    end

    seen.length == 0
  end
end

