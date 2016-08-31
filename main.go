package main

import (
	"fmt"
	"os"
	"strconv"
)

func train(filename string, learningRate float64) {
	p := PerceptronTrainer{
		LearningRate: learningRate,
		Threshold:    1,
	}
	p.LoadSamplesFromCSV(filename)
	p.Init()
	for i := 0; i < 1000; i++ {
		errCount := p.Train()
		fmt.Printf("%d, %s, %d\n", i, JoinFloat(p.Weight, ","), errCount)
		if errCount == 0 {
			break
		}
	}
}

func trainC(filename string, learningRate float64) {
	p := PerceptronTrainer{
		LearningRate: learningRate,
		Threshold:    0,
	}
	p.LoadSamplesFromCSV(filename)
	p.Init()

	delta := 0.02
	for i := 0; i < 50; i++ {
		p.Init()
		converge := p.TrainToEnd(1000) // timeout=1000
		fmt.Printf("%.4f, %d, %s\n", p.LearningRate, converge, JoinFloat(p.Weight, ","))

		p.LearningRate = p.LearningRate + delta
		//if converge == 1000 {
		//	fmt.Printf("not converge\n")
		//}
	}
}

func generate(n int, W []float64, upperThreshold float64, lowerThreshold float64) {
	samples, labels := generateSample(n, W, upperThreshold, lowerThreshold)

	// write line to output
	for i := 0; i < n; i++ {
		fmt.Printf("%s, %d\n", JoinFloat(samples[i], ","), labels[i])
	}
}

func movie() {
	//
	// samples, labels := readDataFromFile("./samplesXOR.csv")
	// learningrate := 0.06
	//
	// countTraining, result := perceptronLearning(samples, labels, learningrate)
	//
	// fmt.Printf("%.2f, %d, %v\n", learningrate, countTraining, result)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Command not provided")
		return
	}
	command := args[0]

	switch command {
	case "generate":

	case "train":
		if len(args) <= 1 {
			fmt.Println("Filename not provided")
			return
		}
		filename := args[1]
		learningRate := 0.02
		if len(args) > 2 {
			learningRateString := args[2]
			if learningRateString != "" {
				value, err := strconv.ParseFloat(learningRateString, 64)
				if err != nil {
					fmt.Printf("Error parsing %s, using default %.2f : %s", learningRateString, learningRate, err)
				} else {
					learningRate = value
				}
			}
		}
		train(filename, learningRate)
	case "trainC":
		if len(args) <= 1 {
			fmt.Println("Filename not provided")
			return
		}
		filename := args[1]
		learningRate := 0.02
		if len(args) > 2 {
			learningRateString := args[2]
			if learningRateString != "" {
				value, err := strconv.ParseFloat(learningRateString, 64)
				if err != nil {
					fmt.Printf("Error parsing %s, using default %.2f : %s", learningRateString, learningRate, err)
				} else {
					learningRate = value
				}
			}
		}
		trainC(filename, learningRate)
	default:
		fmt.Println("Command not supported.")
	}
}
