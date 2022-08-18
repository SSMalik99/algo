package chunks

import (
	"testing"
)

func TestNumberInList(t *testing.T) {
	result := NumberInList([]int{1, 2, 34, 4, 5, 5, 6, 6, 78, 4, 52, 52, 523, 452, 452, 4, 2345, 23, 5234, 52, 345}, 768)

	if result {
		t.Errorf("768 this number should not be contained")
	}
}

func BenchmarkNumberInList(b *testing.B) {
	list := []int{1, 2, 34, 4, 5, 5, 6, 6, 78, 4, 52, 52, 523, 452, 452, 4, 2345, 23, 5234, 52, 345}
	for i := 0; i < b.N; i++ {
		
		NumberInList(list, i)
	}
}




func TestSumOfList(t *testing.T) {
	list := []int{}
	oldSum := 0
	for i := 1; i <= 100; i++ {
		
		list = append(list, i)
		newSum := SumOfList(list)
		
		if newSum != ( oldSum + i ) {
			t.Errorf("sum should be = %v but geting %v", oldSum+i, newSum)
		}
		oldSum = newSum
	}
}

func BenchmarkSumOfList(b *testing.B) {
	list := []int{}

	for i := 0; i < b.N; i++ {

		list = append(list, i)
		SumOfList(list)
	}
}