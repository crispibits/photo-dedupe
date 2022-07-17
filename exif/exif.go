package exif

import (
	"os"
	"time"

	goexif "github.com/rwcarlsen/goexif/exif"
)

// Retrieve the timestamp from a given filename
func GetTimeStamp(fname string) (timestamp time.Time, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}

	x, err := goexif.Decode(f)
	if err != nil {
		return
	}

	timestamp, err = x.DateTime()
	return

}
