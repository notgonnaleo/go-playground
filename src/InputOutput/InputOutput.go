package main
import "fmt"

func main(){
	LoadUserInfo();
}

func LoadUserInfo(){
	var name string
	var age int

	fmt.Printf("What's your name?")
	fmt.Scan(&name)

	fmt.Printf("How old are you?")
	fmt.Scan(&age)

	ShowUserInfo(name, age)
}

func ShowUserInfo(name string, age int){
	fmt.Printf("My name is %s and I'm %d years old.", name, age)
}

/*
* %s for strings
* %d for integers
* %f for floating-point numbers
* %v for the default format of any value
* for string interpolation always use Printf.
*/
