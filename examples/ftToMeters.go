package main
import"fmt"

func main(){
	fmt.Println("Enter measurement in ft: ")
	var ft float64
	fmt.Scanf("%f", &ft)

	var meters float64  = ft / 0.3048

	fmt.Println("measuremnt in meters: ",meters)
}