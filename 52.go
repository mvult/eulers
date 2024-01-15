package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	var i = 1
	for {
		if withinScope(i) {
			if isValid(i) {
				break
			}
		} else {
			i = nextWithinScope(i)
		}

		if i%10000 == 0 {
			fmt.Println(i)
		}
		i++
	}
	fmt.Println(i)
	fmt.Println(time.Since(t1))
}

func withinScope(i int) bool {
	if i < 100 {
		return true
	}

	s := strconv.Itoa(i)
	return string(s[0]) == "1" && string(s[1]) != "7"
}

func nextWithinScope(i int) int {
	s := strconv.Itoa(i)
	tmp := ""
	for i := 0; i < len(s); i++ {
		tmp = tmp + "0"
	}
	tmp = "1" + tmp
	ret, _ := strconv.Atoi(tmp)
	return ret
}

func isValid(i int) bool {
	return isPermutation(i, i*2) && isPermutation(i, i*3) && isPermutation(i, i*4) && isPermutation(i, i*5) && isPermutation(i, i*6)
}

func isPermutation(i1, i2 int) bool {
	s1 := strconv.Itoa(i1)
	s2 := strconv.Itoa(i2)

	ss1 := strings.Split(s1, "")
	ss2 := strings.Split(s2, "")

	sort.Strings(ss1)
	sort.Strings(ss2)

	return strings.Join(ss1, "") == strings.Join(ss2, "")
}
