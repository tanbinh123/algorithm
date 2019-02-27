package easy

import (
	"fmt"
	"testing"
)

/*
1. 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */

// 暴力法
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 两遍哈希表
// 通过hash表降低查找速度, 通过hash索引，使查找O(n)降低为O(1)
func TwoSumTwiceHash(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, v := range nums {
		hash[v] = i
	}
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		j, ok := hash[remain]
		if ok && i != j {
			return []int{i, j}
		}
	}
	return []int{}
}

// 一遍哈希表
// 边存边找
func TwoSumOnceHash(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, v := range nums {
		remain := target - nums[i]
		j, ok := hash[remain]
		if ok && i != j {
			if i > j {
				return []int{j, i}
			}
			return []int{i, j}
		} else {
			hash[v] = i
		}
	}
	return []int{}
}


// 两数之和暴力法测试
func TestTwoSum(t *testing.T) {
	re := TwoSum([]int{3, 2, 4}, 6)
	corr := []int{1, 2}
	for i, v := range corr {
		if re[i] != v {
			t.Errorf("TwoSum failed, error index : %d", i)
		}
	}
}

func TestTwoSumTwiceHash(t *testing.T) {
	re := TwoSumTwiceHash([]int{3, 2, 4}, 6)
	corr := []int{1, 2}
	for i, v := range corr {
		if re[i] != v {
			t.Errorf("TwoSum failed, error index : %d", i)
		}
	}
}

func TestTwoSumOnceHash(t *testing.T) {
	re := TwoSumOnceHash([]int{3, 2, 4}, 6)
	corr := []int{1, 2}
	for i, v := range corr {
		if re[i] != v {
			fmt.Print(re[i])
			t.Errorf("TwoSum failed, error index : %d", i)
		}
	}
}
