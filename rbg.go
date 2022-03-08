package chigo

import (
	"fmt"
	"math"
)

type RGB struct {
	Red, Green, Blue int
}

func NewRGB(i float64) RGB {
	return RGB{
		int(math.Sin(0.1*i)*127 + 128),
		int(math.Sin(0.1*i+2*math.Pi/3)*127 + 128),
		int(math.Sin(0.1*i+4*math.Pi/3)*127 + 128),
	}
}

func (r RGB) Color(s string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r.Red, r.Green, r.Blue, s)
}
