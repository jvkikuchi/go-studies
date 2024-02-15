package main

func Sum(array []int) int {
	var sum int = 0

	for _, number := range array {
		sum += number
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	// Since slicesToSum is an array of arrays, I can access its len, 
	// and define how many items will be on the result array

	// Res 1
	// numberOfSlices := len(slicesToSum)
	// This creates a new array silice with len of the
	// numberOfSlices, initiated with {0, 0, 0}
	// sumsResults := make([]int, numberOfSlices)

	// for index, slice := range slicesToSum {
	// 	// Override the 0s with the results array from the Sum() function
	// 	sumsResults[index] = Sum(slice)
	// }

	
	// Res 2
	// Creates a new empty array, and append the result sumsResults[] end 
	var sumsResults []int

	for _, slice := range slicesToSum {
		sumsResults = append(sumsResults, Sum(slice))
	}

	return sumsResults
}