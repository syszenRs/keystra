package prep

import (
	"fmt"
	"strings"
)

func pop[T any](a []T, k int) []T {
	return a[:len(a)-k]
}

func RemoveAdjacent(s string) string {
	sArray := make([]string, len(s))
	for i, r := range s {
		char := string(r)
		if i > 0 && sArray[len(sArray)-1] == char {
			sArray = pop(sArray, 1)
		} else {
			sArray = append(sArray, char)
		}

		fmt.Printf("char - %s - %d \n", char, i)
	}

	return strings.Join(sArray, "")
}

func RemoveKAdjacent(s string, k int) string {
	type helper struct {
		s string
		c int
	}

	sArray := make([]helper, len(s))

	for _, r := range s {
		char := string(r)

		if len(sArray) == 0 {
			sArray = append(sArray, helper{s: char, c: 1})
		} else {
			last := sArray[len(sArray)-1]
			if last.s == char {
				if last.c == 2 {
					sArray = pop(sArray, 2)
				} else {
					sArray = append(sArray, helper{s: char, c: last.c + 1})
				}
			} else {
				sArray = append(sArray, helper{s: char, c: 1})
			}
		}
	}

	out := ""

	for _, s := range sArray {
		out += s.s
	}

	return out
}
