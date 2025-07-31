pub fn squareRoot(radicand: usize) usize {
    var i: usize = 1;
    while (i < radicand / 2) {
        if (i * i == radicand) {
            break;
        }
        i += 1;
    }

    return i;
}
