# statistics
Statistics package in go similar to statistics in Python.
go get github.com/nesazeri/statistics

go run .\test.go

////////////////////////////////////////////////////////
package main

import (

    "fmt"
    
    "github.com/nesazeri/statistics"
    
)

func main() {

    data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    fmt.Println("Mean:", statistics.Mean(data))
	  median := statistics.Median(data)
    fmt.Println("Median:", median)
    stdev := statistics.Stdev(data)
    fmt.Println("Standard deviation:", stdev)
    
}
