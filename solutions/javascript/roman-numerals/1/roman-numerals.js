const NUMERALS = {
  1000: "M",
  900: "CM",
  500: "D",
  400: "CD",
  100: "C",
  90: "XC",
  50: "L",
  40: "XL",
  10: "X",
  9: "IX",
  5: "V",
  4: "IV",
  1: "I",
};

const FACTORS = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

export const toRoman = (num) => {
  var output = "";

  FACTORS.forEach((factor) => {
    if (!(num < factor)) {
      var multiple = Math.floor(num / factor);
      var remainder = num % factor;

      output += NUMERALS[factor].repeat(multiple);

      num = remainder;
    }
  });

  return output;
};
