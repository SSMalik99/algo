package sorting

import (
	"fmt"
)



// bubble sort is good at small list or almost sorted list 
// big o -> O(n^2) n is the size of list
func BubbleSort(list []int ) {
	fmt.Println(list)

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j + 1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	fmt.Println(list)
}