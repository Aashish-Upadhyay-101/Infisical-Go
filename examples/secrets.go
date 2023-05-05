package main

import (
	"fmt"

	"github.com/infisical/infisical-go"
)

func main() {
	a := infisical.SecretBundle{
		SecretName:  "API_KEY",
		SecretValue: "whatisthis",
	}

	fmt.Println(a)
}
