package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumNums[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Call the SumInts function
	fmt.Printf("SumInts: %v\n", SumInts(ints))

	// Call the SumFloats function
	fmt.Printf("SumFloats: %v\n", SumFloats(floats))

	fmt.Printf("Generic Sum (Int64) & (Float64): %v and %v\n",
		SumNums(ints),
		SumNums(floats),
	)
}
