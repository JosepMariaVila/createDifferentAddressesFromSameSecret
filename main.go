package main

import (
	"fmt"

	"example.com/go/rkey"
)

const secret = "sp6JS7f14BuwFY8Mw6bTtLKWauoUs"

func main() {
	s, _ := rkey.NewFamilySeed(secret)
	pubkey := s.PrivateGenerator.PublicGenerator.Generate(2)
	addr := pubkey.Address()
	fmt.Println("secret:", secret, "address:", addr)
}
