package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var dict map[int]interface{}
var arr []int

/*
1. First generate all primes up to 10000.  Set p.
2. Divide p into groups stemming from the same digits.  E.g., 1234 = [1234,1243, 1324, 1342, 1432, 1423...]
3. Check the solution space at this point..... If
*/

var dict2 map[string][]int

func main() {
	t1 := time.Now()
	setUpPrimes()
	dict2 = make(map[string][]int)

	for _, x := range arr {
		if x > 1000 {
			placeNum(x)
		}
	}

	for _, x := range dict2 {
		if len(x) > 2 {
			if checkSequence(x) {
				break
			}
		}
	}

	fmt.Println(time.Since(t1))
}

func checkSequence(nums []int) bool {
	for i, x := range nums {
		for j, y := range nums {
			for h, z := range nums {
				if i != j && i != h && h != j && x < y && y < z && x != 1487 {
					if checkTrio(x, y, z) {
						fmt.Println(x, y, z)
						return true
					}
				}
			}
		}
	}
	return false
}

func checkTrio(x, y, z int) bool {
	return y-x == z-y
}

func placeNum(i int) {
	lp := lowestPermutation(i)
	dict2[lp] = append(dict2[lp], i)
}

func lowestPermutation(i int) string {
	ss := strings.Split(strconv.Itoa(i), "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func setUpPrimes() {
	arr = primesUpToX(10000)
	dict = make(map[int]interface{})

	for _, x := range arr {
		dict[x] = nil
	}
}

func primesUpToX(x int) []int {
	arr := make([]int, x)
	var ret []int

	for i := 2; i < x; i++ {
		if arr[i] == 0 {
			//Go and mark all multiples
			for j := 2; j*i < x; j++ {
				arr[i*j] = 1
			}
		}
	}
	for i, x := range arr {
		if x == 0 && i > 1 {
			ret = append(ret, i)
		}
	}
	return ret
}
