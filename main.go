package main

import (
	"fmt"
	"math"
	"strconv"
	"awesomeProject/other"
	"awesomeProject/other2"
	"awesomeProject/other3"
	"reflect"
)

func main() {

	fmt.Println("start to test")

	// anoymous function
	func(n int) int {
		fmt.Println(n)
		return n + n
	}(4)

	func(af func(n int)) {
		fmt.Println(af)
	}(func(n int) {
		fmt.Println(n)

	})

	// defer
	// defer stack: FILO
	var deferVariable = 1
	var deferVariable2 = 2
	defer fmt.Printf("defer variable is the value when write it! %d \n", deferVariable)
	defer fmt.Printf("defer variable is the value when write it! %d \n", deferVariable2)
	deferVariable = 3
	deferVariable2 = 4
	fmt.Printf("defer variable is the value %d \n", deferVariable2)

	var a int
	var b string
	a = 4
	b = "avs"
	fmt.Println("a is", a, "b is", b)


	// reflect
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))






	c, f := 14, 18 // shorthand
	const ci int = 43
	var (
		d int = 4
		e int = 3
	)

	fmt.Println("result is", c+d-e*f+ci)
	fmt.Println(math.Min(float64(c), float64(f)))

	var boola bool = true
	if boola {
		var asix int64 = -232
		var athree int32 = -14
		var ausix uint64 = 232
		var authree uint32 = 14
		fmt.Println("result is", asix+int64(athree)+int64(ausix)+int64(authree))

		var af float64 = 3.3
		var anotherInt int64 = 4
		fmt.Println(af + float64(anotherInt))

		fmt.Println("stringfy result is", strconv.Itoa(int(anotherInt))) // stringfy method
	}

	var sum, minor int = ope(4, 5)
	fmt.Println(sum + minor)

	var sum2, _ int = ope(4, 5)
	fmt.Println(sum2)

	// function programming
	fmt.Println(big(4, func(a int) int {
		return a + 100
	}))

	fmt.Println(other.FuncFromOther(other.VariableFromOther))

	other.Export()
	other2.ExportFromOther2()
	other2.Export2FromOther2()
	other3.ExportForConcurrency()
	//other3.ExportForConcurrencyWorker()
	other3.ExportFromOther3OfThirty()

}

func ope(a int, b int) (int, int) {
	c := math.Max(float64(a), float64(b))
	return a * int(c), a - int(c)
}

type newType func(a int) int

func big(a int, small newType) int {
	return small(a)
}
