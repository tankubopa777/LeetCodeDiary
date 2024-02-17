package main

import (
	"fmt"
	"strings"
)



func main() {
    
	// // covert int to string
	// var i int = 42
	// var j string = strconv.Itoa(i)
	// fmt.Println(j)

	// // convert string to int
	// var k string = "42"
	// l, _ := strconv.Atoi(k)
	// fmt.Println(l)

	// // sort string
	// var m string = "Hello"
	// n := []byte(m)
	// fmt.Println(n)

	// // sort list int
	// o := []int{1, 6, 7, 4, 5}
	// sort.Ints(o)
	// p := o
	// fmt.Println(p)

	findTheDifference("abcdefg","abcdefgf")

}

func findTheDifference(s string, t string) {
	var result string
	for _, v := range t {
		println(string(v))
		println(v)
		if strings.Contains(s, string(v)) {
			s = strings.Replace(s, string(v), "", 1)
		} else {
			result = string(v)
		}
	}
	fmt.Println(result)
}
	
	
