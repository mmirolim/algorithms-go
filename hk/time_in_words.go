package hk

import "strings"

// https://www.hackerrank.com/challenges/the-time-in-words/problem
func TimeInWords(hour, minutes int32) string {
	oclock := "o' clock"
	mins := "minutes"
	min := "minute"
	past, to := "past", "to"
	numWords := []string{"", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
		"fourteen", "quarter", "sixteen", "seventeen", "eighteen",
		"nineteen", "twenty",
	}
	tyes := []string{"", "ten", "twenty", "half"}
	if minutes == 0 {
		return numWords[hour] + " " + oclock
	}
	pastOrTo := past
	if minutes > 30 {
		pastOrTo = to
		minutes = 60 - minutes
		hour++
	}
	var out []string
	if minutes == 1 {
		out = []string{numWords[minutes], min}
	} else if minutes == 15 {
		out = []string{numWords[minutes]}
	} else if minutes == 30 {
		out = []string{tyes[minutes/10]}
	} else if minutes <= 20 {
		out = []string{numWords[minutes], mins}
	} else {
		out = []string{tyes[minutes/10] + " " + numWords[minutes%10], mins}
	}

	out = append(out, []string{pastOrTo, numWords[hour]}...)
	return strings.Join(out, " ")
}

func getDigits(n int32) (int, int) {
	if n > 60 || n < 0 {
		panic("invalid number")
	}
	return int(n / 10), int(n % 10)
}
