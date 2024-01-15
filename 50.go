package main

import (
	"fmt"
)

var dict map[int]interface{}
var arr []int
var results map[int]int

func main() {
	arr = primesUpToX(1000000)
	dict = make(map[int]interface{})
	results = make(map[int]int)

	fmt.Println("Generated Primes")

	for _, x := range arr {
		dict[x] = nil
	}
	fmt.Println("Created map of primes")

	for i := 22; i < 600; i++ {
		checkSequencesOfLengthX(i)
		fmt.Println("Checked length ", i)
	}

	max := 0
	answer := 0
	for i, x := range results {
		if x > max {
			answer = i
			max = x
		}
	}

	fmt.Println(answer)
}

func checkSequencesOfLengthX(length int) {
	//Lets say size is 25
	for i := 0; i < len(arr)-length; i++ {
		checkSequence(i, length)
	}
}

func checkSequence(start, length int) {
	//Lets say length is 25, start is 567
	if start+length > len(arr) {
		return
	}

	sum := 0
	for i := 0; i < length; i++ {
		sum = sum + arr[start+i]
	}
	if _, ok := dict[sum]; ok {
		results[sum] = length
	}
}
