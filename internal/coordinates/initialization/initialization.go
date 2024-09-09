package initialization

import (
	coord "CoordinatesChanger/internal/coordinates/coordinatesInterface"
	sysCoor "CoordinatesChanger/internal/coordinates/coordinatesSystems"
	"fmt"
	"strings"
)

func Init(coordinatesSystem string) coord.CoordinatesSystem {
	var (
		first  float64
		second float64
		third  float64
	)
	coordinatesSystem = strings.ToUpper(coordinatesSystem)
	switch coordinatesSystem {

	case "CARTESIAN":
		var cartesian sysCoor.Cartesian
		fmt.Println("Enter the x, y and z coordinates:")
		fmt.Scan(&first, &second, &third)
		cartesian.New(first, second, third)
		return &cartesian

	case "CYLINDRICAL":
		var cylindrical sysCoor.Cylindrical
		fmt.Println("Enter the radius, angle and z coordinates:")
		fmt.Scan(&first, &second, &third)
		cylindrical.New(first, second, third)
		return &cylindrical

	case "SPHERICAL":
		var spherical sysCoor.Spherical
		fmt.Println("Enter the radius, theta and phi coordinates:")
		fmt.Scan(&first, &second, &third)
		spherical.New(first, second, third)
		return &spherical

	default:
		return nil
	}

}
