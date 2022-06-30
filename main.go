package main

import "fmt"

func main() {
	out := MultiplicationTable4(3)
	fmt.Println(out)
}

//func MultiplicationTable(size int) [][]int {
//	out := make([][]int, size)
//	number := 1
//	for i := 0; i < size; i++ {
//		for j := 1; j < size+1; j++ {
//			out[i] = append(out[i], j*number)
//		}
//		number++
//	}
//	return out
//}
//
//func MultiplicationTable2(size int) [][]int {
//	out := make([][]int, size)
//	for i := 0; i < size; i++ {
//		for j := 1; j < size+1; j++ {
//			out[i] = append(out[i], j*i+1)
//		}
//	}
//	return out
//}
//
//func MultiplicationTable3(size int) [][]int {
//	out := make([][]int, size+1)
//	for i := 0; i < size; i++ {
//		for j := 1; j < size; j++ {
//			out[i] = append(out[i], j*i+1)
//		}
//	}
//	return out
//}

func MultiplicationTable4(size int) [][]int {
	out := make([][]int, size)
	for i := range out {
		for j := range out {
			out[i] = append(out[i], j*i)
		}
	}
	return out
}
