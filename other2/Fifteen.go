package other2

import (
	"fmt"
)

func ExportFromOther2() {

	mp := make(map[int]*string)
	// map and slice have been the reference, so no need to have the pointer

	var v1 string = "a1"
	mp[1] = &v1
	fmt.Println(*mp[1])

	var v2 *string = &v1
	mp[2] = v2

	fmt.Println(*mp[1] + *mp[2])

	fmt.Println(takePointer(&v1))
	fmt.Println(v1)

	// operate on the array/slice
	arr := [...]int{1, 2, 3} // slice
	fmt.Println(arr)
	arrPointer := &arr
	(*arrPointer)[0] = 100
	fmt.Println(arr)

	// type
	var per Person = Person{
		firstName: "kobe",
		lastName:  "bryant",
		age:       29,
	}

	var pp *Person = &per
	fmt.Println(pp.firstName)

	changePerson(&per)
	fmt.Println(per)

	var simple = struct {
		name string
		age  int
	}{"kb", 334}

	fmt.Println(simple)

	var emp Employee = Employee{
		title:  "sde",
		person: *pp,
	}

	fmt.Println(emp.person.firstName)

	emp.person.call("hello")
	emp.person.call2("hello")

	// use New-style manually created constructor
	var newPeron Person = New("aa", "bb", 12)
	fmt.Println(newPeron)

}

func takePointer(mystr *string) string {
	return *mystr + "!!!"
}

func changePerson(person *Person) {
	person.lastName = "new" + person.lastName
	person.firstName = "new" + person.firstName

}

type Person struct {
	firstName, lastName string
	age                 int
}

// the method of the Person
func (p Person) call(name string) {
	fmt.Printf("%s %s %d\n", name+" "+p.firstName, p.lastName, p.age)
}

// manually created New-style constructor
func New(fn string, ln string, ag int) Person {
	return Person{
		firstName: fn,
		lastName:  ln,
		age:       ag,
	}
}

// the method of the Person
func (p *Person) call2(name string) {
	fmt.Printf("%s %s %d\n", name+" "+p.firstName, p.lastName, p.age)
}

// type composition
type Employee struct {
	title  string
	person Person
}
