package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}

func bubble(unsorted []int) []int {
	var swapped bool
	for {
		for index := 0; index < len(unsorted) -1; index++ {
			fmt.Fprintf(out, "Comparing %v and %v\n", unsorted[index],unsorted[index+1])
			fmt.Fprintln(out, unsorted[index:index+2])
			swapped = sort(unsorted[index:index+2])
		}
		if (!swapped) {
			return unsorted
		}
	}

	return unsorted
}

func sort(unsortedPair []int) bool {
	if (leftIsGreater(unsortedPair)) {
		unsortedPair[0], unsortedPair[1] = unsortedPair[1], unsortedPair[0]
		fmt.Fprintf(out, "Swapped %v and %v\n", unsortedPair[0],unsortedPair[1])
		return true
	}
	return false
} 

func leftIsGreater(pair []int) bool {
	return (pair[0] > pair[1])
}