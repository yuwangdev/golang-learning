package other

import (
	"fmt"
	"strings"
)

func Export() {
	fmt.Println("from other/eleven")

	// array cannot be resized

	var a [4]int
	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}

	b := [...]int{1, 2, 3, 5}
	b[2] = 4

	fmt.Println(allSrra(a))
	fmt.Println(allSrra(b))

	mat := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(mat[1][1])

	// slice, just the wrapper to the reference of the array
	var sli1 = []int{}
	sli1 = append(sli1, 1)
	sli1 = append(sli1, 2)
	fmt.Println(len(sli1))

	sli2 := b[:]
	fmt.Println("capacity is d%", cap(sli2))
	sli2 = append(sli2, 4)
	fmt.Println(sli2)
	fmt.Println("capacity is d%", cap(sli2))

	sli3 := make([]int, 4, 8) // len and cap
	if sli3 == nil {
		sli3 = append(sli3, 4)
	} else {
		sli3 = append(sli3, sli1...)
	}

	fmt.Println(sli3)

	// variadic
	fmt.Println(takeVarArray(1, 2, 3, 4, 5))

	// map
	var mmp map[int]string = make(map[int]string)
	if mmp != nil {
		mmp[1] = "a"
		mmp[2] = "ab"
		delete(mmp, 1)
		fmt.Println(mmp)
	}

	mmp2 := map[int]string{
		1: "ab",
		2: "cd",
	}

	var two, ok = mmp2[2]
	if ok {
		mmp2[2] = two + "newcd"
	}

	fmt.Println(len(mmp2))
	for key, value := range mmp2 {
		fmt.Printf("key is %d and value is %s\n", key, value)
		// only printf can replace the d s
	}

	mmp3 := mmp2
	mmp3[1] = "new!" // the same with slice, map assignment is also by reference
	fmt.Println(mmp2)

	// string
	str := "asasb"
	for i, j := range str {
		fmt.Println(i)
		fmt.Println(j) // it prints out the asic2 code
	}

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // it prints out chars
		fmt.Printf("%x", str[i]) // it prints out the asic2 code
	}

	fmt.Println("\n-----------")
	fmt.Println(strings.Compare("abc", "abc"))
	fmt.Println(strings.Count("abcd", "cd"))

}

func takeVarArray(input ...int) int {
	var a int = 0
	for _, i := range input {
		a += i
	}
	return a
}

func allSrra(array [4]int) int {

	// array length should be fixed here
	var a int

	for i, j := range array {
		a += i + j
	}
	return a

}
