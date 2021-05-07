package example

type Scalar float64

func (s Scalar) Abs() float64 {
	if s < 0 {
		return float64(-s)
	}
	return float64(s)
}
