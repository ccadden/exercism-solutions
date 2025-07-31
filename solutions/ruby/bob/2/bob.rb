class Bob
  def self.hey(remark)
    return 'Fine. Be that way!' if blank?(remark)

    all_caps = scream_case?(remark)

    if question?(remark)
      return 'Sure.' unless all_caps

      return "Calm down, I know what I'm doing!"
    end

    return 'Whoa, chill out!' if all_caps

    'Whatever.'
  end

  def self.question?(str)
    str.rstrip!
    str.end_with?('?')
  end

  def self.scream_case?(str)
    has_letter = false
    str.each_char do |c|
      lower = ('a'..'z').include?(c)
      upper = ('A'..'Z').include?(c)

      next unless lower || upper
      return false if lower

      has_letter = true
    end

    has_letter && true
  end

  def self.blank?(str)
    whitespace_chars = [' ', "\t", "\n", "\r", "\f", "\v"]
    str.each_char.all? { |char| whitespace_chars.include?(char) }
  end

  private_class_method :question?, :scream_case?, :blank?
end
