/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    if code.trim().len() < 2 {
        return false;
    }

    let nums = code
        .chars()
        .filter(|c| *c != ' ')
        .map(|c| c.to_digit(10))
        .rev()
        .enumerate();

    let mut sum = 0;

    for (i, num) in nums {
        match num {
            None => return false,
            Some(mut val) => {
                if i % 2 != 0 {
                    val = val * 2;

                    if val > 9 {
                        val = val - 9;
                    }
                }

                sum += val
            }
        }
    }

    return sum % 10 == 0;
}
