package main

import (
	"fmt"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		panic("env empty")
	}
	fmt.Println(env)
	return
}
