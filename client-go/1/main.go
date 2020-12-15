package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

var token = os.Getenv("TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	keyName := "secret/data/foo"
	secret, err := c.Read(keyName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("secret %s -> %v", keyName, secret)
	fmt.Println(secret.Data["name"])
}
