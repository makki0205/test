package main

import "fmt"

func main() {
	Lesson2(3)
}
func Lesson2(n int) {
	fmt.Println("start")
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("■")
		}
		fmt.Println("")
	}
}
