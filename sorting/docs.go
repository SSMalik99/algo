package sorting

import "fmt"

func ShowDocumentation() {
	docs := `
		worst case to finish a code is BIG O
		Docs :: Big O -> is the estimation of worst case for a given function or the code

		O(1) ---> when there is a constant time of execution independent of the input

		for example : 

		func Add(a, b int) {
			// whatever value of the "a" and "b" function is gonna take time same as for the smaller number
			return a + b

			// so its constant time function 
		}

		// now O(N), N is the value of max in the below function

		func sumOfMax(max int) {
			for i:=0; i< max; i ++ {
				return (max * ( max + 1 )) / 2
			}

			// so time of the function will depends on the max
			// so Big-o will be O(max) becuase function taking time to loop over things for the as much time
		}

	`

	fmt.Println(docs)
}