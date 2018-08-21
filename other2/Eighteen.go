package other2

import (
	"fmt"
	"strconv"
)

type Callable interface {
	call() string
	yell(ppl string) string
}

// use callable and callable3 as two interfaces for the same people concrete methods
type Callable3 interface {
	call() string
}

type Callable4 interface {
	scream() string
}

// combine two interfaces together as one whole
type CallableBig interface {
	Callable
	Callable4
}

type people struct {
	name string
}

func (ppl people) call() string {
	return ppl.name
}

func (ppl people) yell(na string) string {
	return na + " " + ppl.name
}

type people2 struct {
	name string
	age  int
}

func (ppl *people2) call() string {
	return ppl.name + " " + strconv.Itoa(int(ppl.age))
}

func (ppl *people2) yell(na string) string {
	return na + " " + ppl.name + " " + strconv.Itoa(int(ppl.age))
}

// take any argument by using empty interface
func takeAnyArgu(i interface{}) {

	switch i.(type) {
	case string:
		{
			fmt.Println("it is string")
		}

	default:
		fmt.Printf("type is %T and value is %v \n", i, i)

	}
}

func Export2FromOther2() {
	var ppl people = people{
		name: "kobe",
	}
	var callable Callable = ppl
	fmt.Println(callable.call())
	fmt.Println(callable.yell("ahaha"))

	var ppl2 people2 = people2{
		name: "kobe2",
		age:  32,
	}
	var callable2 Callable = &ppl2
	fmt.Printf(callable2.call())
	fmt.Printf(callable2.yell("ahaha2"))

	var ih interface{} = "hello"
	takeAnyArgu(ih)
	takeAnyArgu(callable2)

	var callable3 Callable3 = ppl
	fmt.Println(callable3.call())

}
