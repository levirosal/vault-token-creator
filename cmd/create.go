package cmd

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

var token = os.Getenv("VAULT_TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func Create(policy string, ttl string) {
	policies := []string{policy}
	client := SetUpClient()

	requestConfig := &api.TokenCreateRequest{
		Policies:  policies,
		TTL:       ttl,
		Renewable: new(bool),
	}
	secret, err := client.Auth().Token().Create(requestConfig)

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to Vault API.")
	}

	fmt.Println("key                Value ")
	fmt.Println("----               ------")
	fmt.Println("token              ", secret.Auth.ClientToken)
	fmt.Println("token_accessor     ", secret.Auth.Accessor)
	fmt.Println("token_duration     ", secret.Auth.LeaseDuration)
	fmt.Println("token_renewable    ", secret.Auth.Renewable)
	fmt.Println("token_policies     ", secret.Auth.TokenPolicies)
	fmt.Println("identify_policies  ", secret.Auth.IdentityPolicies)
	fmt.Println("policies           ", secret.Auth.Policies)
}

func SetUpClient() *api.Client {
	config := &api.Config{
		Address: vault_addr,
	}
	c, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}
	c.SetToken(token)
	return c
}
