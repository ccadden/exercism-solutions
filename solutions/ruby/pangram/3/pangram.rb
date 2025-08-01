class Pangram
  def self.pangram?(phrase)
    letters = phrase.downcase.gsub(/[^a-z]/,'').split('').uniq.length
    letters == 26 ? true : false
  end
end

module BookKeeping
  VERSION=6
end