package main

import (
	"cloud-disk/common/oss"
	"fmt"
)

func main() {
	test := oss.UploadCertificate(1024)
	fmt.Println(test)
}
