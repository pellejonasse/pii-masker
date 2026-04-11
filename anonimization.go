package piimasker

import (
	"strconv"
	"fmt"
	"math/rand/v2"	
)

// not sure about this, but it might be nice to actually preserve the size of numbers
func preserveNumberSize[T Number](v T) T {
	s := fmt.Sprintf("%v", v)
	b := []byte(s)
	for i, c := range b {
		if c >= '0' && c <= '9' {
			b[i] = '9'
		}
	}
	var result T
	fmt.Sscan(string(b), &result)
	return result
}


func anonymizeString(s string) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, len(s))
	for i := range b {
		b[i] = chars[rand.IntN(len(chars))]
	}
	return string(b)
}

func anonymizeInt(v int64) int64 {
	s := strconv.FormatInt(v, 10)
	b := []byte(s)
	for i, c := range b {
		if c >= '0' && c <= '9' {
			if i == 0 || (i == 1 && b[0] == '-') {
				b[i] = byte('1' + rand.IntN(9))
			} else {
				b[i] = byte('0' + rand.IntN(10))
			}
		}
	}
	result, _ := strconv.ParseInt(string(b), 10, 64)
	return result
}

func anonymizeUint(v uint64) uint64 {
	s := strconv.FormatUint(v, 10)
	b := []byte(s)
	for i, c := range b {
		if c >= '0' && c <= '9' {
			if i == 0 {
				b[i] = byte('1' + rand.IntN(9))
			} else {
				b[i] = byte('0' + rand.IntN(10))
			}
		}
	}
	result, _ := strconv.ParseUint(string(b), 10, 64)
	return result
}

func anonymizeFloat(v float64) float64 {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	b := []byte(s)
	firstDigit := true
	for i, c := range b {
		if c >= '0' && c <= '9' {
			if firstDigit {
				b[i] = byte('1' + rand.IntN(9))
				firstDigit = false
			} else {
				b[i] = byte('0' + rand.IntN(10))
			}
		}
	}
	result, _ := strconv.ParseFloat(string(b), 64)
	return result
}
