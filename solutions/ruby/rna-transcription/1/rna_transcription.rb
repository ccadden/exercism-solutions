class Complement
  @@translation = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U'
  }

  def self.of_dna(strand)
    comp = ''
    strand.each_char { |chr|
      return '' unless @@translation.key?(chr.upcase)
      comp << @@translation[chr.upcase]
    }
    return comp
  end
end


module BookKeeping
  VERSION = 4 # Where the version number matches the one in the test.
end