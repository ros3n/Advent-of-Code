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
	res := 0
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)
	for i := range data {
		if data[i] == '(' {
			res += 1
		} else if data[i] == ')' {
			res -= 1
			if res == -1 {
				fmt.Println(i)
				break
			}
		}
	}
	fmt.Println(res)
}
