package main

import "fmt"

func main() {
	s1 := []int{1,2,3}
	fmt.Println(len(s1),cap(s1),s1) // выводит 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2),cap(s2),s2) // выводит 2 2 [2 3]

	for i := range s2 { s2[i] += 20 }

	// всё ещё ссылается на тот же массив
	fmt.Println(s1) // выводит [1 22 23]
	fmt.Println(s2) // выводит [22 23]

	s2 = append(s2,4)

	for i := range s2 { s2[i] += 10 }

	//s1 is now "stale"
	fmt.Println(s1) // выводит [1 22 23]
	fmt.Println(s2) // выводит [32 33 14]
}
