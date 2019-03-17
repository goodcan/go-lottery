package main

import "fmt"

func moveZeroes(nums []int) {
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	if len(nums1) == 0 && len(nums2) == 0 {
		return res
	}

	d := map[int]int{}

	for _, v := range nums1 {
		if _, ok := d[v]; ok {
			d[v]++
		} else {
			d[v] = 1
		}
	}

	for _, v := range nums2 {
		if _, ok := d[v]; ok && d[v] > 0 {
			res = append(res, v)
			d[v]--
		}
	}

	return res
}

func productExceptSelf(nums []int) []int {
	var res []int
	t := map[int]int{}

	for i, v := range nums {
		_, ok := t[v]
		if ok {
			res = append(res, t[v])
		} else {
			p := 1
			for _, j := range nums[:i] {
				p *= j
			}
			for _, j := range nums[i+1:] {
				p *= j
			}
			t[v] = p
			res = append(res, t[v])
		}
	}
	return res
}

func main() {
	t := []int{0, 1, 0, 3, 12}
	moveZeroes(t)
	fmt.Println(t)
	fmt.Println(intersect([]int{1, 2, 3, 4}, []int{2, 4, 6, 8}))
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
