
package main
import (
	"strings"
	) 
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s);
	var splitted = strings.Split(s, " ");
	return len(splitted[len(splitted)-1])

}

func main() {
	var s = "a"
	println(lengthOfLastWord(s))
}