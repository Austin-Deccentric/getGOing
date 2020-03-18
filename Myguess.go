package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	// generate a random number
	rand.Seed(time.Now().UnixNano())
	secretNumbers :=[] int64{2,3,5,8,12,16,7}
	Number := secretNumbers[rand.Intn(len(secretNumbers))]
	var guess int64
	for{
		fmt.Println("\nPlease Input your guess:")
		fmt.Scan(&guess)
	if guess < Number{
		fmt.Println("Guess is too small")
	} else if guess > Number{
		fmt.Println("Guess is too BIG")
	} else {
		fmt.Println("You win")
		break
	}
	}
}
