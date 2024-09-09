package coordinatesSystems

import (
	coord "CoordinatesChanger/internal/coordinates/coordinatesInterface"
	"math"
)

type Cylindrical struct {
	Radius float64
	Angle  float64
	Z      float64
}

func (c Cylindrical) Info() string {
	return "CYLINDRICAL"
}

func (c *Cylindrical) New(args ...float64) {
	c.Radius = args[0]
	c.Angle = args[1]
	c.Z = args[2]
}

func (c *Cylindrical) ConvertIntoCartesian() coord.CoordinatesSystem {
	cartesian := Cartesian{
		X: c.Radius * math.Cos(c.Angle),
		Y: c.Radius * math.Sin(c.Angle),
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

func (c *Cylindrical) ConvertIntoPolar() coord.CoordinatesSystem {
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

func (s *Spherical) New(args ...float64) {
	s.Radius = args[0]
	s.Theta = args[1]
	s.Phi = args[2]
}

func (s *Spherical) ConvertIntoCartesian() coord.CoordinatesSystem {
	cartesian := Cartesian{
		X: s.Radius * math.Sin(s.Phi) * math.Cos(s.Theta),
		Y: s.Radius * math.Sin(s.Phi) * math.Sin(s.Theta),
		Z: s.Radius * math.Cos(s.Phi),
	}
	return &cartesian
}

func (s *Spherical) ConvertIntoCylindrical() coord.CoordinatesSystem {
	return s
}

func (s *Spherical) ConvertIntoSpherical() coord.CoordinatesSystem {
	return s
}

func (s *Spherical) ConvertIntoPolar() coord.CoordinatesSystem {
	return s
}

type Polar struct {
	Radius float64
	Angle  float64
	Z      float64
}

func (p Polar) Info() string {
	return "POLAR"
}

func (p *Polar) New(args ...float64) {
	p.Radius = args[0]
	p.Angle = args[1]
	p.Z = args[2]
}

func (p *Polar) ConvertIntoCartesian() coord.CoordinatesSystem {
	cartesian := Cartesian{
		X: p.Radius * math.Cos(p.Angle),
		Y: p.Radius * math.Sin(p.Angle),
		Z: p.Z,
	}
	return &cartesian
}

func (p *Polar) ConvertIntoCylindrical() coord.CoordinatesSystem {
	return p
}

func (p *Polar) ConvertIntoSpherical() coord.CoordinatesSystem {
	return p
}

func (p *Polar) ConvertIntoPolar() coord.CoordinatesSystem {
	return p
}

type Cartesian struct {
	X float64
	Y float64
	Z float64
}

func (c Cartesian) Info() string {
	return "CARTESIAN"
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
		Radius: math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z),
		Angle:  math.Atan2(c.Y, c.X),
		Z:      c.Z,
	}
	return &cylindrical
}

func (c *Cartesian) ConvertIntoSpherical() coord.CoordinatesSystem {
	radius := math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z + c.Z)
	spherical := Spherical{
		Radius: radius,
		Theta:  math.Atan2(c.Y, c.X),
		Phi:    math.Acos(c.Z / radius),
	}
	return &spherical
}

func (c *Cartesian) ConvertIntoPolar() coord.CoordinatesSystem {
	polar := Polar{
		Radius: math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z),
		Angle:  math.Atan2(c.Y, c.X),
		Z:      c.Z,
	}
	return &polar
}
