package main

import "fmt"

func World(i int) string {
	if i == 5 {
		return "World!"
	} else {
		return "New World!"
	}
}
func main() {
	fmt.Print("Hello ")
	suffix := World(5)
	fmt.Println(suffix)
}
