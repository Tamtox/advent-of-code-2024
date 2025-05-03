package day2

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Task1() {
	path := filepath.Base("/Day_2")
	fmt.Println("Path: ", path)
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fileString := string(file)
	lines := strings.Split(fileString, "\n")
	safeAmount := 0
	for _, line := range lines {
		numsStrArr := strings.Split(line, " ")
		numsArr := make([]int, 0)
		for _, numStr := range numsStrArr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
				return
			}
			numsArr = append(numsArr, num)
		}
		isSafe := true
		pattern := "increasing"
		if numsArr[0] > numsArr[1] {
			pattern = "decreasing"
		}
		for i := 1; i < len(numsArr); i++ {
			num := numsArr[i]
			prevNum := numsArr[i-1]
			diff := num - prevNum
			if pattern == "increasing" && (diff < 1 || diff > 3) {
				isSafe = false
				break
			} else if pattern == "decreasing" && (diff > -1 || diff < -3) {
				isSafe = false
				break
			}
		}
		fmt.Println("Line: ", line, " is safe: ", isSafe)
		if isSafe {
			safeAmount++
		}
	}
	fmt.Println("Safe amount: ", safeAmount)
}
