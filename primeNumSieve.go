package main

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
