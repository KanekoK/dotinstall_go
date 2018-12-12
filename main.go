package main

import "fmt"

func main() {
	// const name = "taguchi"
	// name = "fkoji"
	// fmt.Println(name)
	const (
		sun = iota
		mon
		tue
	)
	fmt.Println(sun, mon, tue)
}
