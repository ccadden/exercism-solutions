class RotationalCipher
  def self.rotate(text, rotation)
    rotated = ''

    text.each_char do |char|
      rotated << case char
                 when 'A'..'Z'
                   transform(char, rotation, 'A')
                 when 'a'..'z'
                   transform(char, rotation, 'a')
                 else
                   char
                 end
    end

    rotated
  end

  def self.transform(char, rotation, start)
    char = char.ord
    start = start.ord

    (((char - start + rotation) % 26) + start).chr
  end

  private_class_method :transform
end
