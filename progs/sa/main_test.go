package main

import (
	"fmt"
	"testing"

	"birc.au.dk/gsa/shared"
)

func Test_Output(t *testing.T) {

	genome, read := "abababbbbabaaaab", "aa"

	sa := shared.LsdRadixSort(genome)
	fmt.Println(sa)
	for i, idx := range sa {
		fmt.Println(i, genome[idx:])
	}
	lower, upper := BinarySuffixArraySearch(genome, read, sa)
	for _, v := range sa {
		idx := sa[v]
		if v >= lower && v < upper {
			if genome[idx:len(read)+idx] != read {
				t.Error("should be identical")
			}
		} else {
			if sa[v]+len(read) < len(genome) {
				if genome[idx:len(read)+idx] == read {
					t.Error("should not be identical")
				}
			}
		}
	}

}
