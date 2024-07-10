package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var username string = "Ankush"
	fmt.Println("My name is " + username + " Singh")
	fmt.Printf("Name is type of %T \n", username)

	name := "Section"
	fmt.Println(name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	userName, _ := reader.ReadString('\n')
	fmt.Println("Hello " + userName + " Welcome to the Go Programming Language!")

	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please give us a ratin between 1 and 5: ")
	rating, _ := reader.ReadString('\n')

	addedRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New Rating is:", addedRating+1)
	}

	time := time.Now()
	fmt.Println("Current Time is: ", time)

	fmt.Println(time.Format("01-02-2006 Monday 15:04:05"))
}
