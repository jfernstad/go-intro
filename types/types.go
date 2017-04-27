package main

// Multiple imports using ()
import (
	"fmt"
)

// Golang basic types
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias for uint8
//
// rune // alias for int32
//      // represents a Unicode code point
//
// float32 float64
//
// complex64 complex128

// Two new separate types
type MyType int
type someType int

// Capitalized `types`, `functions`, `methods`, `variables` are Exported, meaning Public
// Non-capitalized are Not Exported, meaning Private, as in other packages can not reach them directly

type MyCoolObject struct { // Object declaration
	PublicMember  string
	privateMember int
}

func main() {

	var myInt uint16 // full variable declaration. Yup, type AFTER variable name
	myFloat := 0.2   // implicit variable declaration `:=`, type is inferred

	// Object instantiation
	someObj := MyCoolObject{}             // type inferred
	someObj = MyCoolObject{"Hi There", 1} // using same variable, someObj is garbage collected, no memory leak
	someObj = MyCoolObject{PublicMember: "Hello", privateMember: 1}

	var myAry [3]int   // int-array with 3 slots
	var someAry [5]int // this is a completely different type
	var slice []int    // A `slice` is not an array

	myAry = [3]int{1, 2, 3} // Yup, weird. Creates and initializes a new array
	slice = someAry[1:2]    // A `slice` points to a backing array, `[start_idx:end_idx]`

	// This is an error since someAry is of type `[5]int`, not `[3]int`
	// someAry = [3]int{1, 2, 3}

	// This is a swap
	myAry[0], myAry[1] = myAry[1], myAry[0]

	// Not using variables IS AN ERROR
	fmt.Printf("%d, %f, %+v, %v, %v, %v\n", myInt, myFloat, someObj, myAry, someAry, slice)
	// Prints: 0, 0.200000, {PublicMember:Hello privateMember:1}, [0 0 0], [0 0 0 0 0], [0]
}

// https://tour.golang.org/moretypes/1
