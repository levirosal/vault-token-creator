package main

import (
	"fmt"
	"os"

	"github.com/levirosal/vault-token-creator/cmd"
)

var (
	policy string
	ttl    string
)

func main() {
	if len(os.Args)%2 == 0 || len(os.Args) == 1 {
		panic("Missing arguments.")
	}

	for i := 1; i+1 < len(os.Args); i += 2 {
		policyArg := os.Args[i]
		ttlArg := os.Args[i+1]

		if policyArg[0:8] == "-policy=" && ttlArg[0:5] == "-ttl=" {
			policy = policyArg[8:len(policyArg)]
			ttl = ttlArg[5:len(ttlArg)]

			if policy == "" || ttl == "" {
				panic("Missing arguments. Policy and ttl cannot be null.")
			}

			fmt.Println("Creating ...")
			cmd.Create(policy, ttl)
		} else {
			println("Wrong arguments: ")
			println(policyArg)
			println(ttlArg)
		}
		fmt.Println("-----------------------------------")
	}
}
