class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    m = /:\s+(.*)$/.match(@line)
    return m[1].strip if !m.nil?
  end

  def log_level
    m = /\[(\w+)\]/.match(@line)
    return m[1].strip.downcase if !m.nil?
  end

  def reformat
    m = /\[(\w+)\]: (.*)/.match(@line)
    return (m[2].strip + " (" + m[1].downcase.strip + ")") if !m.nil?
  end
end
