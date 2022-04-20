package main

import "fmt"

func main() {
	var hen [6]float64
	hen[0] = 3
	hen[1] = 4
	hen[2] = 8
	hen[3] = 6
	hen[4] = 30
	hen[5] = 2.5
	henNum := len(hen)
	var total float64
	for _, value := range hen {
		total += value
	}
	avg := total / float64(henNum)
	fmt.Printf("total=%v,avg=%.2f", total, avg)
}
