pub fn is_armstrong_number(num: u32) -> bool {
    let len_nums: u32 = num.to_string().len() as u32;

    return num as u64
        == num.to_string().chars().fold(0u64, |sum, val| {
            sum + (val.to_digit(10).unwrap() as u64).pow(len_nums)
        });
}
