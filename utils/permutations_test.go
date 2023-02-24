package utils

import (
	"testing"
)

func TestDoForPermutations(t *testing.T) {
	vals := []string{"a", "b", "c"}
	found := 0
	DoForElementPermutations(vals, func(i int, p []string) {
		found++
		// fmt.Println(p)
	})
	if found != 6 {
		t.Fatal("expected 6 permutations, found", found)
	}
	if vals[0] != "a" {
		t.Fatal("expected vals[0] to be a, found", vals[0])
	}
}
