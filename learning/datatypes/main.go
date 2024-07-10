package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Single bool
}

func star() {
	var ankush = User{"Ankush", "ankushsingh.dev@gmail.com", 20, true}
	fmt.Println("User is: ", ankush)
}

func main() {
	fmt.Println("Hello")

	num := 23
	// var ptr *int
	var ptr = &num
	fmt.Println("Value of ptr is: ", ptr)
	fmt.Println("Value of ptr is: ", *ptr)
	fmt.Printf("Value of ptr is: %T ", ptr)

	var vegList [3]string
	vegList[0] = "Potato"
	vegList[1] = "Tomato"
	vegList[2] = "cauliflower"
	fmt.Println("VegList is: ", vegList)
	fmt.Println("VegList is: ", len(vegList))

	var fruitsList = [3]string{"Apple", "Banana", "Orange"}
	fmt.Println("Fruits List is: ", fruitsList)

	fmt.Printf("Datatype of fruitslist is %T \n", fruitsList)

	var stuList = []string{}
	stuList = append(stuList, "Ankush", "Abhay", "Harshit")
	fmt.Println("Student List is: ", stuList)

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	fmt.Println("Languages are: ", languages)
	fmt.Println("PY stands for: ", languages["PY"])

	delete(languages, "PY")
	fmt.Println("Languages are: ", languages)

	abhay := User{"Abhay", "abhay@go.dev", 19, true}
	fmt.Println("User is: ", abhay)
	fmt.Printf("Details about Abhay: %+v \n", abhay)

	fmt.Printf("Email of Abhay is %v and his age is %v.\n", abhay.Email, abhay.Age)

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

	star()
}
