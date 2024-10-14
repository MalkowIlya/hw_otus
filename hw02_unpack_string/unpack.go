package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var repeat = func(writeString string, repeatCount int, result *strings.Builder) {
	if repeatCount > 0 {
		result.WriteString(strings.Repeat(writeString, repeatCount))
	}
}

func Unpack(str string) (string, error) {
	var result strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if i == 0 && unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		}

		if runes[i] == '\\' {
			if err := escapedChar(runes, &i, &result); err != nil {
				return "", err
			}
		} else {
			if err := simpleChar(runes, &i, &result); err != nil {
				return "", err
			}
		}
	}

	return result.String(), nil
}

func escapeItem(runes []rune, i *int, result *strings.Builder, escapedChar string) error {
	if *i+2 < len(runes) && unicode.IsDigit(runes[*i+2]) {
		repeatCount, _ := strconv.Atoi(string(runes[*i+2]))
		repeat(escapedChar, repeatCount, result)
		*i += 2
	} else {
		result.WriteString(escapedChar)
		*i++
	}
	return nil
}

func escapedChar(runes []rune, i *int, result *strings.Builder) error {
	literals := map[rune]string{'n': `\n`, 't': `\t`, 'r': `\r`, 's': `\s`}
	next := runes[*i+1]

	if literal, ok := literals[next]; ok {
		return escapeItem(runes, i, result, literal)
	}

	if next == '\\' {
		return escapeItem(runes, i, result, `\`)
	}

	if unicode.IsDigit(next) {
		return escapeItem(runes, i, result, string(next))
	}

	return ErrInvalidString
}

func simpleChar(runes []rune, i *int, result *strings.Builder) error {
	current := runes[*i]

	switch {
	case *i+1 < len(runes) && unicode.IsDigit(runes[*i+1]):
		repeatCount, _ := strconv.Atoi(string(runes[*i+1]))
		if repeatCount > 0 {
			repeat(string(current), repeatCount, result)
		}
		*i++
	case unicode.IsDigit(current):
		return ErrInvalidString
	default:
		result.WriteRune(current)
	}

	return nil
}
