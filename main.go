package main
import (
	"bufio"
	"strings"
	"os"
	"fmt"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		incoming := scanner.Text()
		lowercase := strings.ToLower(incoming)
		fmt.Fprintln(os.Stdout,lowercase)
	}
}
