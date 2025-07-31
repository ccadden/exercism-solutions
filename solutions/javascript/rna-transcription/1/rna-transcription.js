var DnaTranscriber = function() {

}

DnaTranscriber.prototype.toRna = function(strand) {
  var map = {
    'G': 'C',
    'C': 'G',
    'T': 'A',
    'A': 'U'
  }

  rna = ''
  for(var i = 0; i < strand.length; i++) {
    if(!(strand[i] in map)) {
      throw new Error('Invalid input')
    }
    rna += map[strand[i]]
  }
  return rna
}

module.exports = DnaTranscriber