var Hamming = function() {

}

Hamming.prototype.compute = function(strandA, strandB) {
  // Check that strands are equal in length
  if (strandA.length != strandB.length) {
    throw new Error('DNA strands must be of equal length.')
  }
  else {
    hamming = 0
    for(var i = 0; i < strandA.length; i++) {
      if(strandA[i] !== strandB[i]) {
        hamming ++
      }
    }
  }
  return hamming
}

module.exports = Hamming