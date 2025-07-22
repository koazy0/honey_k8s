package main

import (
	"fmt"
	"honey_server/internal/logic/utils"
)

func main() {

	v := utils.IPToInt("192.168.1.2")
	fmt.Println(utils.IntToIP(v))

}
