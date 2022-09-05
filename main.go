package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"
	// "strconv"
	// "strings"
	"github.com/SSMalik99/algo/chunks"
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
	// fmt.Println(time.Now())
	// fmt.Println("------------", letFindEffectiveDifitForSum(18), "------------")
	// fmt.Println(time.Now())
	// fmt.Println("------------", numberOfSum(18), "------------")
	// fmt.Println(time.Now())
	// fmt.Println("------------", findNumberWithDigitSum(18), "------------")
	// fmt.Println(time.Now())
	// fmt.Println(primeFactorNumber(300))

	// fmt.Println(fibonacciSeries(10))

	// fmt.Println(greatestCommonDivisorUsingPrimeFactor(675, 465))
	// fmt.Println(greatestCommonDivisorUsingPrimeFactor(2, 3))

	// fmt.Println(greatestCommonDivisorUsingEuclidean(675, 465))

	fmt.Println("------Program exit!------")
}


// GCD with Euclidean formula
/*
	The Euclidean Algorithm for finding GCD(A,B) is as follows:
		If A = 0 then GCD(A,B)=B, since the GCD(0,B)=B, and we can stop.  
		If B = 0 then GCD(A,B)=A, since the GCD(A,0)=A, and we can stop.  
		Write A in quotient remainder form (A = B⋅Q + R)
		
		Find GCD(B,R) using the Euclidean Algorithm since GCD(A,B) = GCD(B,R)
*/
func greatestCommonDivisorUsingEuclidean(numbers... int) int {
	if len(numbers) > 2 {
		panic(fmt.Sprintf(chunks.Yellow, "\n\n\tThis function is only to find GCD between two numbers", chunks.Reset))	
	}

	if numbers[0] == 0 {
		return numbers[1]
	}else if numbers[1] == 0 {
		return numbers[0]
	}

	// Write A in quotient remainder form (A = B⋅Q + R)
	if numbers[0] > numbers[1] {
		
		temp := numbers[0]
		numbers[0] = numbers[1]
		numbers[1] = temp % numbers[1]
	}else {
		temp := numbers[1]
		numbers[1] = numbers[0]
		numbers[0] = temp % numbers[0]
	}
	return greatestCommonDivisorUsingEuclidean(numbers...)
}


// Greatest common divisor GCD with prime factors
func greatestCommonDivisorUsingPrimeFactor(numbers... int) int {
	if len(numbers) > 2 {
		panic(fmt.Sprintf(chunks.Yellow, "\n\n\tThis function is only to find GCD between two numbers", chunks.Reset))
	}
	number, divisiblity := checkNumeberDivisible(numbers[0], numbers[1])
	if divisiblity {
		return number
	}
	primeFactors1 := primeFactorNumber(numbers[0])
	primeFactors2 := primeFactorNumber(numbers[1])
	overAllFactor := []int{}
	compareAndGenerateGCD := func (factors1, factors2 []int) int {
		for _, factor := range factors1 {
			if chunks.NumberInList(factors2, factor) {
				factors2 = removeNumberFromSlice(factors2, factor, false)
				overAllFactor = append(overAllFactor, factor)
			}
		}
		total := 1
		for _, fac := range overAllFactor {
			total *= fac	
		}
		return total
	}

	if len(primeFactors1) > len(primeFactors2) {
		return compareAndGenerateGCD(primeFactors1, primeFactors2)	
	}

	return compareAndGenerateGCD(primeFactors2, primeFactors1)
	
}

func checkNumeberDivisible(num1, num2 int) (int, bool) {
	// check for the zero numbers
	if num1 == 0 {
		return num2, true
	}
	if num2 == 0 {
		return num1, true
	}

	if num1 > num2 {
		if num1 % num2 == 0 {
			return num2, true
		}
	}else {
		if num2 % num1 == 0 {
			return num1, true
		}
	}

	return 0, false
}


func removeNumberFromSlice(slice []int, number int, deleteAll bool) []int {
	var newSlice []int

	for index, num := range slice {

		if deleteAll {
			if num != number {
				newSlice = append(newSlice, num)
			}
		}else{

			if num == number {
				newSlice = append(newSlice, slice[index+1:]...)
				newSlice = append(newSlice, slice[:index]...)
				break;
			}
		}
	}
	return newSlice
}


func fibonacciSeries(numberOfTerms int) *[]int {
	fibonaci := []int{0, 1}
	for i := len(fibonaci); i < numberOfTerms; i++ {
		fibonaci = append(fibonaci, fibonaci[i-1] + fibonaci[i-2])
	}
	return &fibonaci // instead of sharing whole the slice share the address
}

func primeFactorNumber(num int) []int {
	var factors = []int{} // factors' stack
	nextPrime := nextPrimeNumber(1) //prime number next to 1
	// for loop to add prime factor to the factors stack
	for {
		if num < 2 {
			break // break the loop if number is no more activated
		}
		// for loop to find next prime number
		for {
			if num % nextPrime == 0 {				
				break // break the loop if next Prime is available
			}
			nextPrime = nextPrimeNumber(nextPrime)
		}
		// add prime numbr to the stack
		if num % nextPrime == 0 {
			factors = append(factors, nextPrime)
		}
		num = num / nextPrime // make number equal to the quotient
	}
	return factors
}
// find next prime number after a given number
func nextPrimeNumber(number int) int {
	if number == 0 || number == 1 {
		return 2 // return very first prime
	}
	number += 1 // increase the number to find next number
	for {
		if isPrime(number) {
			break // break loop if number is prime 
		}
		number += 1 // increase number again
	}
	return number // return the number
}

// check number is prime or not
func isPrime(number int) bool {
	if number == 2 {
		return true // return first prime
	}
	// loop to divide the number
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}

	return true
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

