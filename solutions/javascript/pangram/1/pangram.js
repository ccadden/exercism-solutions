var Pangram = function(phrase) {
  this.phrase = phrase
}
Pangram.prototype.isPangram = function() {
  alphabet = 'abcdefghijklmnopqrstuvwxyz'
  for(var i = 0; i < this.phrase.length; i++) {
    alphabet = alphabet.replace(this.phrase[i].toLowerCase(),'')
  }
  pangram = alphabet.length === 0 ? true : false
  return pangram
}
module.exports = Pangram