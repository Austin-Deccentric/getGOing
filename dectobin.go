package main

import(
	"fmt"
	"strconv"
)

func conv(n int) string{
	return strconv.FormatInt(int64 (n), 2)
}

func main() {
	var s string
	//var i int
	fmt.Println("Enter a postive Number:")
	for{
		_, err := fmt.Scan(&s)
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else{
			fmt.Println(conv(i))
			break
		}
	}
}