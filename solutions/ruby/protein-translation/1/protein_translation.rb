class InvalidCodonError < StandardError
end

class Translation
  def self.of_rna(strand)
    res = []
    
    (0...strand.length).step(3).each do |i|
      codon = strand[i...i+3]
      p = protien(codon)
      
      break if p === 'STOP'
      
      res << protien(codon)
    end

    res
  end

  private_class_method def self.protien(codon)
    case codon
      when 'AUG' then 'Methionine'
      when 'UUU', 'UUC' then 'Phenylalanine'
      when 'UUA', 'UUG' then 'Leucine'
      when 'UCU', 'UCC', 'UCA', 'UCG' then 'Serine'
      when 'UAU', 'UAC' then 'Tyrosine'
      when 'UGU', 'UGC' then 'Cysteine'
      when 'UGG' then 'Tryptophan'
      when 'UAA', 'UAG', 'UGA' then 'STOP'
      else raise InvalidCodonError
    end
  end
end