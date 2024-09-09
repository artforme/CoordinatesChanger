package main

import (
	"CoordinatesChanger/internal/coordinates/coordinatesInterface"
	"CoordinatesChanger/internal/coordinates/initialization"
	"fmt"
	"strings"
)

func main() {
	// Run the main function of the application
	fmt.Println("Hello, World!")

	var (
		tempSystemCoordinates string
	)
	fmt.Println("Enter the coordinates system you want to start with:")
	fmt.Println("CARTESIAN, CYLINDRICAL, SPHERICAL")
	fmt.Scan(&tempSystemCoordinates)

	var systemsCoordinatesObject coordinatesInterface.CoordinatesSystem

	systemsCoordinatesObject = initialization.Init(tempSystemCoordinates)

	if systemsCoordinatesObject == nil {
		fmt.Println("Invalid coordinates system")
		return
	}

	fmt.Println("The coordinates system you have chosen is:", systemsCoordinatesObject.Info())
	var decimalPlaces uint
	fmt.Println("Enter the number of decimal places:")
	fmt.Scan(&decimalPlaces)
	fmt.Println("The coordinates are:", systemsCoordinatesObject.String(decimalPlaces))

	for {
		fmt.Println("Enter the coordinates system you want to convert to:")
		fmt.Println("CARTESIAN, CYLINDRICAL, SPHERICAL")
		fmt.Scan(&tempSystemCoordinates)
		tempSystemCoordinates = strings.ToUpper(tempSystemCoordinates)
		switch tempSystemCoordinates {
		case "CARTESIAN":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian()
		case "CYLINDRICAL":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian().ConvertIntoCylindrical()
		case "SPHERICAL":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian().ConvertIntoSpherical()
		default:
			fmt.Println("finishing the program")
			return
		}
		fmt.Println("The coordinates system you have chosen is:", systemsCoordinatesObject.Info())
		fmt.Println("The coordinates are:", systemsCoordinatesObject.String(decimalPlaces))
	}

}
