package main

import (
	"fmt"
	"math/cmplx"
)

func loop(num int) int {
    sum := 0 
    for i := 0; i < num; i++ { 
        sum += i 
    }
    return sum 
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	arr1 := []interface{}{1, 2, 3, "pranesh"}
	arr2 := [5]interface{}{4,5,6,7,8}
  
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr1)


}




day := 5

switch day {
case 1,3,5:
	fmt.Println("Odd day")
case 2,4,6:
	fmt.Println("Even day")
default:
	fmt.Println("iNVALID day")
}