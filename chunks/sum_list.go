package chunks

func SumOfList(list []int) int  {
	sum := 0
	for _, number := range list {
		sum = sum + number
	}	

	return sum
}

