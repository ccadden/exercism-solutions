<?php

class LuckyNumbers
{
    public function sumUp(array $digitsOfNumber1, array $digitsOfNumber2): int
    {
        return (int) implode($digitsOfNumber1) + (int) implode($digitsOfNumber2);
    }

    public function isPalindrome(int $number): bool
    {
        return (int) strrev($number) == $number;
    }

    public function validate(string $input): string
    {
        if(empty($input) && $input !== '0') {
            return 'Required field';
        }

        if((int) $input <= 0) {
            return 'Must be a whole number larger than 0';
        } 

        return '';
    }
}
