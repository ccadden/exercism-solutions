var Gigasecond = function(birthday) {
  this.birthday = birthday
}

Gigasecond.prototype.date = function() {
  return new Date(this.birthday.getTime() + Math.pow(10,9)*1000) // adds 10^9 seconds to the birthday time
}

module.exports = Gigasecond