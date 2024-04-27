package main
import "fmt"

func main(){
	fmt.Print("Numbers divisible by 3 between 1 and 100 are: ",)
	for i := 1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Print(i, " ")
		}
	}
}