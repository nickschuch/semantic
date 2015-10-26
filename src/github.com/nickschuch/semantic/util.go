package main

import (
	"fmt"
	"os"
)

func Exit(m string) {
	fmt.Println(m)
	os.Exit(1)
}
