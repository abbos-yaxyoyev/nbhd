package datetime

import "time"

func Generate() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}
