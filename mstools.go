package mstools

import (
	"log"
	"math"
	"os"
)

//Round function to round, with precision 2
func Round(input float64) float64 {
	var round float64
	pow := math.Pow(10, 2)
	digit := pow * input
	round = math.Ceil(digit)
	return round / pow
}

//errCheck log and die on error
func errCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// does file exist?
func fExists(fName string) (exist bool , err error) {
	exist = false
	if _, err := os.Stat(fName); err == nil {
		exist = true
	}
	return
}
