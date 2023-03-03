package statistics

import "sort"

// Median calculates the median of a slice of float64 values
func Median(data []float64) float64 {
    sort.Float64s(data)
    n := len(data)
    if n%2 == 0 {
        // If the number of elements is even, average the middle two
        return (data[n/2-1] + data[n/2]) / 2
    } else {
        // If the number of elements is odd, return the middle element
        return data[n/2]
    }
}
