package mstools

import (
	"log"
	"math"
	"os"
)

//Round
func Round(val float64, places int ) (float64) {
    var round float64
    pow := math.Pow(10, float64(places))
    digit := pow * val
    _, div := math.Modf(digit)
    if div >= 0.5 {
        round = math.Ceil(digit)
    } else {
        round = math.Floor(digit)
    }
    return round / pow
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
