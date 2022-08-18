package chunks

func NumberInList(list []int, num int) bool{
	for _,i := range list {
		if i == num {
			return true
		}
	}
	return false
}