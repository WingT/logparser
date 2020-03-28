package logparser

import (
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

type Entry map[string]string

func ParseTextFormattedLog(log string) Entry {
	var s scanner.Scanner
	s.Init(strings.NewReader(log))
	s.Filename = "default"
	s.IsIdentRune = func(ch rune, i int) bool {
		// https://github.com/sirupsen/logrus/blob/v1.0.5/text_formatter.go#L143-L156
		return ch == '-' || ch == '.' || ch == '/' || ch == '@' || ch == '^' || ch == '+' || ch == '_' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}

	entry := make(Entry)
	var key string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tt := s.TokenText()
		if len(tt) > 0 && tt[0] == '"' {
			tt, _ = strconv.Unquote(tt)
		}
		if tt == "=" {
			continue
		}
		if key == "" {
			key = tt
		} else {
			entry[key] = tt
			key = ""
		}
	}
	return entry
}
