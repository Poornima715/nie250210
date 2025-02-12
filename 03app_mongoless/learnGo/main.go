package main
import(
	"fmt"//format module by which we can print the output
)

type Car struct {
	Id int
	Number string
	Model string
	Type string
}
func main() {
	// car1 := Car{Id:101, Number:"KA29E3031", Model:"Ambassadar", Type:"Sedan"}
	var car1 Car = Car{Id:101, Number:"KA29E3031", Model:"Ambassadar", Type:"Sedan"}
	fmt.Println(car1)
	fmt.Println(car1.Type)

	var cars[] Car = [] Car{
		{Id:103, Number:"KA29E3031", Model:"Ambassadar", Type:"Sedan"},
		{Id:104, Number:"KA28E4731", Model:"Toyota", Type:"SUV"},
		
	}
	fmt.Println(cars)
	var car2* Car= &car1
	fmt.Println(car2.Model)
	// fmt.Println("Hello World!");
	// first:= 10
	// second:= 20
	// sum:= first+second
	// var a int=10
	// var b int=20
	// var sum int =a+b
	// fmt.Println(sum)
	// var salaries[] float32=[] float32{1.0,2.0}
	// salaries:= [] float32{1.0,2.0}
	// fmt.Println(salaries)

}