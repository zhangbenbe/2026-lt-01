package main

import (
	"math"
	"sort"
)

// 1
func twoSum(nums []int, target int) []int {
	var cnt = make(map[int]int)
	var result []int
	for index, num := range nums {
		if _, ok := cnt[target-num]; ok {
			result = []int{index, cnt[target-num]}
		} else {
			cnt[num] = index
		}
	}
	return result
}

// 11
func maxArea(height []int) int {
	var result = 0
	var left, right = 0, len(height) - 1
	for left < right {
		result = max(result, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

// 14
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}

// 15
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, length := make([][]int, 0), 0, 0, 0, len(nums)

	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && index < end {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			addNum := nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}

	return result
}

// 16
func threeSumClosest(nums []int, target int) int {
	length, res, diff := len(nums), 0, math.MaxInt32
	if length > 2 {
		sort.Ints(nums)
		for i := 0; i < length-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, length-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res = sum
					diff = abs(sum - target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}
