package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama string = "Andika Andriana"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(nama))

	var sha = sha1.New()
	sha.Write([]byte(encodedString))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println("Loading Variabel..")
	fmt.Println("Original\t\t:", nama)
	fmt.Println("Original -> Encoded\t:", encodedString)
	fmt.Println("Encoded -> Hashed\t:", encryptedString)
	fmt.Println("Proses Selesai..")
}
