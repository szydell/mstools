package mstools

import (
	"log"
	"math"
	"os"
)

//Round function to round, with precision 2
func Round(input float64) float32 {
	var round float64
	pow := math.Pow(10, float64(2))
	digit := pow * input
	round = math.Ceil(digit)
	return float32(round / pow)
}

//errCheck log and die on error
func errCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// does file exist?
func fExists(fName string) (bool, error) {
	if _, err := os.Stat(fName); err == nil {
		return true, err
	}
	return false, err
}
