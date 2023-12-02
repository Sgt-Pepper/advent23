package day01

import (
	"log/slog"
	"strconv"
	"strings"
	"unicode"
)

func EvalDocument(s string) int {
	var total int = 0
	for _, line := range strings.Split(string(s), "\n") {
		line = replaceStringsForPart2(line) // uncomment for part 1
		first, last := getDigits(line)
		tens, _ := strconv.Atoi(string(first))
		ones, _ := strconv.Atoi(string(last))
		slog.Debug("-------->", "result", (10*tens)+ones)
		total += (10 * tens) + ones
	}
	return total
}

func getDigits(s string) (rune, rune) {
	var first, last, zero rune
	for _, r := range s {
		if unicode.IsDigit(r) {
			if first == zero {
				first = r
			}
			last = r
		}
	}
	return first, last
}

func replaceStringsForPart2(s string) string {
	slog.Debug("o", "string", s)
	s = strings.ReplaceAll(s, "one", "one1one")
	s = strings.ReplaceAll(s, "two", "two2two")
	s = strings.ReplaceAll(s, "three", "three3three")
	s = strings.ReplaceAll(s, "four", "four4four")
	s = strings.ReplaceAll(s, "five", "five5five")
	s = strings.ReplaceAll(s, "six", "six6six")
	s = strings.ReplaceAll(s, "seven", "seven7seven")
	s = strings.ReplaceAll(s, "eight", "eight8eight")
	s = strings.ReplaceAll(s, "nine", "nine9nine")
	//	s = strings.ReplaceAll(s, "zero", "0")
	slog.Debug("-------->", "string", s)
	return s

}
