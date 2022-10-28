package shared

import (
	"fmt"
	"math"
)

/*Performs binary search on
First we identify a match. When we the match, we branch out and find upper and lower bound.*/
func BinarySearch(genome string, read string, suffixArray []int) (int, int) {
	// default integer divison rounds down
	low := 0
	high := len(suffixArray) - 1

	//edgecase where read is empty. Return empty interval (genome[0:0[)
	if read == "" {
		return 0, 0
	}

	for low < high-1 {
		mid := high - (high-low)/2
		saIndex := suffixArray[mid]
		//check if match (only on strings at least as long as read)
		if saIndex+len(read) < len(genome) {
			if genome[saIndex:saIndex+len(read)] == read {
				//a match has been found now search for upper and lower bound.
				upperBound := upperBound(mid, high, suffixArray, genome, read)
				lowerBound := lowerBound(mid, low, suffixArray, genome, read)
				return lowerBound, upperBound
			}
		}
		fmt.Println(genome[saIndex:])
		if genome[saIndex:] > read {
			high = mid
		} else {
			low = mid + 1
		}
	}
	//case where no match is found. Interval genome[0:0[ is empty.
	return 0, 0
}

func upperBound(match int, high int, suffixArray []int, genome string, read string) int {
	low := match

	//we break when we are left with two values. The last element that matches, and the first element that does not.
	for low < high-1 {
		//round up
		mid := low + int(math.Ceil(float64(high-low)/2.0))
		saIndex := suffixArray[mid]

		if saIndex+len(read) < len(genome) {
			if genome[saIndex:saIndex+len(read)] > read {
				//mid is too high
				high = mid
			} else {
				//we have some match (might not be greatest match)
				low = mid
			}
			//case where suffix is shorter than read
		} else {
			high = mid
		}

	}
	//We return the idx of first element that is not included, since we define intervals as '[i,j['. Result is either high or high+1.
	if suffixArray[high]+len(read) < len(genome) {
		if genome[suffixArray[high]:suffixArray[high]+len(read)] == read {
			return high + 1
		} else {
			return high
		}
	} else {
		//we look at higher element that is also shorter than our pattern
		return high
	}
}
func lowerBound(match int, low int, suffixArray []int, genome string, read string) int {
	high := match

	//we break when we are left with two values. The last element before match, and the first match.
	for low < high-1 {

		mid := low + (high-low)/2
		saIndex := suffixArray[mid]

		if genome[saIndex:] < read {
			//mid is too low
			low = mid
		} else {
			//we have some match (might not be smallest match)
			high = mid
		}

	}
	//check if value low is the match, and return if it is. Otherwise match is low+1. (high)
	if suffixArray[low]+len(read) < len(genome) {
		if genome[suffixArray[low]:suffixArray[low]+len(read)] == read {
			return low
		} else {
			return high
		}
	} else {
		return high
	}

}
