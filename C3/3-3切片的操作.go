package main

import "fmt"

func main() {
	sliceappend()

	sliceops()
}

func sliceappend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Extending slice")
	fmt.Println(arr)
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
}

func sliceops() {
	fmt.Println("Creating slice")
	var s []int //zreo value for slice is nil

	for i := 0; i < 100; i++ {
		printslice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printslice(s1)

	s2 := make([]int, 16)

	s3 := make([]int, 10, 32) //type, len, cap
	printslice(s2)
	printslice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printslice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printslice(s2)

}

func printslice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
