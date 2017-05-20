package mstools

import (
    "log"
    "math"
    "os"
    "errors"
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

//FileOrDir returns if it is a file (f), or dir (d) or 'u' and an error if does not exist
func FileOrDir(fName string) (fType byte, err error) {
    fi, err := os.Stat(fName)
    if err != nil {
        return 'u', err
    }
    switch mode := fi.Mode(); {
    case mode.IsDir():
        fType = 'd'
    case mode.IsRegular():
        fType = 'f'
    default:
        err = errors.New("unkown type")
    }
    return
}
