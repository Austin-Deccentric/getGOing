package main

import(
	"fmt"
	"strconv"
)
// Conv converts the user input to equivalnet binary representation
func conv(n int) string{
	return strconv.FormatInt(int64 (n), 2) 
}

func main() {				
	var s string 				//Initialize a variable of type string
	fmt.Println("Enter a postive Number:")
	for{
		_, err := fmt.Scan(&s) 		//Read from standard input using Scan and a pointer. _, err because i want to capture any error
		i, err := strconv.Atoi(s) 	// Converts the input to an interger
		if err != nil {           	// Checks for err in the conversion and repeats the stdio operation
			fmt.Println("Enter a valid number")
		} else{
			fmt.Println(conv(i)) 	//If no error is encountered, call on the conversion function: conv
			break		     	// exit the loop
		}
	}
}
