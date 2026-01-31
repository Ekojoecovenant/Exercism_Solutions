//
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const format = (name, number) => {
  let term = "";
  const numbeg = number % 100
  const numLast = number % 10
  console.log(numLast)

  if (numbeg >= 11 && numbeg <= 13) {
    term = "th";
  } else if (numLast === 1) {
    term = "st"; 
  } else if (numLast === 2) {
    term = "nd";
  } else if (numLast === 3) {
    term = "rd";
  } else {
    term = "th";
  }

  const message = `${name}, you are the ${number}${term} customer we serve today. Thank you!`;
  return message;
};
