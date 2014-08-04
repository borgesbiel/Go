package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argsIn := os.Args[1:]
	numbers := make([]int, len(argsIn))

	for i, n := range argsIn {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("[%s] not a valid number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	indicePivot := len(n) / 2
	pivot := n[indicePivot]
	n = append(n[:indicePivot], n[indicePivot+1:]...)

	minor, major := split(n, pivot)

	return append(
		append(quicksort(minor), pivot),
		quicksort(major)...)
}

func split(numbers []int, pivot int) (minor []int, major []int) {
	for _, n := range numbers {
		if n <= pivot {
			minor = append(minor, n)
		} else {
			major = append(major, n)
		}
	}

	return minor, major
}