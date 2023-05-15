package utils

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Take in lat,long in a string style coordinates and output two float64's for lat and long instead
func ProcessCoords(coords string) (float64, float64, error) {
	// Split input coords on comma
	co := strings.Split(coords, ",")

	// Convert latitude to float64
	lat, err := strconv.ParseFloat(co[0], 64)
	if err != nil {
		fmt.Println("Cannot parse coordinates, are they entered in the correct format?")
		log.Fatal(err)
		return 0, 0, err
	}

	// Convert longitude to float64
	long, err := strconv.ParseFloat(co[1], 64)
	if err != nil {
		fmt.Println("Cannot parse coordinates, are they entered in the correct format?")
		log.Fatal(err)
		return 0, 0, err
	}

	return lat, long, nil
}

// Check that only a single flag is in use. Returns bools for which flag
// flag has been set as coords, city
func CheckFlags(coordFlag string, cityFlag string) (bool, bool, error) {
	// Set default values
	runCity := true
	runCoords := true

	// Check if coords flag was set
	if coordFlag == "" {
		runCoords = false
	}

	// Check if city flag was set
	if cityFlag == "" {
		runCity = false
	}

	// User must set only 1 flag, not none, not both
	if !runCity && !runCoords {
		return false, false, errors.New("must set either the city or coords flag")
	} else if runCity && runCoords {
		return false, false, errors.New("use only one flag from city and coords")
	}

	return runCoords, runCity, nil

}
