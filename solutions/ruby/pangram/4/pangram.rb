class Pangram
  def self.pangram?(phrase)
    letters = phrase.downcase.gsub(/[^a-z]/,'').chars.uniq.length
    letters == 26
  end
end

module BookKeeping
  VERSION=6
end