package slicer

func Average(data []float64) float64 {

	if len(data) == 0 {
		return 0.0
	}

	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	total := 0.0
	for _, v := range data {
		total += v
	}

	return total / float64(len(data))
}
