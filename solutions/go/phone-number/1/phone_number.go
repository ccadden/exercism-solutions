package phonenumber

import (
	"fmt"
	"regexp"
	"errors"
)

var nonDigits = regexp.MustCompile(`[^0-9]`)
var invalidPhoneNumber = errors.New("Invalid Phone Number")

func parseTenDigitNumber(digits string) (string, bool) {
	firstDigit := digits[0:1]
	fourthDigit := digits[3:4]
	switch {
	case firstDigit == "1":
		return "", false
	case firstDigit == "0":
		return "", false
	case fourthDigit == "1":
		return "", false
	case fourthDigit == "0":
		return "", false
	default:
		return digits, true
	}
}

func parseElevenDigitNumber(digits string) (string, bool) {
	if digits[0:1] != "1" {
		return "", false
	} else {
		return parseTenDigitNumber(digits[1:])
	}
}

func parsePhoneNumber(digits string) (string, bool) {
	length := len(digits)

	switch {
	case length > 11:
		return "", false
	case length < 10:
		return "", false
	case length == 10:
		return parseTenDigitNumber(digits)
	case length == 11:
		return parseElevenDigitNumber(digits)
	}

	if length == 11 {
	} else if length == 10 {
		digitSlice := digits[0:1]
		if digitSlice == "1" || digitSlice == "0" {
			return "", false
		} else {
			return digits, true
		}
	}

	return "", false
}

func Number(phoneNumber string) (string, error) {
	sanitized, valid := parsePhoneNumber(nonDigits.ReplaceAllString(phoneNumber, ""))

	if !valid {
		return "", invalidPhoneNumber
	}

	return sanitized, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	//`(613) 995-0253`

	number, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s",number[0:3], number[3:6], number[6:]), nil
}
