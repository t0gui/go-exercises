package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Do 1 ao 3 -->")
	fmt.Println(slice[0:3])
	fmt.Print("Do 5 ao 10 -->")
	fmt.Println(slice[4:])
	fmt.Print("Do 2 ao 7 -->")
	fmt.Println(slice[1:7])
	fmt.Print("Do 3 ao 9 -->")
	fmt.Println(slice[2:9])
	fmt.Print("Do 1 ao 9 usando len()-->")
	fmt.Println(slice[0:(len(slice) - 1)])
}
