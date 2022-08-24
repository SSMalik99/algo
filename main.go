package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"
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



// byte is 8 bit, bit is 0 or 1

func convertDecimalToOtherBase(num, base int) string {
	const charSet = "0123456789ABCDEF"
	// var result string
    var sb strings.Builder	

	for num > 0 {
		rem := num % base
		num = num / base
		sb.WriteByte(charSet[rem])
	}
	
	return reverseByDirectLoop(sb.String());
}

func getIndexFromString(str string, char string) int {
	
	for index, val := range str {
		if char == string(val) {
			return index
		}
	}

	return -1
}

// base should be 2 - 16
func convertAnyBaseToDecimal(otherBaseValue string, base int) int {
	const charSet = "0123456789ABCDEF"

	if base > 16 {
		panic("Base should be in between 2 - 16")
	}

	otherBaseValue = reverseByDirectLoop(otherBaseValue)

	var decimal int

	for powerIndex, v := range otherBaseValue {

		mathValue := getIndexFromString(charSet, string(v))//strings.Index(charSet, string(v))
		if mathValue < 0 {
			panic("you are entring wrong data for the base")
		}
		decimal = decimal + (mathValue * int(math.Pow(float64(base), float64(powerIndex))))

	}
	return decimal;
}

func convertAnyBaseToAny(baseValue string, originalBase, finalBase int) string {

	return convertDecimalToOtherBase(convertAnyBaseToDecimal(baseValue, originalBase), finalBase)
	
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
	// fmt.Println(findEntityIsPalilndrom("joj"))
	// fmt.Println(findEntityIsPalilndrom("jojytrrrtadf"))
	// fmt.Println(findEntityIsPalilndrom(76567))
	// fmt.Println(findEntityIsPalilndrom(7656765))


	// convert decimal to other base
	// fmt.Println(convertDecimalToOtherBase(12, 2))
	// fmt.Println(convertDecimalToOtherBase(12, 8))
	// fmt.Println(convertDecimalToOtherBase(12, 16))
	/**
		1100
		14
		C
		1000101011101110000011000110
		1053560306
		8AEE0C6
	*/
	// fmt.Println(convertDecimalToOtherBase(145678534, 2))
	// fmt.Println(convertDecimalToOtherBase(145678534, 8))
	// fmt.Println(convertDecimalToOtherBase(145678534, 16))
	

	// convert any base to decimal
	// fmt.Println(convertAnyBaseToDecimal("1100", 2))
	// fmt.Println(convertAnyBaseToDecimal("14", 8))
	// fmt.Println(convertAnyBaseToDecimal("C", 16))
	
	// fmt.Println(convertAnyBaseToDecimal("1000101011101110000011000110", 2))
	// fmt.Println(convertAnyBaseToDecimal("1053560306", 8))
	// fmt.Println(convertAnyBaseToDecimal("8AEE0C6", 16))


	// Convert base to base

	// fmt.Println(convertAnyBaseToAny("1100", 2, 8))
	// fmt.Println(convertAnyBaseToAny("14", 8, 16))
	// fmt.Println(convertAnyBaseToAny("C", 16, 2))
	// fmt.Println(convertAnyBaseToAny("1000101011101110000011000110", 2, 16))
	// fmt.Println(convertAnyBaseToAny("1053560306", 8, 10))
	// fmt.Println(convertAnyBaseToAny("8AEE0C6", 16, 2))

	
	// fmt.Println(FindNumbersInList([]int{2, 5, 3, 4}, 7))


	fmt.Println(time.Now())
	// fmt.Println("------------", findNumberWithDigitSum(100), "------------")
	// fmt.Println(time.Now())

	// fmt.Println("------------", numberOfSum(100), "------------")

	// fmt.Println("------------", letFindEffectiveDifitForSum(100), "------------")
	fmt.Println(time.Now())
	fmt.Println("------------", letFindEffectiveDifitForSum(18), "------------")
	fmt.Println(time.Now())
	fmt.Println("------------", numberOfSum(18), "------------")
	fmt.Println(time.Now())
	fmt.Println("------------", findNumberWithDigitSum(18), "------------")
	fmt.Println(time.Now())

	fmt.Println("------Program exit!------")
}



func sumOfDigits(number int) int {
	
	quotient := number
	var newNum int
	for quotient > 0 {
		newNum = newNum + (quotient % 10)
		quotient /= 10
	}
	return newNum
}

/*
// to find the number whose digits sum is 100
OutPut:
	------Program started!------
	start time : 2022-08-24 17:39:24.8072856 +0530 IST m=+0.007834801
	100000000012 --> loop count
	Answer : 199999999999
	end time : 2022-08-24 18:33:21.8627263 +0530 IST m=+3237.063275501
*/
func letFindEffectiveDifitForSum(digitSum int) int {
	lowerBound := math.Pow(10, 1) - 1
	upperBound := math.Pow(10, 2) - 1
	power := 3
	var count int
	for {
		count += 1
		uppperBoundSum := sumOfDigits(int(upperBound))
		
		if digitSum == uppperBoundSum {
			return int(upperBound)
		}
		fmt.Println(digitSum, uppperBoundSum)
		if digitSum > uppperBoundSum {
			lowerBound = upperBound
			upperBound = math.Pow(10,float64(power)) - 1
			power += 1
			continue
		}
		break
	}
	
	for lowerBound < upperBound {
		count += 1
		if sumOfDigits(int(lowerBound)) == digitSum {
			fmt.Println(count)
			return int(lowerBound)
		}
		lowerBound += 1
	}

	return 0;
}


// one way is to calculate with lower and upper bound
// To find number whose digit sum is 100 tooks below parimeter
// start time : 2022-08-24 13:56:38.7467926 +0530 IST m=+0.006289301
// end time : 2022-08-24 15:07:24.2370053 +0530 IST m=+4245.496502001
// loop count to calculate bounds = 11
// loop count : 100000000012 // loop count is not including for loop to get sum of digit
func numberOfSum(digitSum int) int {
	lowerBound := math.Pow(10, 1) - 1
	upperBound := math.Pow(10, 2) - 1
	power := 3
	var count int
	for {

		count += 1
		uppperBoundSum := sumOfDigits(int(upperBound))
		if digitSum == uppperBoundSum {
			return int(upperBound)
		}else if digitSum > uppperBoundSum {
			lowerBound = upperBound
			upperBound = math.Pow(10,float64(power)) - 1
			power += 1
		}else {
			break
		}
		
	}
	var number int;
	for i := lowerBound; i < upperBound; i++ {
		count += 1
		if sumOfDigits(int(i)) == digitSum {
			number = int(i)
			break
		}
	}
	fmt.Println(count)
	return number
}

// one direct way
// To find number whose digit sum is 100 tooks below parimeter
// Start Time : 2022-08-24 11:11:52.4158211 +0530 IST m=+0.003609801
// End Time : 2022-08-24 12:14:13.1429595 +0530 IST m=+3740.730773501
// loop count : 199999999900 // loop count is not including for loop to get sum of digit
// answer : 199999999999
func findNumberWithDigitSum(digitSum int) int {
	number := digitSum
	var count int
	for {
		count += 1
		quotient := number
		var newNum int
		for quotient > 0 {
			newNum = newNum + (quotient % 10)
			quotient /= 10
		}
		if newNum == digitSum {
			break
		}
		number += 1
	}
	fmt.Println(count)
	return number
}

func FindNumbersInList(list []int, sum int) (int, int) {
	
	previousIndexes := [2]int{-1, -1}

	for outerIndex, outerValue := range list {
		for innerIndex, innerValue := range list {
			if outerIndex == innerIndex {
				continue
			}
			if outerValue + innerValue == sum {	
				if (previousIndexes[0] != -1) && (previousIndexes[1] != -1) {
					
					// here we can apply code to filter the indexes we want
						// let's implement simple one where we will require only those indexes which
						// have greater difference value mean if (-1, -3 given then -1 shuld be answer)
					// of their original values
					if (outerValue - innerValue) < (list[previousIndexes[0]] - list[previousIndexes[1]]) {
						continue
					}
				}

				previousIndexes[0] = outerIndex 
				previousIndexes[1] = innerIndex
			}
		}
	}

	return previousIndexes[0], previousIndexes[1]
}

