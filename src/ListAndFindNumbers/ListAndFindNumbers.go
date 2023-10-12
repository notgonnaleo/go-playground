package main
import "fmt"

var array [5]int = [5]int{10,45,21,4,9}
var Index int

func main(){
	List(array)
	
	fmt.Printf("Pull one item from list:")
	fmt.Scanf("%d", &Index)
	fmt.Print(Find(Index, array))
}

func List(array[5]int){
	for i, v := range array {
		fmt.Println(i,v)
	}
}

func Find(Index int, array[5]int) int{
	return array[Index]
}