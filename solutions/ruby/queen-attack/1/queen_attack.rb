class Queens
  attr_reader :white, :black
  
  def initialize(white:, black: [])
    @white = white
    @black = black
    raise ArgumentError.new() unless params_valid?
  end

  def params_valid?
    return false if white.any?{ |n| n < 0 || n > 7 }
    return false if black.any?{ |n| n < 0 || n > 7 }

    return true
  end

  def attack?
    return true if share_rank?
    return true if share_file?
    return true if share_diagonal?

    return false
  end

  private

  def share_file?
    white[0] == black[0]
  end

  def share_rank?
    white[1] == black[1]
  end

  def share_diagonal?
    (white[0] - black[0]).abs == (white[1] - black[1]).abs
  end
end