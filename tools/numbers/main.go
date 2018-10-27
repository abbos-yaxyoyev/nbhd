package numbers

import (
	"fmt"
	"strconv"
)

func TruncateFloat(n float64, p int) float64 {
	i := fmt.Sprintf("%."+strconv.Itoa(p)+"f", n)
	r, _ := strconv.ParseFloat(i, p)
	return r
}
