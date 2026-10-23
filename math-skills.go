package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func average(nums []float64) float64 {
	var sum float64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / float64(len(nums))
}

func median(nums []float64) float64 {
	numbers := make([]float64, len(nums))
	copy(numbers, nums)
	var temp float64
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if j+1 != len(numbers) {
				if numbers[j+1] < numbers[j] {
					temp = numbers[j+1]
					numbers[j+1] = numbers[j]
					numbers[j] = temp
				}
			}
		}
	}
	if len(numbers)%2 == 0 {
		mid := len(numbers) / 2
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[len(numbers)/2]
}

func variance(nums []float64) float64 {
	avr := average(nums)
	var res float64
	var sum float64
	for i := 0; i < len(nums); i++ {
		res = avr - nums[i]
		res *= res
		sum += res
	}
	return sum / float64(len(nums))
}

func std(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <inputfile>")
		return
	}

	inputFile := os.Args[1]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(data)
	lines := strings.Split(input, "\n")
	var numbers []float64

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid number in file:", line)
			return
		}

		numbers = append(numbers, float64(value))
	}
	var x float64
	x = variance(numbers)
	fmt.Println("Average:", int(math.Round(average(numbers))))
	fmt.Println("Median:", int(math.Round(median(numbers))))
	fmt.Println("Variance:", int(math.Round(variance(numbers))))
	fmt.Println("Standard Deviation:", int(math.Round(std(x))))

}
