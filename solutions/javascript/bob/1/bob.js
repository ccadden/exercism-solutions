var Bob = function() {

}
Bob.prototype.hey = function(phrase) {
  // Question
  if(phrase.toUpperCase() == phrase && phrase.match(/.*[a-zA-Z]+.*/g)) {
    return 'Whoa, chill out!'
  }
  else if(phrase.match(/.*\?$/g)) {
    return 'Sure.'
  }
  else if(phrase.match(/^ *$/g)) {
    return 'Fine. Be that way!'
  }
  else {
    return 'Whatever.'
  }
}

module.exports = Bob