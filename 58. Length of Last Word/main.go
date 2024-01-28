package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)



func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    scanner.Scan()
S := scanner.Text()
_ = S // to avoid unused error

// fmt.Fprintln(os.Stderr, "Debug messages...")
// convert the string to a byte array
_ = string(S)
for i := 0; i < len(S); i++ {
	fmt.Println(S[i])
	// S[i] to int
	x, _ := strconv.Atoi(string(S[i]))
	for j := 0; j < x ; j++ {
		fmt.Print("*")
	}
	
}
	
	
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				return count
			}
		} else {
			count++
		}
	}
	return count
}