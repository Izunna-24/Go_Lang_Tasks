package main	

import "fmt"

func main(){
	fmt.Print("Enter degree in F: ")
	var fareinheit int
	fmt.Scanln(&fareinheit)

	celcuis := (fareinheit - 32) * 5/9

	fmt.Println("The equivalence of the temp you entered in celcuius is ", celcuis)
}