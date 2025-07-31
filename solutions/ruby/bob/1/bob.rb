class Bob
  def self.hey(remark)
    remark.strip!
    return "Whoa, chill out!" if remark.upcase == remark && /.*[a-zA-Z]+.*/.match?(remark)
    return "Sure." if /^.*\?$\Z/.match?(remark)
    return "Fine. Be that way!" if remark.empty?
    "Whatever."
  end
end

module BookKeeping
  VERSION = 1
end