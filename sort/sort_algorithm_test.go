package sort

import (
	"testing"
)

// 选择排序测试
func TestSelectionSort(t *testing.T) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	corr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 10}
	sortArr := SelectionSort(arr)
	for i, v := range sortArr {
		if corr[i] != v {
			t.Errorf("SelectionSort failed, error index : %d", i)
		}
	}
}

// 选择排序微基准
func BenchmarkSelectionSort(b *testing.B) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}

// 快速排序测试
func TestQuickSort(t *testing.T) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	corr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 10}
	sortArr := QuickSort(arr)
	for i, v := range sortArr {
		if corr[i] != v {
			t.Errorf("QuickSort failed, error index : %d", i)
		}
	}
}

// 快速排序微基准
func BenchmarkQuickSort(b *testing.B) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}

// 冒泡排序测试
func TestBubbleSort(t *testing.T) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	corr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 10}
	sortArr := BubbleSort(arr)
	for i, v := range sortArr {
		if corr[i] != v {
			t.Errorf("BubbleSort failed, error index : %d", i)
		}
	}
}

// 冒泡排序微基准
func BenchmarkBubbleSort(b *testing.B) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}

// 直接插入排序测试
func TestDirectInsertSort(t *testing.T) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	corr := []int{1, 1, 2, 2, 3, 3, 4, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 10}
	sortArr := DirectInsertSort(arr)
	for i, v := range sortArr {
		if corr[i] != v {
			t.Errorf("DirectInsertSort failed, error index : %d", i)
		}
	}
}

// 直接插入排序微基准
func BenchmarkDirectInsertSort(b *testing.B) {
	arr := []int{3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3,
		3, 7, 2, 9, 1, 4, 6, 8, 10, 5, 6, 4, 8, 1, 4, 2, 7, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DirectInsertSort(arr)
	}
}
