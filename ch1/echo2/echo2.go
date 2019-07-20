//Echo2 prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("Index %d, Arg: %s\n", index, arg) //exercise 1.2
	}
}