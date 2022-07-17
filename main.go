package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/crispibits/photo-dedupe/exif"
	"github.com/crispibits/photo-dedupe/md5"
)

var md5Sums map[string][]string

func init() {
	md5Sums = make(map[string][]string)
}

func main() {

	// TODO - use Viper or other cmdline parsing
	dir := os.Args[1]

	fmt.Printf("Looking in %s\n", dir)

	// TODO - change to WalkDir
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		fmt.Printf("Checking %s\n", path)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			timestamp, err := exif.GetTimeStamp(path)
			if err != nil {
				return nil
			}
			y, m, d := timestamp.Date()
			moveFile(path, y, int(m), d, md5.MD5File(path))
		}
		return nil
	})
	for k, v := range md5Sums {
		if len(v) > 1 {
			fmt.Printf("%s %s\n", k, v)
		}
	}
}

//func doMd5Sum(k string, v string) {
//	md5Sums[k] = append(md5Sums[k], v)
//}

func moveFile(from string, y int, m int, d int, md5 string) {
	yDir := fmt.Sprintf("%d", y)
	if !exists(yDir) {
		os.Mkdir(yDir, 0755)
	}
	mDir := fmt.Sprintf("%d/%02d", y, m)
	if !exists(mDir) {
		os.Mkdir(mDir, 0755)
	}
	dDir := fmt.Sprintf("%d/%02d/%02d", y, m, d)
	if !exists(dDir) {
		os.Mkdir(dDir, 0755)
	}
	// TODO - set suffix dependent on file type
	to := fmt.Sprintf("%s/%s.jpg", dDir, md5)
	fmt.Printf("Moving %s to %s\n", from, to)
	if err := os.Rename(from, to); err != nil {
		log.Fatal(err)
	}
}

func exists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
