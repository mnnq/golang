package main

//we use import to get access to some code is written inside another package
//give my package main access to all of the code all the func inside fmt
//standard package fmt= format, print difference information
import "fmt"

func main() {

	s := make([]int, 11)
	for i, _ := range s {
		s[i] = i
	}

	for i, a := range s {
		if a%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

/*go build just compiles the program
go fmt foramts all the code each file
go test runs any tests associated with current project

reusable

executable: when you execute the file creates an executable type
reusable: libraries or code dependencies, reusable logic, etc

the name of the packagew that you use creates an executable or reusable (the word package main
makes Go create a executable)
*/
