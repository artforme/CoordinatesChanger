package coordinatesInterface

type CoordinatesSystem interface {
	Info() string
	String(decimalPlaces uint) string
	New(args ...float64)
	ConvertIntoCartesian() CoordinatesSystem
	ConvertIntoCylindrical() CoordinatesSystem
	ConvertIntoSpherical() CoordinatesSystem
}
