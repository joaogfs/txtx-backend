package main

import (
	"regexp"
)

func replace(text string, pat string, replacement string) string {
	return string(regexp.MustCompile(pat).ReplaceAll([]byte(text), []byte(replacement)))
}
