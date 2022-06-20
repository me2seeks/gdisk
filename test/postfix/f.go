package main

import (
	"cloud-disk/common/globalkey"
	"fmt"
)

func main() {
	fmt.Println(globalkey.Postfix["APP"])
}
