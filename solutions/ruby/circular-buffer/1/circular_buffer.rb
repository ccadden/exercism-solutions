class CircularBuffer
  attr_reader :buffer, :len

  def initialize(len)
    @buffer = []
    @len = len
  end

  def read
    raise BufferEmptyException if buffer.empty?

    @buffer.shift
  end

  def write(val)
    raise BufferFullException if buffer_full?

    @buffer.push(val)
  end

  def write!(val)
    @buffer.shift if buffer_full?
    @buffer.push(val)
  end

  def clear
    @buffer = []
  end

  private

  def buffer_full?
    buffer.length == len
  end

  class BufferEmptyException < StandardError
  end

  class BufferFullException < StandardError
  end
end
