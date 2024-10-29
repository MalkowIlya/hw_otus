package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res strings.Builder
	var lastChar rune

	for i, ch := range str {
		if unicode.IsDigit(ch) {
			count, err := strconv.Atoi(string(ch))

			if err != nil || i == 0 || unicode.IsDigit(lastChar) {
				return "", ErrInvalidString
			}

			if count == 0 && res.Len() > 0 {
				resStr := res.String()
				res.Reset()
				res.WriteString(resStr[:len(resStr)-1])
			} else if count > 0 {
				res.WriteString(strings.Repeat(string(lastChar), count-1))
			}
			lastChar = ch
			continue
		}

		res.WriteRune(ch)
		lastChar = ch
	}

	return res.String(), nil
}
