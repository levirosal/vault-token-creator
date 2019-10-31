package main

import (
	"fmt"
	"os"
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

		if policyArg[0:8] == "-policy=" && ttlArg[0:5] == "-TTL=" {
			policy = policyArg[8:len(policyArg)]
			ttl = ttlArg[5:len(ttlArg)]
			fmt.Println("Creating ......")
			//cmd.Create(policy, ttl)
			println(policy)
			println(ttl)
			fmt.Println("-----------------------------------")
		} else {
			panic("Wrong arguments.")
		}
	}
}
