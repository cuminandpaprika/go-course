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
	for index := 0; index < len(unsorted) -1; index++ {
		fmt.Fprintf(out, "Comparing %v and %v\n", unsorted[index],unsorted[index+1])
		if (unsorted[index] > unsorted[index+1]) {
			unsorted[index+1], unsorted[index] = unsorted[index], unsorted[index+1]
			fmt.Fprintf(out, "Swapping %v and %v\n", unsorted[index],unsorted[index+1])
		}
	}
	return unsorted
}