package main

import (
	"fmt"

	"github.com/karalabe/go-init-bug/bad"
	"github.com/karalabe/go-init-bug/good"
)

func main() {
	fmt.Println("Good order:", good.Size((*good.B)(nil)))
	fmt.Println("Bad order: ", bad.Size((*bad.B)(nil)))
}
