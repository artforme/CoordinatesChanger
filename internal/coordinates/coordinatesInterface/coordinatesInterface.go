package coordinatesInterface

type CoordinatesSystem interface {
	Info() string
	New(args ...float64)
	ConvertIntoCartesian() CoordinatesSystem
	ConvertIntoCylindrical() CoordinatesSystem
	ConvertIntoSpherical() CoordinatesSystem
	ConvertIntoPolar() CoordinatesSystem
}
