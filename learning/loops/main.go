package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the content to write in the file: ")
	content, _ := reader.ReadString('\n')

	file, err := os.Create("./test.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./test.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Data read from the file is:\n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// func main() {

// 	var i int

// 	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// 	for i = 0; i < 5; i++ {
// 		println(days[i])
// 	}

// 	for d := range days {
// 		println(d, days[d])
// 	}

// 	var val int = 0

// 	for val < 10 {
// 		if val == 5 {
// 			val++
// 			goto lco
// 		} else {
// 			fmt.Println(val)
// 		}
// 		val++
// 	}
// lco:
// 	fmt.Println("Hello")
// }
