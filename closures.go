package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	i := 0
	for {
		fmt.Println(nextInt())
		i++
		if i == 3 {
			break
		}
	}

	newInts := intSeq()
	fmt.Println(newInts())
}