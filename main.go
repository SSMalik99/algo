package main

import (
	"fmt"
	"github.com/SSMalik99/algo/chunks"

)


func main() {
	fmt.Println("------Program started!------")
	list := []int{1,2,34,4,5,5,6,6,78,4,52,52,523,452,452,4,2345,23,5234,52,345}
	
	fmt.Println("Is Number in list ?",chunks.NumberInList(list, 768))
	fmt.Println("Sum of list = ", chunks.SumOfList(list))
	
	fmt.Println("------Program exit!------")
}

