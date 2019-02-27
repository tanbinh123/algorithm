package sort

// 选择排序
func SelectionSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 0; j < length-i; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[maxIndex], arr[length-i-1] = arr[length-i-1], arr[maxIndex]
	}
	return arr
}

// 快速排序
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return quickSortPartition(arr, 0, len(arr)-1)
}

func quickSortPartition(arr []int, low int, high int) []int {
	if low >= high {
		return arr
	}
	i, j := low, high
	pivotKey := arr[i]
	for i < j {
		for j > i && arr[j] >= pivotKey {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

		for i < j && arr[i] <= pivotKey {
			i ++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	quickSortPartition(arr, low, i-1)
	quickSortPartition(arr, i+1, high)
	return arr
}

// 冒泡排序
func BubbleSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 0; i < length-1; i ++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 直接插入排序
func DirectInsertSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	var insertValue int
	for i := 1; i < length; i++ {
		insertValue = arr[i]
		var j int
		for j = i - 1; j >= 0; j-- {
			if arr[j] <= insertValue {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = insertValue
	}
	return arr
}
