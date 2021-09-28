package main

import (
	"flag"
	"fmt"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func main() {
	key := flag.String("key", "", "Key to sign with")
	password := flag.String("password", "123456", "Password of the key")
	plaintext := flag.String("plaintext", "", "Text to sign")

	flag.Parse()

	msg, err := helper.SignCleartextMessageArmored(*key, []byte(*password), *plaintext)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
}
