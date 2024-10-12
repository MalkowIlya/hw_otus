package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result []rune
	var lastRune rune
	var isDigit bool
	var nextDefault bool

	for _, r := range str {
		switch {
		case unicode.IsDigit(r) && !nextDefault:
			if isDigit {
				return "", ErrInvalidString
			}

			if lastRune == 0 {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(r))
			if count == 0 {
				result = result[0 : len(result)-1]
			} else {
				for i := 0; i < count-1; i++ {
					result = append(result, lastRune)
				}
			}

			isDigit = true
		case string(r) == `\` && !nextDefault:
			nextDefault = true
		default:
			result = append(result, r)
			lastRune = r
			isDigit = false
			nextDefault = false
		}
	}

	return string(result), nil
}
