package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: genome-file reads-file\n")
		os.Exit(1)
	}
	genome := os.Args[1]
	reads := os.Args[2]

	//suffixArray := shared.LsdRadixSort(genome)

	fmt.Println("Search in", genome, "for the reads in", reads)
}

func upperBound(match int, high int, suffixArray []int, genome string, read string) int {
	mid := high - (high-match)/2
	low := match

	for high != low {
		saIndex := suffixArray[mid]

		if low == high {
			return low + 1
		}

		if genome[saIndex:] > read {
			if mid == high {
				return low + 1
			}
			high = mid
		} else {
			low = mid + 1
		}
		mid = high - (high-low)/2
	}
	return high + 1
}
func lowerBound(match int, low int, suffixArray []int, genome string, read string) int {
	mid := match - (match-low)/2
	high := match

	for high != low {
		saIndex := suffixArray[mid]

		if low == high {
			return low
		}
		if genome[saIndex:] < read {
			if high == mid {
				return low
			}
			high = mid
		} else {
			low = mid + 1
		}
		mid = high - (high-low)/2

	}
	return low - 1
}

func BinarySuffixArraySearch(genome string, read string, suffixArray []int) (int, int) {
	// rounds down
	mid := (len(suffixArray) - 1) / 2
	low := 0
	high := len(suffixArray) - 1
	//m := len(read)

	for low != high {
		fmt.Println(low, mid, high)
		saIndex := suffixArray[mid]

		//check if match only on strings at least as long as read
		if saIndex+len(read) < len(genome) {
			if genome[saIndex:saIndex+len(read)] == read {
				fmt.Println(genome[saIndex:len(genome)-1], "yiha")
				upperBound := upperBound(mid, high, suffixArray, genome, read)
				lowerBound := lowerBound(mid, low, suffixArray, genome, read)
				return lowerBound, upperBound
			}
		}

		if genome[saIndex:] > read {
			fmt.Println(genome[saIndex:], read)
			high = mid
		} else {
			low = mid + 1

		}
		mid = high - (high-low)/2

	}
	return -1, -1
}
