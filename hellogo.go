package main

import (
	"bufio"
	"fmt"
	"os"
)

var pl = fmt.Println

func main() {
	pl("What is your name? ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	pl("Hello", name)
}
