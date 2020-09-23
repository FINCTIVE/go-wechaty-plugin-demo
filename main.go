package main

import "fmt"

func main() {
	var m map[int]int = make(map[int]int)
	m[1]++
	fmt.Println(m[1])
}
