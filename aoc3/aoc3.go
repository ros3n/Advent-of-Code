package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)

	x, y := 0, 0
	counter := 0

	type boolmap map[int]bool
	visited := make(map[int]boolmap)

	for i := range data {
		switch {
		case data[i] == '^':
			y += 1
		case data[i] == 'v':
			y -= 1
		case data[i] == '>':
			x += 1
		case data[i] == '<':
			x -= 1
		}
		if visited[x] == nil {
			visited[x] = make(boolmap)
		}
		if !visited[x][y] {
			visited[x][y] = true
			counter += 1
		}
	}
	fmt.Println(counter)
}
