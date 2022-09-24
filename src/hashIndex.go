package hashIndex

import (
	"fmt"
	"log"
	"os"
)

// fileSize is the maximum size of a segmentFile
const fileSize = 1024

// elemSize is the maximum allowed size of a key to be stored
const elemSize = 64

// Init prepares a hash index for use, and returns the map and data file to be used
func Init() (map[string]int, []*os.File) {
	f, err := os.Create("../files/file1.csv")
	if err != nil {
		log.Fatalf("couldn't initialize data file: %s\n", err)
	}
	dataFiles := []*os.File{f}
	m := make(map[string]int)
	return m, dataFiles
}

// Append adds or updates a key-value pair
func Append(k string, v int, segFiles []*os.File, m map[string]int, segFileIndex *int) {
	f := segFiles[*segFileIndex]
	if len(k) > elemSize {
		fmt.Printf("Failed to add key value pair to hash index: character limit exceeded (lim: %v)\n", elemSize)
		return
	}
	fs, err := f.Stat()
	if err != nil {
		fmt.Printf("failed to get location of appended info: %s\n", err)
		return
	}
	loc := int(fs.Size())
	if loc > fileSize {
		if f.Close() != nil {
			fmt.Printf("Failed to close segment file %s: %s", fs.Name(), err)
		}
		*segFileIndex++
		fNext, err := os.Create(fmt.Sprintf("../files/file%v.csv", len(segFiles)))
		if err != nil {
			fmt.Printf("Failed to create new segment file: %s", err)
			return
		}
		segFiles = append(segFiles, fNext)
		Append(k, v, segFiles, m, segFileIndex)
	}

	n, err := f.WriteString(fmt.Sprintf("%s,%v\n", k, v))
	if err != nil {
		fmt.Printf("failed to write key value pair to file: %s\n", err)
		return
	}
	fmt.Printf("%v bytes written at %v bytes of segment file %v\n", n, loc, *segFileIndex)
	m[k] = loc + *segFileIndex*fileSize // byte offset value
}

/*
// LookUp returns the associated value of a given key
func LookUp(k string, f *os.File, m map[string]int) string {
	loc := m[k]
	v := make([]byte, elemSize)
	n, err := f.ReadAt(v, int64(loc))
	fmt.Printf("Read %v bytes from segment file")
}
*/

// Write adds a key value pair to a file, and keeps the location of the data in a hash map

// compact removes duplicate entries in a data file
// func compact()

// merge aggregates two data files together
// func merge()
