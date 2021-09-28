package main

import (
	"flag"
	"fmt"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func main() {
	name := flag.String("name", "Alice Bob", "Name to generate key for")
	email := flag.String("email", "alice@example.com", "Email to generate key for")
	password := flag.String("password", "123456", "Password to protect key with")

	generate := flag.Bool("generate", false, "Generate a key")

	flag.Parse()

	if *generate {
		key, err := helper.GenerateKey(*name, *email, []byte(*password), "x25519", 0)
		if err != nil {
			panic(err)
		}

		fmt.Println(key)
	}
}
