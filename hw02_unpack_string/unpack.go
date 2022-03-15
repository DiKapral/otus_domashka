package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	r := []rune(s)
	if s == "" {
		return "", nil
	}
	if unicode.IsDigit(r[0]) {
		return "", ErrInvalidString
	}
	var v1 string
	var v2 string
	for _, simvl := range r {

		if !unicode.IsDigit(simvl) {
			v2 = string(simvl)
			if v2 == "\n" {
				v2 = "\\n"
			}
			v1 = builder(v1, v2)
		} else {
			sh, _ := strconv.Atoi(string(simvl))
			if v2 == "" {
				return "", ErrInvalidString
			}
			if sh == 0 {
				v1 = v1[:len(v1)-1]
				continue
			}
			v1 = builder(v1, strings.Repeat(v2, sh-1))
			v2 = ""
		}
	}
	return v1, nil
}

func builder(str1, str2 string) string {
	var builder strings.Builder
	builder.Grow(len(str1) + len(str2))
	builder.WriteString(str1)
	builder.WriteString(str2)
	return builder.String()
}
