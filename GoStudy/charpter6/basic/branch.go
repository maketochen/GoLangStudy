package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	const filename = "basic/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}
