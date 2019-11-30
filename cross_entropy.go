package main

import (
	"fmt"
	"math"
	"os"
)
var labels []string = []string{"pig","cow","car","bus"}
var yTrue []float64 = []float64{1,2,3,4}
var yPred []float64 = []float64{0.65,0.15,0.25,0.5,0.35,0.85,0.75,0.5}
var yProbs [][]float64 = [][]float64{
									[]float64{1,0,0,0},
									[]float64{.3,.05,.55,.1},
									[]float64{.5,.0,.0,.5},
									[]float64{.0,.35,0,.65},
									[]float64{.5,.0,.0,.5},
}

func simpleMean(arr []float64) float64 {
	/*Returns the mean of an array of float64 values*/
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	var mean = sum / float64(len(arr))
	return mean
}
/*
The crossEntropy function will return the negative log-likelihood based
on the predicted probability of the target class
*/
func crossEntropy(yTrue []float64, yPred []float64) []float64{
	var totalCE []float64
	for index, value := range yTrue {
		fmt.Println("Preds:", yPred[index])
		fmt.Println(value)
		y := value
		p := yPred[index]

		var CE = -(y * math.Log(p) + (1-y) * math.Log(1-p))
		fmt.Println("CE: ", CE)
		totalCE = append(totalCE, CE)
	}
	//fmt.Println(mat.Sum(totalCE))
	var meanCE = simpleMean(totalCE)
	fmt.Println("Cross Entropy: ", meanCE)
	return yPred
}

func crossEntropy1(yTrue []float64, yProbs [][]float64, labels []string) ([]float64) {

	var uniqueClasses = unique(yTrue)

	if len(uniqueClasses) <= 1 {
		fmt.Println(uniqueError())
		os.Exit(3)
	}

	var totalCE []float64
	for index, value := range yTrue {
		//fmt.Println("Preds:", yPred[index])
		//fmt.Println(value)
		y := value
		p := yPred[index]

		var CE = -(y * math.Log(p) + (1-y) * math.Log(1-p))
		//fmt.Println("CE: ", CE)
		totalCE = append(totalCE, CE)
	}
	//fmt.Println(mat.Sum(totalCE))
	var meanCE = simpleMean(totalCE)
	fmt.Println("Cross Entropy: ", meanCE)
	return yPred
}

func main() {
	//slice := [...][]int{}


	arr := [3][2]float64{}
	copyDimensions(arr)
	fmt.Println(arr)
	//fmt.Println(yProbs)
	//crossEntropy(yTrue, yPred)
	//crossEntropy1(yTrue, yProbs, labels)
	//u := unique(yTrue)
	//fmt.Println(u)
	//fmt.Println(uniqueError())

	clip(yProbs, 1e-15, 1-1e-15)
	//declareArray(2)
	//fmt.Println(arr)
	//elems := 8

}
