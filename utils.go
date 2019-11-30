package main

import (
	"errors"
	"fmt"
)

func uniqueError() error {
	return errors.New("a minimum of two classes required. please enter correct ground truth values")
}

func copyDimensions (matrix [][]float64) {
	duplicate := make([][]float64, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]float64, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	fmt.Println(duplicate)
}

type Matrix [][]float64

func clip(arr Matrix, min float64, max float64) [5][4]float64 {
	/*
	Function to clip upper and lower bounds of values between 0 and 1.
	Useful for situations requiring Log functions which can return inf. or NAN
	values.
	Inputs: arr: 2D array of float 64s e.g. [[0,1],[0.5,0.5]]
	        min: float64, minimum value for returned elements
			max: float64, maximum value for returned elements
	Output: outputArray: array with same dimensionality as input array
	*/
	// Only working for explicit dimensions.
	fmt.Println("Length:", len(arr))

	var outputArray [5][4]float64
	var replacement float64
	for i, value := range arr {
		for j, element := range value {
			if element < min {
				replacement = min
			} else if element > max {
				replacement = max
			} else {
				replacement = element
			}
			outputArray[i][j] = replacement
		}
		fmt.Println(outputArray)

	}
	return outputArray
}

func unique(arr []float64) []float64{
	/*
	Functions takes a 1-dimensional array and returns an
	array containing a single instance of each unique element.
	Input: arr: 1D-array(float64)
	Output: unique: 1D-array(float64)
	*/
	var unique []float64
	for _, i := range arr{
		skip := false
		for _, j := range unique{
			if i == j{
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, i)
		}
	}
	return unique
}

//func declareArray()  {
//	elems := 8
//	var slice = make([...][]int,elems)
//
//	fmt.Println(slice)
//
//}