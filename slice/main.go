package main

import "fmt"

func main() {
	fmt.Println("initial array")
	arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)

	fmt.Println("get slice 1:4 - len = 3 and capacity = 5 (because slice starts not from first element of initial array)")
	sl1 := arr[1:4]
	fmt.Println(sl1)

	fmt.Println("add elem to slice - it replaced next element in initial array")
	sl1 = append(sl1, 7)
	fmt.Println(sl1)
	fmt.Println(arr)

	fmt.Println("add elem to slice - it replaced next element in initial array")
	sl1 = append(sl1, 8)
	fmt.Println(sl1)
	fmt.Println(arr)

	fmt.Println("change first element in slice - analogical element in array was changed too")
	sl1[0] = 9
	fmt.Println(sl1)
	fmt.Println(arr)

	fmt.Println("add more elems to slice - slice was copied and now it's pointer also points to new array")
	fmt.Println("so, we don't see any changes in initial array")
	sl1 = append(sl1, 10, 11, 12)
	fmt.Println(sl1)
	fmt.Println(arr)
}
