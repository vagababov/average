package main

import (
	"fmt"
	"math/rand"
	"time"
)

func average(d []float64) float64 {
	s := 0.
	for _, f := range d {
		s += f
	}
	return s / float64(len(d))
}

// d[0] presumed the newest data for simplicity.
// linear weighted average s = (n*d[0] + (n-1)*d[1]...+1*d[n-1])/sum(1..n)
func weightedAvg(d []float64) float64 {
	s := 0.
	sz := float64(len(d))
	l := len(d) - 1
	for i := range d {
		fi := float64(i)
		s += (sz - fi) * d[l-i]
	}
	return s / (sz * (sz + 1) / 2)
}

// exponentially decaying weights for the average.
// we're using 2N as total sum of weights, and approximating 1/2 as step.
func expAvg2N(d []float64, step float64) float64 {
	s := 0.
	//	sz := float64(len(d))
	m := step
	l := len(d) - 1
	for i := range d {
		s += m * d[l-i]
		m *= (1 - step)
		//	fmt.Printf("S = %v M = %v\n", s, m)
	}
	return s
}

func randData(num int, min, max float64) []float64 {
	if min >= max {
		panic("min >= max")
	}
	ret := make([]float64, num)
	d := max - min
	for i := 0; i < num; i++ {
		ret[i] = rand.Float64()*d + min
	}
	return ret
}

func increasingData(num int, min, max float64) []float64 {
	if min >= max {
		panic("min >= max")
	}
	ret := make([]float64, num)
	step := (max - min) / float64(num)
	for i := 0; i < num; i++ {
		ret[i] = min + float64(i)*step
	}
	return ret
}

func decreasingData(num int, min, max float64) []float64 {
	if min >= max {
		panic("min >= max")
	}
	ret := make([]float64, num)
	step := (max - min) / float64(num)
	for i := 0; i < num; i++ {
		ret[i] = max - float64(i)*step
	}
	return ret
}

func randomStats() {
	rand.Seed(time.Now().UnixNano())
	ds := randData(200, 1, 3600)
	fmt.Print("num_run,avg,linavg")
	for s := 0.1; s < 1; s += 0.1 {
		fmt.Printf(",expavg(%v)", s)
	}
	fmt.Println()

	for i := 60; i < len(ds); i++ {
		d := ds[i-60 : i]
		fmt.Printf("%d,%v,%v", i-60, average(d), weightedAvg(d))
		for s := 0.1; s < 1; s += 0.1 {
			fmt.Printf(",%v", expAvg2N(d, s))
		}
		fmt.Println()
	}
	fmt.Println()
}

func increasingStats() {
	rand.Seed(time.Now().UnixNano())
	ds := increasingData(200, 1, 3600)
	fmt.Print("num_run,avg,linavg")
	for s := 0.1; s < 1; s += 0.1 {
		fmt.Printf(",expavg(%v)", s)
	}
	fmt.Println()

	for i := 60; i < len(ds); i++ {
		d := ds[i-60 : i]
		fmt.Printf("%d,%v,%v", i-60, average(d), weightedAvg(d))
		for s := 0.1; s < 1; s += 0.1 {
			fmt.Printf(",%v", expAvg2N(d, s))
		}
		fmt.Println()
	}
	fmt.Println()
}

func decreasingStats() {
	rand.Seed(time.Now().UnixNano())
	ds := increasingData(200, 1, 3600)
	fmt.Print("num_run,avg,linavg")
	for s := 0.1; s < 1; s += 0.1 {
		fmt.Printf(",expavg(%v)", s)
	}
	fmt.Println()

	for i := 60; i < len(ds); i++ {
		d := ds[i-60 : i]
		fmt.Printf("%d,%v,%v", i-60, average(d), weightedAvg(d))
		for s := 0.1; s < 1; s += 0.1 {
			fmt.Printf(",%v", expAvg2N(d, s))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	randomStats()
	//	increasingStats()
}
