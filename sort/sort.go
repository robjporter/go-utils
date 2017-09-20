package sort

func BubbleSort(list []byte) []byte {
	for n := len(list); n != 0; {
		newn := 0
		for i := 1; i < n; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				newn = i
			}
		}
		n = newn
	}
	return list
}
