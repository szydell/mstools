package mstools

import (
	"log"
	"math"
	"os"
)

//Round function to round, with precision 2
func Round(input float64) float64 {
	var round float64
	digit := input * 100
	round = math.Ceil(digit)
	return round / 100
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
