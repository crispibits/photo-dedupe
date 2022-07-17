package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func MD5File(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	//Get the 16 bytes hash
	hashInBytes := h.Sum(nil)[:16]

	//Convert the bytes to a string
	return hex.EncodeToString(hashInBytes)
}
