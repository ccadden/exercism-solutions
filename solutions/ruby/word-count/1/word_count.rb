class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    count_hash = Hash.new(0)
    words = @phrase.scan(/[a-zA-Z0-9]+'?[a-zA-Z]*\b/)
    words.each do |word|
      count_hash[word.downcase] = count_hash[word.downcase] + 1
    end
    count_hash
  end
end

module BookKeeping
  VERSION = 1
end