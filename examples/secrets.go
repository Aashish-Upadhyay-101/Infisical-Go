package main

import (
	"fmt"

	"github.com/infisical/infisical-go"
)

func main() {
	a := infisical.InfisicalClient("INFISICAL_SERVICE_TOKEN")
	fmt.Println(a)
}
