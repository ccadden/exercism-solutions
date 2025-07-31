class RunLengthEncoding
  def self.encode(input)
    return "" if input.empty?
    output = ""
    sliced = input
    pos = 0
    while true do
      # find a match
      m = /(\D)\1*/.match(sliced)
      break if m.nil?
      # find starting and ending indicies to determine repetition
      start = m.begin(0)
      finish = m.end(0)
      run = (finish - start).to_s
      run = "" if run == "1"
      output << run + m[0][0] 
      pos = pos + finish
      # create new substring to operate on
      sliced = input[pos..-1]
    end
    output
  end

  def self.decode(input)
    return "" if input.empty?
    output = ""
    slice = input
    pos = 0
    while true do
      # find match
      m = /(\d+)?(\D){1}/.match(slice)
      break if m.nil?
      # get capture groups
      captures = m.captures 
      repeat = captures[0]
      repeat = "1" if repeat.nil?
      output << captures[1] * repeat.to_i
      # reslice
      pos = pos + m.end(0)
      slice = input[pos..-1]
    end
    output
  end
end

module BookKeeping
  VERSION = 3
end