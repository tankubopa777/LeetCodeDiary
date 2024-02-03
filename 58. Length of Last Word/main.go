package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)



func main() {
    
	fmt.Println("Enter the low value")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	low, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Enter the high value")
	scanner.Scan()
	high, _ := strconv.Atoi(scanner.Text())
	fmt.Println(sequentialDigits(low, high))

}
	
	


func sequentialDigits(low int, high int) []int {
	var result []int
	var num int
	for i := 1; i < 9; i++ {
		
	}
}