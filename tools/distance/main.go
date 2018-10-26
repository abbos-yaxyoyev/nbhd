package distance

import (
	"math"
)

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Calculate(f []float64, s []float64) float64 {

	lat1 := f[0] * math.Pi / 180
	lon1 := f[1] * math.Pi / 180
	lat2 := s[0] * math.Pi / 180
	lon2 := s[1] * math.Pi / 180

	radius := 6378100.0

	h := hsin(lat2-lat1) + math.Cos(lat1)*math.Cos(lat2)*hsin(lon2-lon1)

	r := 2 * radius * math.Asin(math.Sqrt(h))
	r = math.Round(r)

	return r
}
