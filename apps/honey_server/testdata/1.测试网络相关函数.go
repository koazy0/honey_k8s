package main

import (
	"fmt"
)

func main() {

	user := ""
	_, err := fmt.Scanln(&user)
	if err != nil {
		fmt.Println(err)
	}
}
