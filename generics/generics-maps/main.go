package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	m := make(CustomMap[int, string])
	m[3] = "3"
	m[2] = "2"

	fmt.Printf("%+v", m)

}
