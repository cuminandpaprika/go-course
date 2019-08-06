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
	var isFinishedSorting bool
	for {
		isFinishedSorting = true
		for index := 0; index < len(unsorted) -1; index++ {
			fmt.Fprintf(out, "Comparing %v and %v\n", unsorted[index],unsorted[index+1])			
			if (leftIsGreater(unsorted[index:index+2])) {
				swap(unsorted[index:index+2])
				isFinishedSorting = false
			}
		}
		if (isFinishedSorting) {
			return unsorted
		}
	}
}

func swap(unsortedPair []int) {
	fmt.Fprintf(out, "Swapping %v and %v\n", unsortedPair[0],unsortedPair[1])
	unsortedPair[0], unsortedPair[1] = unsortedPair[1], unsortedPair[0]	
} 

func leftIsGreater(pair []int) bool {
	return pair[0] > pair[1]
}