package mstools

import (
	"log"
	"math"
	"os"
)

//Round float64, precision 2 digits
func Round(input float64) float64 {
	return math.Ceil(input * 100) / 100
}

//ErrCheck log and die on error
func ErrCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//FExists does file exist?
func FExists(fName string) (exist bool , err error) {
	exist = false
	if _, err := os.Stat(fName); err == nil {
		exist = true
	}
	return
}
