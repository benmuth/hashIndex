package hashIndex

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// fileSize is the maximum size of a segmentFile
const maxFileSize = 1024

// elemSize is the maximum allowed size of a key to be stored
const elemSize = 64

// Init prepares a hash index for use, and returns the map and data file to be used
func Init() ([]map[string]int, []*os.File) {
	f, err := os.Create("../files/file1.csv")
	if err != nil {
		log.Fatalf("couldn't initialize data file: %s\n", err)
	}
	dataFiles := []*os.File{f} //TODO: change from slice of files to just opening files by name/number
	maps := []map[string]int{{}}

	return maps, dataFiles
}

// Append adds a key-value pair to the end of a segment file
func Append(k string, v int, segFiles []*os.File, maps []map[string]int, fileIndex *int) {
	f := segFiles[*fileIndex]
	if len(k) > elemSize {
		fmt.Printf("Failed to add key value pair to hash index: character limit exceeded (lim: %v)\n", elemSize)
		return
	}
	fs, err := f.Stat()
	if err != nil {
		fmt.Printf("failed to get location of appended info: %s\n", err)
		return
	}
	loc := int(fs.Size())  //loc keeps track of the end of the file
	if loc > maxFileSize { //create a new segFile and a new map that points to it
		if f.Close() != nil {
			fmt.Printf("Failed to close segment file %s: %s", fs.Name(), err)
		}
		*fileIndex++
		fNext, err := os.Create(fmt.Sprintf("../files/file%v.csv", len(segFiles)))
		if err != nil {
			fmt.Printf("Failed to create new segment file: %s", err)
			return
		}
		segFiles = append(segFiles, fNext)
		m := make(map[string]int)
		maps = append(maps, m)
		Append(k, v, segFiles, maps, fileIndex)
	}
	maps[*fileIndex][k] = loc // byte offset value
	n, err := f.WriteString(fmt.Sprintf("%s,%v\n", k, v))
	if err != nil {
		fmt.Printf("failed to write key value pair to file: %s\n", err)
		return
	}
	fmt.Printf("%v bytes written at %v bytes of segment file %v\n", n, loc, *fileIndex)
}

// LookUp returns the associated value of a given key
func LookUp(k string, segFiles []*os.File, maps []map[string]int) string {
	var loc, index int
	for i := 0; i < len(maps); i++ {
		if v, ok := maps[i][k]; ok {
			loc = v
			index = i
		}
	}
	f := segFiles[index]
	line := make([]byte, elemSize)

	n, err := f.ReadAt(line, int64(loc))
	if err != nil {
		fmt.Printf("Failed to lookup value in segment file: %s\n", err)
	}
	fmt.Printf("Read %v bytes from segment file\n", n)
	return cleanLine(line)
}

// compact removes duplicate entries in a data file
// func compact()

// merge aggregates two data files together
// func merge()

func cleanLine(b []byte) string {
	s := string(b)
	val, _, ok := strings.Cut(s, "\n")
	if ok {
		return val
	}
	fmt.Println("value already cleaned")
	return ""
}
