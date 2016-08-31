package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(A []float64, B []float64) float64 {
	var b float64
	for i := 0; i < len(A); i++ {
		b = b + A[i]*B[i]
	}
	return b
}

func updateWeight(weight []float64, sample []float64, c float64) []float64 {
	W := make([]float64, len(weight))
	for i := 0; i < len(weight); i++ {
		W[i] = weight[i] + c*sample[i]
	}
	return W
}

// var threshold float64 = 2.5 // change me

type TrainingSample struct {
	Sample []float64
	Label  int
}

type PerceptronTrainer struct {
	Samples      []TrainingSample
	LearningRate float64
	Weight       []float64
	Threshold    float64
}

func (p *PerceptronTrainer) Init() error {
	if len(p.Samples) == 0 {
		return fmt.Errorf("No sample loaded.")
	}
	length := len(p.Samples[0].Sample)
	p.Weight = make([]float64, length)
	// More initial values
	return nil
}

func (p *PerceptronTrainer) LoadSamplesFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newSample := TrainingSample{
			Sample: make([]float64, 0),
		}
		rawData := scanner.Text()
		stringdata := strings.Split(rawData, ",")
		for j := 0; j < len(stringdata)-1; j++ {
			value, err := strconv.ParseFloat(strings.Trim(stringdata[j], " "), 64)
			if err != nil {
				return err
			}
			newSample.Sample = append(newSample.Sample, value)
		}
		label, err := strconv.Atoi(strings.Trim(stringdata[len(stringdata)-1], " "))
		if err != nil {
			return err
		}
		newSample.Label = label
		p.Samples = append(p.Samples, newSample)
	}
	return nil
}

func (p *PerceptronTrainer) Train() int {
	errCount := 0
	for _, sample := range p.Samples {
		result := sum(sample.Sample, p.Weight)
		y := 0
		if result > p.Threshold {
			y = 1
		} else {
			y = -1
		}
		if y != sample.Label {
			errCount++
			c := p.LearningRate * float64(sample.Label-y)
			p.Weight = updateWeight(p.Weight, sample.Sample, c)
		}
	}
	return errCount
}

func (p *PerceptronTrainer) TrainToEnd(timeout int) int {
	for i := 0; i < timeout; i++ {
		errCount := p.Train()
		if errCount == 0 {
			return i
		}
	}
	return timeout
}
