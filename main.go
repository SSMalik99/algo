package main

import (
	"bytes"
	"fmt"
	"math"
	// "strconv"
	// "strings"
	// "github.com/SSMalik99/algo/chunks"
)

func reversStringByReverseLoop(objStr string) string {
	length := len(objStr)
	// loopLength := length/2
	var newString bytes.Buffer

	for i := length - 1; i >=0 ; i-- {
		newString.WriteString(string(objStr[i]))
	}
	return newString.String()
}

func reverseByDirectLoop(objStr string) string {
	var anotherString string
	
	for _, v := range objStr {
		anotherString = string(v) + anotherString
	}

	return anotherString
}

func DigitInNumber(num int) int {

	// one way
	return int(math.Floor(math.Log10(float64(num)) + 1))
	
	// 2nd way to find digit in number
	// newNum := num
	// count := 0
	// for newNum > 0 {
	// 	fmt.Println(newNum)
	// 	newNum = newNum / 10
	// 	count += 1
	// }
	// return count

}

func reverseNumber(num int) int {
	digitCount := DigitInNumber(num)

	var reverse int
	// start from one becaus in 1234 digts are 4 but if we want to reverse then 
	// 4 should be multiply by 1000(10 ^ 3) not by 10000 (10 ^ 4)
	for i := 0; i < digitCount; i++ {
		reminder := num % 10
		reverse = reverse + reminder * int(math.Pow(10, float64(digitCount-i-1)))
		num = num / 10
	}

	return reverse
}

// find a number or string is palindrom or not
func findEntityIsPalilndrom(argument interface{}) bool {

	switch argument.(type) {
		case int:
			if argument == reverseNumber(argument.(int)) {
				return true
			}
		case string:
			if argument == reverseByDirectLoop(argument.(string)) {
				return true
			}
	}

	return false

}


func main() {
	fmt.Println("------Program started!------")

	// list := []int{1,2,34,4,5,5,6,6,78,4,52,52,523,452,452,4,2345,23,5234,52,345}
	// fmt.Println("Is Number in list ?",chunks.NumberInList(list, 768))
	// fmt.Println("Sum of list = ", chunks.SumOfList(list))
	// fmt.Println(reversStringByReverseLoop("grtv"))
	// fmt.Println(reverseByDirectLoop("grtv"))

	// reverse a number
	// fmt.Println(reverseNumber(121))

	// find entity is palindrome
	fmt.Println(findEntityIsPalilndrom("joj"))
	fmt.Println(findEntityIsPalilndrom("jojytrrrtadf"))
	fmt.Println(findEntityIsPalilndrom(76567))
	fmt.Println(findEntityIsPalilndrom(7656765))

	fmt.Println("------Program exit!------")
}

