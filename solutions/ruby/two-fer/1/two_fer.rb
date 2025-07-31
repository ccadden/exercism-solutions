class TwoFer
  def self.two_fer(name=nil)
    fer = name.nil? ? 'you' : name

    "One for #{fer}, one for me."
  end
end