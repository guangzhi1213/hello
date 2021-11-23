package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello. codeup.")
	envs := os.Environ()
	fmt.Println(envs)
}
