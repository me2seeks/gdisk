package main

import (
	"cloud-disk/common/upload"
	"fmt"
)

func main() {
	test := upload.UploadCertificate(1024)
	fmt.Println(test)
}
