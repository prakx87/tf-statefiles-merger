package main

import (
	"fmt"
	"flag"
)


func main() {
	var source_file_1 string
	var source_file_2 string
	

	flag.StringVar(&source_file_1, "source1", "", "Path for first source file")
	flag.StringVar(&source_file_2, "source2", "", "Path for second source file")

	flag.Parse()

	fmt.Println(source_file_1)
	fmt.Println(source_file_2)
}