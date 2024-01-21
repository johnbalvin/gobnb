package search

import "time"

type InputData struct {
	Coordinates CoordinatesInput
	Check       Check
	ZoomValue   int
}
type Check struct {
	In  time.Time
	Out time.Time
}

type CoordinatesInput struct {
	Ne CoordinatesValues
	Sw CoordinatesValues
}
type CoordinatesValues struct {
	Latitude float64
	Longitud float64
}
