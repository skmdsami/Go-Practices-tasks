package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func enums_example(day Day) {
	if Friday == day {
		fmt.Println("start of the week-end")
	}else{
		fmt.Println("it's not an start of the week-end")
	}
}

func main() {
	age := 22 // there should be no un-used variable in the program, shorthand of declaring a variable
	var name , last_name string
	name = "phani"
	last_name = "phani"
	fmt.Println(name, last_name)
	var age1 uint = 22
	const age2 uint = 23 // it will not throw error because these are checked at compile , if these are not used
	var animal string = "phani"
	fmt.Println("animal was",animal)
	fmt.Println(age1)
	fmt.Println(age)
	fmt.Println("This is a Go Program")

	for {
		fmt.Println("this keeps on iterating , so we need to mention a base condition to stop it")
		break
	}

	for i:=0 ; i < 12; i++ {
		fmt.Println("This is the common way of doing it ")
	}

	i := 0
	for i <= 10 {
		i += 1
	}// we can use the same for keyword for doing while , loop, normal for loop operations and we also have range as well

	for j := range 10 { // starts with zero and ends at 10
		if j%2 == 0 {
			continue
		}
	}
	number := 2
	switch number {
	case 2 :
		fmt.Println("even")
	fmt.Println("number-2")
	case 3 : fmt.Println("odd")
	default:
		fmt.Println("Other number")
	}


	fmt.Println("Let's get Started with Arrays")

	var array[5] uint ;
	array[0] = 1;

	array2 := [2]uint{1,1}
	fmt.Println(array2)
	array[1] = 2
	array[2] = 5
	array[3] = 6
	array[4] = 12
	fmt.Println("These are the 2 ways of creating an array in GO")
	fmt.Println("slice of the array was",array[:3]) // 0 to up to 3rd index means 0,1,2 indexes values will be printed out
	namee := [5]string {"phani"}
	fmt.Println(namee)

	// make(Type, length [, capacity])
	mapp := make(map[string]int)
	mapp["phani"] = 21
	delete(mapp, "phani")

	sum :=  add(2,3)
	fmt.Println("sum of 2,3 was", sum)
	multiplication, subraction := math_calculations(5,2)
	fmt.Println("for values 2,3 multiplication and subraction was", multiplication, subraction)
	person_struct_call()
	enums_example(Friday)
}


func GenericAdd[T int | uint | float64 | float32](a T, b T) T {
	return a + b
}


type numericTypes interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func genericAdd[T numericTypes] (a T, b T) T {
	return a + b
}

// or instead of doing it by yourselfes we can import the constrains.Ordered and we can do [T constrains.Ordered] such that all numeric times will be
// done "golang.org/x/exp/constraints
/*
import "golang.org/x/exp/constraints"

func add[T : constraints.Ordered](a T,b T) T {...}
*/

func add(a int, b int ) int {
	return	a + b
}

func math_calculations(a int, b int) (int, int) {
	return a*b, a-b
}


type person struct {
	name string
	age uint8
}

func person_constructor(name string, age uint8) person {
	return person { name, age}
}

func (p person) get_name() string { // this is my way for implementing functions to a struct
	return p.name
}

func (p person) get_age() uint8 {
	return p.age
}


func person_struct_call() {
	var p person = person_constructor("phani", 21)

	fmt.Println("name :",p.get_name())
	fmt.Println("age :",p.get_age())
}

//* Note- Go uses Pascal case and Camel case naming conventions