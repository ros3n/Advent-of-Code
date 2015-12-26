package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func area(dimensions []int) (int, int) {
	areas := make([]int, 3)
	areas[0] = dimensions[0] * dimensions[1]
	areas[1] = dimensions[1] * dimensions[2]
	areas[2] = dimensions[0] * dimensions[2]
	sort.Ints(areas)
	return 2*areas[0] + 2*areas[1] + 2*areas[2], areas[0]
}

func volume(dimensions []int) int {
	return dimensions[0] * dimensions[1] * dimensions[2]
}

func main() {
	file, err := os.Open(os.Args[1])
	check(err)
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "x")
		dimensions := make([]int, 3)
		for i := range data {
			dimensions[i], err = strconv.Atoi(data[i])
			check(err)
		}
		// area, smallestSideArea := area(dimensions)
		sort.Ints(dimensions)
		result += 2*dimensions[0] + 2*dimensions[1] + volume(dimensions)
	}

	check(scanner.Err())
	fmt.Println(result)
}
