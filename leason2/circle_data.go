package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var circleArea float64

	fmt.Println("enter circle eare")
	_, err := fmt.Scanf("%f\n", &circleArea)
	if err != nil {
		println("Fscanf: ", err.Error())
		os.Exit(0)
	}

	radius, circumference := getCircleDataByArea(circleArea)

	println("Radius: ", fmt.Sprintf("%.6f", radius))
	println("circumference: ", fmt.Sprintf("%.6f", circumference))
}

/**
 * Return information about radius and circumference by area
 *
 */
func getCircleDataByArea(area float64) (float64, float64) {

	radius := math.Sqrt(area / math.Pi)
	circumference := 2 * radius * math.Pi

	return radius, circumference
}
