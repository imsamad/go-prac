package main;

import "fmt"


func main() {
	// way to init slices
	// 1. [len]type{...elems...}
	// 2. from array
	// 3. make function
	one := []int{1,2,3}
	two := [...]int{1,2,3}

	fmt.Println(cap(one))
	fmt.Println(len(one))
	
	fmt.Println(one)	
	fmt.Println(two)	

	// from array

	arr := []int{1,2,3,4,5,6}
	three := arr[3:4]

	fmt.Println(arr)
	fmt.Println(three)
	fmt.Println(cap(three)) // important to remember cap would from stat till end of array
	fmt.Println(len(three))

	len_ := 10
	cap_ := 15
	four := make([]int, len_, cap_)
	fmt.Println(four)

	// ways to modify

	// specific elem
	slice := []int{1,2,3,4,5}
	slice[1] = 111
	fmt.Println(slice)

	// append individual elems
	slice = append(slice, 1,2,3,4,5)
	fmt.Println(slice)

	slice2 := []int{9,8,7}

	// append slice to another slice
	slice2 = append(slice, slice2...)

	fmt.Println(slice2)

	// Change lemgth
	fmt.Printf("===============================\n")
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5] // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1)) // cap would still remain origLen - startOfSlice
	
	fmt.Printf("===============================\n")
	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	copySliceMemoryEff()
}

func copySliceMemoryEff() {
	fmt.Println("copySliceMemoryEff")
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
  
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Printf("neededNumbers = %v\n", neededNumbers)
	fmt.Printf("length = %d\n", len(neededNumbers))
	fmt.Printf("capacity = %d\n", cap(neededNumbers))

	// numbersCopy:=[len(neededNumbers)]int{}
	// l
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
  
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
	// The capacity of the new slice is now less than the capacity of the original slice because the new underlying array is smaller.
}