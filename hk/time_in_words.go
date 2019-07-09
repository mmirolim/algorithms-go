package hk

import "strings"

// https://www.hackerrank.com/challenges/the-time-in-words/problem
func TimeInWords(hour, minutes int32) string {
	oclock := "o' clock"
	mins := "minutes"
	min := "minute"
	quarter := "quarter"
	half := "half"
	past, to := "past", "to"
	digits := []string{"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}
	teens := []string{"ten", "eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen", "seventeen", "eighteen",
		"nineteen",
	}
	tyes := []string{"0", "ten", "twenty"}
	pastOrTo := past
	if minutes > 30 {
		pastOrTo = to
		minutes = 60 - minutes
		hour++
	}
	tens, digit := getDigits(minutes)
	out := []string{}
	if tens == 0 && digit != 0 {
		if digit == 1 {
			out = []string{digits[digit], min, pastOrTo}
		} else {
			out = []string{digits[digit], mins, pastOrTo}
		}
	} else if tens == 1 {
		if digit == 5 {
			out = []string{quarter, pastOrTo}
		} else {
			out = []string{teens[digit], mins, pastOrTo}
		}
	} else if tens >= 2 {
		if tens == 3 && digit == 0 {
			out = []string{half, past}
		} else {
			out = []string{tyes[tens], digits[digit], mins, pastOrTo}
		}
	}
	tens, digit = getDigits(hour)
	var hours []string
	if tens == 0 {
		hours = []string{digits[digit]}
	} else {
		hours = []string{teens[digit]}
	}

	if len(out) == 0 {
		hours = append(hours, oclock)
	}
	out = append(out, hours...)
	return strings.Join(out, " ")
}

func getDigits(n int32) (int, int) {
	if n > 60 || n < 0 {
		panic("invalid number")
	}
	return int(n / 10), int(n % 10)
}
