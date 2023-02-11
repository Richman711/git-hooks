package main

import "fmt"

func World(i int) string {
	if i == 5 {
		return "1World!"
	} else {
		return "New World!"
	}
}
func main() {
	fmt.Print("Hello ")
	suffix := World(2)
	fmt.Println(suffix)
}
