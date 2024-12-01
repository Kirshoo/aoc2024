package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
	"slices"
	"math"
)

func part1(filePath string) (int, error) {
	// Read from an input file
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	leftList, rightList := []int{}, []int{}
	
	rd := bufio.NewReader(file)
	for {	// Read file line by line
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}
		
		// Trim the white spaces
		line = strings.TrimSpace(line)

		// Get fields we are interested in
		fields := strings.Fields(line)

		// Put the values into the respective lists
		left, _ := strconv.Atoi(fields[0])
		right, _ := strconv.Atoi(fields[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	// Sort the list to compare ith smallest values
	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0
	for i := 0; i < len(leftList); i++ {
		sum += int(math.Abs(float64(rightList[i] - leftList[i])))
	}

	return sum, nil
}

func part2(filePath string) (int, error) {
	// Read from an input file
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	leftList, rightSet := []int{}, make(map[int]int)
	
	rd := bufio.NewReader(file)
	for {	// Read file line by line
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}
		
		// Trim the white spaces
		line = strings.TrimSpace(line)

		// Get fields we are interested in
		fields := strings.Fields(line)

		// Put the values into the respective lists
		left, _ := strconv.Atoi(fields[0])
		right, _ := strconv.Atoi(fields[1])
		leftList = append(leftList, left)
		rightSet[right] += 1
	}
	
	similarity := 0
	for i := 0; i < len(leftList); i++ {
		similarity += leftList[i] * rightSet[leftList[i]]
	}

	return similarity, nil
}

func main() {
	filename := "./input/task.txt"

	answer1, err := part1(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer1)


	answer2, err := part2(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer2)
}
