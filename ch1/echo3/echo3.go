//Same as echo1 and 2 but uses Join for effeciency
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Name of command:" + os.Args[0]) //exercise 1.1
	fmt.Println(strings.Join(os.Args[1:], " "))
	//Also fmt.Println(os.Args[1:]) would work if formatting doesn't matter
}