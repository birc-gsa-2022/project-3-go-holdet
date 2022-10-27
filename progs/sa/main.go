package main

import (
	"fmt"
	"os"
	"shared"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genome := os.Args[1]
	reads := os.Args[2]

	suffixArray := shared.LsdRadixSort(genome)

	for read := range reads {

	}

	fmt.Println("Search in", genome, "for the reads in", reads)
}

func upperBound(match int, high int, suffixArray []int, genome string, read string) int {
	mid := high - (high-match)/2
	low := match

	for true {
		saIndex := suffixArray[mid]

		if low == high {
			return low
		}

		if genome[saIndex:] > read {
			high = mid
		} else {
			low = mid + 1
		}
		mid = high - (high-low)/2
	}
}

func RunSingle(genome string, read []string) (int, int) {
	suffixArray := shared.LsdRadixSort(genome)
	return BinarySuffixArraySearch(genome, read, suffixArray)
}

func lowerBound(match int, high int, suffixArray []int, genome string, read string) int {
	mid := high - (high-match)/2
	low := match

	for true {
		saIndex := suffixArray[mid]

		if low == high {
			return low
		}

		if genome[saIndex:] < read {
			high = mid
		} else {
			low = mid + 1
		}
		mid = high - (high-low)/2
	}
	return -1
}

func BinarySuffixArraySearch(genome string, read string, suffixArray []int) (int, int) {
	// rounds down
	mid := (len(suffixArray) - 1) / 2
	low := 0
	high := len(suffixArray) - 1
	//m := len(read)

	for true {
		saIndex := suffixArray[mid]
		if genome[saIndex:len(genome)-1] == read {
			upperBound := upperBound(saIndex, high, suffixArray, genome, read)
			lowerBound := lowerBound(saIndex, high, suffixArray, genome, read)
			return lowerBound, upperBound
		}

		if genome[saIndex:] > read {
			high = mid
		} else {
			low = mid + 1

		}
		mid = high - (high-low)/2

	}
	return -1, -1
}
