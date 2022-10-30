package main

import (
	"fmt"
	"testing"

	"birc.au.dk/gsa/shared"
)

func Test_Output(t *testing.T) {

	genome, read := "aka", "aka"

	sa := shared.LsdRadixSort(genome)
	fmt.Println(sa)
	for i, idx := range sa {
		fmt.Println(i, genome[idx:])
	}
	lower, upper := shared.BinarySearch(genome, read, sa)
	fmt.Println(lower, upper)
	fmt.Println("exact matches are in the interval:", lower, "to", upper)
	for _, v := range sa {
		idx := sa[v]

		//check if all suffixes in the interval matches and that all suffixes outside do not match.
		if v >= lower && v < upper {
			if genome[idx:len(read)+idx] != read {
				t.Error("They ARE NOT identical. But should be at idx:", v)
			}
		} else {
			if sa[v]+len(read) < len(genome) {
				if genome[idx:len(read)+idx] == read {
					t.Error("They ARE identical. But should be at idx:", v)
				}
			}
		}
	}

}

/*
func Test_cmp_with_old_handin(t *testing.T) {

	shared.SortFile("./testdata/output.txt")
	shared.SortFile("./testdata/handin3_reference.txt")

	if !shared.CmpFiles("./testdata/test_result.txt", "./testdata/h1_naive_results.txt") {
		t.Errorf("files are not identical")
	}
}
*/
