package other3

import (
	"os"
	"fmt"
	"errors"
)

func ExportFromOther3OfThirty() {

	_, error := os.Open("non_existed_file")
	if error != nil {
		fmt.Println("file is not existed")
		fmt.Println(error.Error())
		//return // if there is no return, the following code will be executed
	}

	if _, err := returnWithError(-3); err != nil {
		fmt.Println(err.Error())
	}

	if _, err := returnWithError(0); err != nil {
		fmt.Println(err.Error())
	}

	var ln = "abs"
	tc := newTeach(nil, &ln)
	fmt.Println(tc)

}

func returnWithError(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("the input cannot be lower than 0")
	} else if i == 0 {
		return 0, fmt.Errorf("the input %d is wrong\n", i)

	} else {
		return i + 1, nil
	}
}

type teacher struct {
	fn string
	ln string
}

func rec() {
	r := recover()
	if r != nil {
		fmt.Println("recovered!")

	}
}

// only pointer or struct can be nil; int/string cannot be set to nil
func newTeach(fn *string, ln *string) (teacher) {
	defer rec()
	if fn == nil {
		panic("firstname cannot be nill")
	}

	if ln == nil {
		panic("lastname cannot be nill")
	}
	return teacher{
		fn: *fn,
		ln: *ln,
	}
}
