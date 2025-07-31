class SecretHandshake
  COMMANDS = ['wink', 'double blink', 'close your eyes', 'jump']
  
  def initialize(number)
    @binary_as_array = number.is_a?(Integer) ? ('%05b' % number).split('').reverse : []
  end

  def commands
    result = []
    
    @binary_as_array.slice(0,4).each_with_index do |c, i|
      result.push(COMMANDS[i]) if c == '1'
    end

    result.reverse! if @binary_as_array[-1] == '1'

    result
  end
end