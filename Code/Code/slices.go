package main

import "fmt"

func cities() {
	cities := []string{"Seattle,", "Boston,", "San Diego,"}
	fmt.Println(cities)
}

func intSlice() {
	arr := [11]int{2, 7, 11, 14, 16, 18, 20, 22, 25, 39, 43}
	slice := arr[2:9]
	fmt.Println(slice)
	subslice := arr[1:4]
	fmt.Println(subslice)
}

func lenCapSlice() {
	slice2 := make([]int, 9, 14)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	arr2 := [11]int{2, 7, 11, 14, 16, 18, 20, 22, 25, 39, 43}
	slice3 := arr2[5:9]
	fmt.Println(cap(arr2))
	fmt.Println(cap(slice3))
}

func strSlice() {
	arr3 := [7]int{2, 7, 14, 18, 20, 39, 43}
	slice4 := arr3[:4]
	fmt.Println(arr3)
	fmt.Println(slice4)
	slice4[3] = 23
	fmt.Println("Subsitute element at index 2 with a different int")
	fmt.Println(arr3)
	fmt.Println(slice4)
}

func sliceAppend() {
	arr4 := [5]int{2, 7, 14, 18, 20}
	slice5 := arr4[3:5]
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5, 39, 43)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}

func sliceAppend2() {
	arr5 := [5]int{2, 7, 14, 18, 20}
	slice6 := arr5[:3]
	arr6 := [5]int{4, 14, 28, 36, 40}
	slice7 := arr6[:3]
	slice8 := append(slice6, slice7...)
	fmt.Println(slice8)
}

func sliceAppend3() {
	arr7 := [5]int{2, 7, 14, 18, 20}
	int1 := 1
	fmt.Println(arr7)
	slice9 := arr7[:int1]
	slice10 := arr7[int1+2:]
	slice11 := append(slice9, slice10...)
	fmt.Println(slice11)
}

func sourceSlice() {
	sourceslice := []int{2, 7, 14, 18, 20}
	targetslice := make([]int, 5)
	number := copy(targetslice, sourceslice)
	fmt.Println(targetslice)
	fmt.Println("Total count of copied elements: ", number)
	arr8 := []int{2, 7, 14, 18, 20, 39, 43}
	fmt.Println("Provides the index location of each element in the array and then loops through the entire array.")
	for index, value := range arr8 {
		fmt.Println(value, " is in index", index)
	}
	for _, loop_result := range arr8 {
		fmt.Println(loop_result)
	}
}

func main() {
	cities()
	intSlice()
	lenCapSlice()
	strSlice()
	sliceAppend()
	sliceAppend2()
	sliceAppend3()
	sourceSlice()

}
