package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// get a string as input from the user
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
    	panic("cannot read from stdin")
	}
	inp = strings.TrimSpace(inp)


	total_ascii := 0
	for i := 0; i < len(inp); i++ {
		total_ascii += int(inp[i])
	}
	// check if the sum of the ascii code is less than 200 or not
	if total_ascii < 200 {
		fmt.Println("err")
		return
	}
	for total_ascii > 399 {
		total_ascii /= 2
	}
	fmt.Println(total_ascii)

}