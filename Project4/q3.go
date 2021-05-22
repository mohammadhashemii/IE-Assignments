package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Course struct {
	Name 		string 		`json:"name"`
	Id 			int	   		`json:"id"`
	Unit 		int	   		`json:"unit"`
	Teacher		string		`json:"teacher"`
	TAs			[]string	`json:"TAs"`
}

func (c *Course) String() string {
	return fmt.Sprintf("%s (%d)\nby %s", c.Name, c.Unit, c.Teacher)
}

func main() {

	var num_lines int
	fmt.Scanf("%d", &num_lines)

	_json := ``
	scanner := bufio.NewScanner(os.Stdin)
	for i:= 0; i < num_lines; i++ {
		scanner.Scan()
		_json += scanner.Text()
	}
	course := Course{}
	json.Unmarshal([]byte(_json), &course)
	fmt.Print(course.String())
}
