package algorithms

func BubbleSort(arr []int) []int {
	swaped := false
	for !swaped {
		swaped = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
				swaped = false
			}
		}
	}
	return arr
}
