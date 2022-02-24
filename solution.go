package square

import (
	"fmt"
	"math"
)
type calculator int

const (
	SidesSquare   = 4
	SidesTriangle = 3
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum calculator) float64 {
	if calculator = SidesSquare {
		return math.Pow(sideLen, 2)
	}
	fmt.Println(CalcSquare(10.0))

	if calculator = SidesTriangle {
		return (sideLen * sideLen) * ((math.Sqrt(3)) / 4)
	}
	fmt.Println(CalcSquare(10.0))

	if calculator = SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	}
	fmt.Println(CalcSquare(10.0))
} else
return 0
}


