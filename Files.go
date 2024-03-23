package main

import (
	"fmt"
	"os"
)

func read_file(path string) string {
	dat, err := os.ReadFile(path)
	if err == nil {
		return string(dat)
	}
	fmt.Println("{}", err.Error())
	return ""
}
