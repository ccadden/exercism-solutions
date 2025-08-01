// Please implement the `ComputationError.IllegalArgument` error.
pub const ComputationError = error{
    IllegalArgument,
};

pub fn steps(number: usize) anyerror!usize {
    if (number < 1) {
        return ComputationError.IllegalArgument;
    }

    var n = number;
    var numSteps: usize = 0;

    while (n != 1) {
        if (n % 2 == 0) {
            n /= 2;
        } else {
            n = (3 * n) + 1;
        }

        numSteps += 1;
    }

    return numSteps;
}
