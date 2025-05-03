package day1

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Task 1
func Task1() {
	filePath := filepath.Join("day1", "input.txt")
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)
	for _, line := range lines {
		numsArr := strings.Split(line, "   ")
		num1, err1 := strconv.Atoi(numsArr[0])
		if err1 != nil {
			fmt.Println("Error converting string to int: ", err1)
			return
		}
		nums1 = append(nums1, num1)
		num2, err2 := strconv.Atoi(numsArr[1])
		if err2 != nil {
			fmt.Println("Error converting string to int: ", err2)
			return
		}
		nums2 = append(nums2, num2)
	}
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	len1 := len(nums1)
	len2 := len(nums2)
	len := len1
	if len2 < len1 {
		len = len2
	}
	sum := 0
	for i := 0; i < len; i++ {
		num1 := nums1[i]
		num2 := nums2[i]
		diff := 0
		if num1 < num2 {
			diff = num2 - num1
		} else if num2 < num1 {
			diff = num1 - num2
		}
		sum += diff
	}
	fmt.Println("Diffs: ", sum)
}

// Task 2
func Task2() {
	filePath := filepath.Join("day1", "input.txt")
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	nums1 := make([]int, 0)
	nums2Map := make(map[int]int)
	for _, line := range lines {
		numsArr := strings.Split(line, "   ")
		num1, err1 := strconv.Atoi(numsArr[0])
		if err1 != nil {
			fmt.Println("Error converting string to int: ", err1)
			return
		}
		nums1 = append(nums1, num1)
		num2, err2 := strconv.Atoi(numsArr[1])
		if err2 != nil {
			fmt.Println("Error converting string to int: ", err2)
			return
		}
		num2Entry := nums2Map[num2]
		if num2Entry == 0 {
			nums2Map[num2] = 1
		} else {
			nums2Map[num2] = num2Entry + 1
		}
	}
	sum := 0
	for _, num := range nums1 {
		num2Entry := nums2Map[num]
		sum += num * num2Entry
	}
	fmt.Println("Sum: ", sum)
}
