package main

import (
	"fmt"
	"testing"

	"birc.au.dk/gsa/shared"
)

func Test_Output(t *testing.T) {

	genome, read := "akasakan$", "aka"

	sa := shared.LsdRadixSort(genome)
	fmt.Println(sa)
	for i, idx := range sa {
		fmt.Println(i, genome[idx:])
	}
	lower, upper := shared.BinarySearch(genome, read, sa)
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

func TestVaryingAlphabets(t *testing.T) {
	matches := 0
	Alphabets := []shared.Alphabet{
		shared.English, shared.DNA, shared.AB}

	for _, v := range Alphabets {
		genome, reads := shared.BuildSomeFastaAndFastq(300, 8, 20, v, 11)

		parsedGenomes := shared.GeneralParserStub(genome, shared.Fasta, len(genome)+1)
		parsedReads := shared.GeneralParserStub(reads, shared.Fastq, len(reads)+1)

		//iterate all genomes
		for _, gen := range parsedGenomes {
			sa := shared.LsdRadixSort(gen.Rec)
			//iterate all reads
			for _, read := range parsedReads {
				lower, upper := shared.BinarySearch(gen.Rec, read.Rec, sa)
				matches += upper - lower

				//verify that the bounds are correct
				for _, v := range sa {
					idx := sa[v]

					//check if all suffixes in the interval matches and that all suffixes outside do not match.
					if v >= lower && v < upper {
						if gen.Rec[idx:len(read.Rec)+idx] != read.Rec {
							t.Error("They ARE NOT identical. But should be at idx:", v)
						}
					} else {
						if sa[v]+len(read.Rec) < len(gen.Rec) {
							if gen.Rec[idx:len(read.Rec)+idx] == read.Rec {
								t.Error("They ARE identical. But should be at idx:", v)
							}
						}
					}
				}
			}
		}

	}
	fmt.Println("a total of", matches, " matches was found in the test.")
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
