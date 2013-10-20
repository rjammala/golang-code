package main

import "fmt"

func findLength(n uint64) int {
	length := 1

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
			length++
		} else {
			n = 3*n + 1
			length++
		}
	}

	return length
}

func main() {

	max := 0
	for i := 2; i < 1000000; i++ {
		tempResult := findLength(uint64(i))
		if tempResult > max {
			max = tempResult
		}
	}

	fmt.Println(max)
}
