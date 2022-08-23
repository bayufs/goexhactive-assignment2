package main

import "fmt"

func main() {
	var name = []string{
		"bayu",
		"fajar",
		"sidik",
		"nugraha",
		"muthus",
		"muhamad",
		"hilmi",
		"eka",
		"bintang",
		"mail"}

	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}
