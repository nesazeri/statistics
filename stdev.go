package statistics

import "math"

// Stdev calculates the sample standard deviation of a slice of float64s.
func Stdev(data []float64) float64 {
    if len(data) < 2 {
        return math.NaN()
    }

    mean := Mean(data)
    var variance float64
    for _, x := range data {
        variance += math.Pow(x-mean, 2)
    }
    variance /= float64(len(data) - 1)

    return math.Sqrt(variance)
}

// PopStdev calculates the population standard deviation of a slice of float64s.
func PopStdev(data []float64) float64 {
    if len(data) < 2 {
        return math.NaN()
    }

    mean := Mean(data)
    var variance float64
    for _, x := range data {
        variance += math.Pow(x-mean, 2)
    }
    variance /= float64(len(data))

    return math.Sqrt(variance)
}
