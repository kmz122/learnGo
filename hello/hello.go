package main

import (
	"fmt"

	"github.com/kmz122/greetings"
	"rsc.io/quote/v4"
)

func greeting(name string) string {
	return fmt.Sprintf("Hello %s! Happy learning Go.", name)
}

func main() {
	fmt.Println(greeting("Kyi Myo Zaw"))
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Kyi"))
}
