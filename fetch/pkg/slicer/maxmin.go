package slicer

func Max(data []float64) float64 {

	if len(data) == 0 {
		return 0.0
	}
	m := data[0]
	for _, v := range data {
		if m < v {
			m = v
		}
	}
	return m
}

func Min(data []float64) float64 {

	if len(data) == 0 {
		return 0.0
	}
	m := data[0]
	for _, v := range data {
		if m > v {
			m = v
		}
	}
	return m
}
