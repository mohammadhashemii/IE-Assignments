package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {

	var langName string
	var m int
	fmt.Scanf("%s %d", &langName, &m)

	reader := bufio.NewReader(os.Stdin)
	feat := make([]string, m)
	for i := 0; i < m; i++ {
		feature, err := reader.ReadString('\n')
		if err != nil {
			panic("cannot read from stdin")
		}
		feature = strings.TrimSpace(feature)
		feat[i] = strings.ToLower(feature)
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		panic("cannot read from stdin")
	}
	var n int
	n, err = strconv.Atoi(strings.TrimSpace(line))

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		feature, err := reader.ReadString('\n')
		if err != nil {
			panic("cannot read from stdin")
		}
		feature = strings.TrimSpace(feature)
		if contains(feat, strings.ToLower(feature)) {
			ans[i] = "yes"
		} else {
			ans[i] = "no"
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}

}
