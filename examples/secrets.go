package main

import (
	"fmt"

	"github.com/infisical/infisical-go"
)

func main() {
	client := infisical.InfisicalClient("st.6427edc1bb17b13ef6cfb41b.72c099963eee6c414972744ff216f0e0.8804617c6779b5f34af9edb71c95a862")
	fmt.Println(client)
}
