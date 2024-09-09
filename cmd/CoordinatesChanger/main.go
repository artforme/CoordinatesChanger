package main

import (
	"CoordinatesChanger/internal/coordinates/coordinatesInterface"
	"CoordinatesChanger/internal/coordinates/initialization"
	"fmt"
	"strings"
)

func main() {
	// Run the main function of the application
	var tempSystemCoordinates string
	systemsCoordinatesObject := InitCoordinates(tempSystemCoordinates)

	fmt.Println("The coordinates system you have chosen is:", systemsCoordinatesObject.Info())
	var decimalPlaces uint
	fmt.Println("Enter the number of decimal places:")
	fmt.Scan(&decimalPlaces)
	fmt.Println("The coordinates are:", systemsCoordinatesObject.String(decimalPlaces))

	for {
		fmt.Println("Enter the coordinates system you want to convert to:")
		fmt.Println("CARTESIAN type 1, CYLINDRICAL type 2, SPHERICAL type 3, Change the coordinates type 4")
		fmt.Scan(&tempSystemCoordinates)
		tempSystemCoordinates = strings.ToUpper(tempSystemCoordinates)
		switch tempSystemCoordinates {
		case "1":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian()
		case "2":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian().ConvertIntoCylindrical()
		case "3":
			systemsCoordinatesObject = systemsCoordinatesObject.ConvertIntoCartesian().ConvertIntoSpherical()
		case "4":
			systemsCoordinatesObject = InitCoordinates(tempSystemCoordinates)
		default:
			fmt.Println("finishing the program")
			return
		}
		fmt.Println("The coordinates system you have chosen is:", systemsCoordinatesObject.Info())
		fmt.Println("The coordinates are:", systemsCoordinatesObject.String(decimalPlaces))
	}

}

func InitCoordinates(tempSystemCoordinates string) coordinatesInterface.CoordinatesSystem {
	fmt.Println("Enter the coordinates system you want to start with:")
	fmt.Println("CARTESIAN, CYLINDRICAL, SPHERICAL")
	fmt.Scan(&tempSystemCoordinates)

	var systemsCoordinatesObject coordinatesInterface.CoordinatesSystem

	systemsCoordinatesObject = initialization.Init(tempSystemCoordinates)

	if systemsCoordinatesObject == nil {
		fmt.Println("Invalid coordinates system")
		return nil
	}
	return systemsCoordinatesObject
}
