package Types

import "math"

type Net struct {
	Name  string
	Zones []int
}

type Coordinate struct {
	Lat float64
	Lon float64
}

func New(x float64, y float64) Coordinate {
	return Coordinate{((2 * math.Atan(math.Exp(((((12000000-y)/2.003750834e7)*180)*3.141592653589793)/180))) - 1.5707963267948966) * 57.29577951308232, (x / 2.003750834e7) * 180}
}

type Info struct {
	Coord             Coordinate
	Elevation         int16
	Id                string
	IsTransferStation bool
	Level             int16
	Name              string
	Nets              []Net
	Omc               int
	ParentId          string
	ParentName        string
}
