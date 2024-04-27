package main
import "fmt"



func main(){
	fmt.Print("Enter a name: ")
	var inputCollected string
	fmt.Scanln(&inputCollected)
	var result string = "My name is "+ inputCollected
	fmt.Println(result)


	

	
}

