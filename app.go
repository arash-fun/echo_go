package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("arr is %T", arr)

	a := make([]int, 2, 12)
	fmt.Println(a)

	b := arr[2:5]
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)

	b[1] = 12345
	fmt.Println(arr)

	matric()

}
func test2(myfunc func(int) int) {
	fmt.Println("this function do like this ....")
	myfunc(165423)
}

func abs(z int) int {
	if z > 0 {
		return z
	}
	return -z
}

func matric() {
	sum := 0
	var matrix [2][2]int
	fmt.Println("enter your matric \n")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Element : matric[%d][%d]", i, j)
			fmt.Scan(&matrix[i][j])

		}
	}
	fmt.Println(matrix)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum += matrix[i][j]
		}
	}
	fmt.Printf("sum matrix is %d", sum)
}
