package other

import (
	"fmt"
)

func init() {
	fmt.Println("override the init of oth.go")
}

var VariableFromOther int = 555

func FuncFromOther(a int) int {

	switch {
	case a > 5:
		fmt.Println("a>5")
	default:
		fmt.Println("a!>5")
	}

	switch b := a + 4; b {
	case 4, 5, 6:
		fmt.Println("a==456")
	default:
		fmt.Println("a is not 456")
	}

	if a > 10 {
		fmt.Println("a is larger than 10")
	} else if a == 10 {
		fmt.Println("a is 10")
	} else {
		fmt.Println("a is less thjna 10")
	}

	if i := 3; i > 0 {
		fmt.Println("init i is larger than 0");
	}

	for i := 1; i < 3; i++ {
		a++;
	}

	i := 4;
	for ; i <= 6; i++ { // while style
		if i == 6 {
			break
		}
		a++;
	}

	return a + 5;
}
