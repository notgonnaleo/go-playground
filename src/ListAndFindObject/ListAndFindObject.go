package main
import "fmt"

type PersonType struct {
	Id int
	Name string
	Age int
}

var Id int;

var Persons = []PersonType{
	{
		Id: 1,
		Name: "Leonardo",
		Age: 20,
	},
	{
		Id: 2,
		Name: "Samuel",
		Age: 26,
	},
	{
		Id: 3,
		Name: "Guilherme",
		Age: 21,
	},
}

func List(){
	for _, person := range Persons{
		fmt.Println(person.Id, person.Name, person.Age)
	}
}

func Find(Id int) PersonType{
	for _, person := range Persons{
		if person.Id == Id{
			return person
		}
	}
	return PersonType{}
}

func main(){
	List()

	fmt.Println("Select one person:")
	fmt.Scanf("%d", &Id)

	foundPerson := Find(Id)
	fmt.Println(foundPerson)
}