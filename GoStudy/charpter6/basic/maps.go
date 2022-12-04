package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name": "zhangsan",
		"sex":  "man",
	}

	m2 := make(map[string]string)

	var m3 map[string]int

	m4 := new(map[string]string) //PS: &map[]

	fmt.Println(m1, m2, m3, m4)
	name, ok := m1["name"] //ok 判断是否存在 true false
	fmt.Println("\n", name, ok)
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println("\n", name, ok)
	fmt.Println(m1, m2, m3, m4)
}
