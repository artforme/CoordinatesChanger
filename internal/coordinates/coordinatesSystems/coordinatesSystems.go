package coordinatesSystems

import (
	coord "CoordinatesChanger/internal/coordinates/coordinatesInterface"
	"fmt"
	"math"
)

type Cylindrical struct {
	Radius float64
	Phi    float64
	Z      float64
}

func (c Cylindrical) Info() string {
	return "CYLINDRICAL"
}

func (c *Cylindrical) String(decimalPlaces uint) string {
	format := fmt.Sprintf("Radius: %% .%df, Phi: %% .%df, Z: %% .%df", decimalPlaces, decimalPlaces, decimalPlaces)
	return fmt.Sprintf(format, c.Radius, c.Phi, c.Z)
}

func (c *Cylindrical) New(args ...float64) {
	c.Radius = args[0]
	c.Phi = args[1]
	c.Z = args[2]
}

func (c *Cylindrical) ConvertIntoCartesian() coord.CoordinatesSystem {
	cartesian := Cartesian{
		X: c.Radius * math.Cos(c.Phi),
		Y: c.Radius * math.Sin(c.Phi),
		Z: c.Z,
	}
	return &cartesian
}

func (c *Cylindrical) ConvertIntoCylindrical() coord.CoordinatesSystem {
	return c
}

func (c *Cylindrical) ConvertIntoSpherical() coord.CoordinatesSystem {
	return c
}

type Spherical struct {
	Radius float64
	Theta  float64
	Phi    float64
}

func (s Spherical) Info() string {
	return "SPHERICAL"
}

func (s *Spherical) String(decimalPlaces uint) string {
	format := fmt.Sprintf("Radius: %% .%df, Theta: %% .%df, Phi: %% .%df", decimalPlaces, decimalPlaces, decimalPlaces)
	return fmt.Sprintf(format, s.Radius, s.Theta, s.Phi)
}

func (s *Spherical) New(args ...float64) {
	s.Radius = args[0]
	s.Theta = args[1]
	s.Phi = args[2]
}

func (s *Spherical) ConvertIntoCartesian() coord.CoordinatesSystem {
	cartesian := Cartesian{
		X: s.Radius * math.Sin(s.Theta) * math.Cos(s.Phi),
		Y: s.Radius * math.Sin(s.Theta) * math.Sin(s.Phi),
		Z: s.Radius * math.Cos(s.Theta),
	}
	return &cartesian
}

func (s *Spherical) ConvertIntoCylindrical() coord.CoordinatesSystem {
	return s
}

func (s *Spherical) ConvertIntoSpherical() coord.CoordinatesSystem {
	return s
}

type Cartesian struct {
	X float64
	Y float64
	Z float64
}

func (c Cartesian) Info() string {
	return "CARTESIAN"
}

func (c *Cartesian) String(decimalPlaces uint) string {
	format := fmt.Sprintf("X: %% .%df, Y: %% .%df, Z: %% .%df", decimalPlaces, decimalPlaces, decimalPlaces)
	return fmt.Sprintf(format, c.X, c.Y, c.Z)
}

func (c *Cartesian) New(args ...float64) {
	c.X = args[0]
	c.Y = args[1]
	c.Z = args[2]
}

func (c *Cartesian) ConvertIntoCartesian() coord.CoordinatesSystem {
	return c
}

func (c *Cartesian) ConvertIntoCylindrical() coord.CoordinatesSystem {
	cylindrical := Cylindrical{
		Radius: math.Sqrt(c.X*c.X + c.Y*c.Y),
		Phi:    math.Atan2(c.Y, c.X),
		Z:      c.Z,
	}
	return &cylindrical
}

func (c *Cartesian) ConvertIntoSpherical() coord.CoordinatesSystem {
	radius := math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z)
	if radius == 0 {
		return &Spherical{Radius: 0, Theta: 0, Phi: 0}
	}
	spherical := Spherical{
		Radius: radius,
		Theta:  math.Atan(math.Sqrt(c.X*c.X+c.Y*c.Y) / c.Z),
		Phi:    math.Atan2(c.Y, c.X),
	}
	return &spherical
}
