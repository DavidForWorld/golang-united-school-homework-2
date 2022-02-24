package square
import "math"
type calculator int
const (
	SidesSquare   = 4
	SidesTriangle = 3
	SidesCircle   = 0
)
func CalcSquare(sideLen float64, sidesNum calculator) float64 {
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}

	if sidesNum == SidesTriangle {
		return (sideLen * sideLen) * ((math.Sqrt(3)) / 4)
	}

	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	} else {
		return 0 }
}
