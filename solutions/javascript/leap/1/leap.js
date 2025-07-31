//
// This is only a SKELETON file for the "Leap" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Year = function(input) {
//
// YOUR CODE GOES HERE
//
this.year = input
};

Year.prototype.isLeap = function() {
//
// YOUR CODE GOES HERE
//
  if(this.year % 4){
    // Not evenly divisible by 4
    return false
  }
  else {
    if(this.year % 100){
      // Not divisible by 100
      return true
    }
    else {
      if(this.year % 400) {
        // Divisible by 100 but not 400
        return false
      }
      else {
        return true
      }
    }
  }
};

module.exports = Year;