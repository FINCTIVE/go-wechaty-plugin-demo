package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0.05
	fmt.Println(time.Duration(a) * time.Hour)
}
