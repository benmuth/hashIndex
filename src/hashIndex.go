package hashIndex

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// fileSize is the maximum size of a segmentFile
const maxFileSize = 1024

// elemSize is the maximum allowed size of a key to be stored
const elemSize = 64

type HashIndex struct {
	m        map[string]int
	segFiles []*os.File
	offset   int
}

// Init prepares a hash index for use, and returns the map and data file to be used
func Init() *HashIndex {
	f, err := os.Create("../files/f1")
	if err != nil {
		log.Fatalf("couldn't initialize data file: %s\n", err)
	}
	segFiles := []*os.File{f} //TODO: change from slice of files to just opening files by name/number
	m := make(map[string]int)
	return &HashIndex{m: m, segFiles: segFiles, offset: 0}
}

// current data format:
// lenByte|{length}|value(len=lenByte|length)
// if lenByte > 255 (max value of 1 byte), then look at length to determine how long the value is
// length would have to be greater than one byte, so you need a lenByte for length
// so, given an offset in a file, the first byte at that offset would indicate the length of the 'value'
// byte string, unless the first byte == 255, in which case you'd look at the next byte. The second byte
// in this case would indicate how many bytes (x) it takes to encode the length of the value. Then the third
// byte and on would be the length of the value. Then the value string would come after x bytes

// base case: len(val) < 255 => lenByte < 255 && data = []byte(lenByte|value)
// else: len(val) >= 255 => lenByte = 255 && data = []byte(lenByte=255|lenBytes([]byte)|value)

func Write(key string, val string, f *os.File, m map[string]int) error {
	if len(val) == 0 {
		return fmt.Errorf("failed to write to database: no value string provided.")
	}
	fs, err := f.Stat()
	if err != nil {
		return err
	}
	loc := fs.Size()

	data := encode(val)

	// append value to file at byte offset loc
	n, err := f.WriteAt(data, loc)
	if err != nil {
		return err
	}
	fmt.Println(n)
	// add key and byte offset to map
	m[key] = int(loc)
	return nil
}

func Read(key string, f *os.File, m map[string]int) (string, error) {
	// look up key in map
	offset, ok := m[key]
	if !ok {
		return "", fmt.Errorf("Failed to read key from file: key is not in map\n")
	}
	s, err := decode(offset, f)
	if err != nil {
		return "", err
	}
	return s, nil
}

// A value string is encoded with a length prefix byte, like this:
// 		"hello" -> []byte{5, 'h', 'e', 'l', 'l', 'o'}
// If the value string is longer than 255 bytes, multiple bytes are needed to encode the string, like this:
// 		"aLongStriii...ing" 	 ->	[]byte{255, 2, 1, 13, 'a', 'L', 'o'...}
// 					^255 'i's (len=269)    ^    ^  ^--^- string length bytes (1 * 2^8 + 13)
//										   |	bytes needed to encode string length (2 bytes for 269 chars)
//										   255 indicates long string

// encode encodes the value string into a length-prefixed slice of bytes.
func encode(val string) []byte {
	l := len(val) // maybe should be converted to bytes before length is measured?
	buf := make([]byte, 1)
	if l <= 255 { // the length of the value string can be encoded in 1 byte
		buf[0] = byte(l)
		for i := 1; i < len(buf); i++ {
			buf[i] = byte(val[i])
		}
		buf = append(buf, []byte(val)...)
	} else { // the length of the value string needs multiple bytes to be encoded
		buf[0] = 255
		// find the number of bytes needed to encode the length of the value string
		n := 0
		for x := l; x != 0; {
			x >>= 8
			n++
		}
		buf = append(buf, byte(n))
		// b is the number indicating the length of the value string
		b := make([]byte, n)
		for i := n - 1; i >= 0; i-- {
			b[i] = byte(l)
			l >>= 8
		}
		buf = append(buf, b...)
		buf = append(buf, val...)
	}
	return buf
}

// decode decodes a slice of bytes from a source at a given offset, and returns the string.
func decode(offset int, source io.ReaderAt) (string, error) {
	// find length of value string
	firstByte := make([]byte, 1)
	if _, err := source.ReadAt(firstByte, int64(offset)); err != nil {
		return "", err
	}
	length := firstByte[0]
	if length < 255 {
		s := make([]byte, length)
		if _, err := source.ReadAt(s, int64(offset+1)); err != nil {
			return "", err
		}
		return string(s), nil
	} else {
		secondByte := make([]byte, 1)
		if _, err := source.ReadAt(secondByte, int64(offset+1)); err != nil {
			return "", err
		}
		numLengthBytes := secondByte[0]
		valLength := make([]byte, numLengthBytes)
		if _, err := source.ReadAt(valLength, int64(offset+2)); err != nil {
			return "", err
		}
		l := 0
		for i := 0; i < int(numLengthBytes); i++ {
			l <<= 8
			l += int(valLength[i])
		}
		s := make([]byte, l)
		if _, err := source.ReadAt(s, int64(offset+2)+int64(numLengthBytes)); err != nil {
			return "", err
		}
		return string(s), nil
	}
}

// Append adds a key-value pair to the end of a segment file
// func Append(k string, v string, segFiles []*os.File, maps []map[string]int, fileIndex *int) {
// 	f := segFiles[*fileIndex]
// 	if len(k) > elemSize {
// 		log.Printf("Failed to add key value pair to hash index: character limit exceeded (lim: %v)\n", elemSize)
// 		return
// 	}
// 	fs, err := f.Stat()
// 	if err != nil {
// 		log.Printf("failed to get location of appended info: %s\n", err)
// 		return
// 	}
// 	loc := int(fs.Size())  //loc keeps track of the end of the file
// 	if loc > maxFileSize { //create a new segFile and a new map that points to it
// 		if f.Close() != nil {
// 			log.Printf("Failed to close segment file %s: %s", fs.Name(), err)
// 		}
// 		*fileIndex++
// 		fNext, err := os.Create(fmt.Sprintf("../files/f1", len(segFiles)))
// 		if err != nil {
// 			log.Printf("Failed to create new segment file: %s", err)
// 			return
// 		}
// 		segFiles = append(segFiles, fNext)
// 		m := make(map[string]int)
// 		maps = append(maps, m)
// 		Append(k, v, segFiles, maps, fileIndex)
// 	}
// 	maps[*fileIndex][k] = loc // byte offset value
// 	n, err := f.WriteString(fmt.Sprintf("%v%s%v%s", len(k), k, len(v), v))
// 	if err != nil {
// 		log.Printf("failed to write key value pair to file: %s\n", err)
// 		return
// 	}
// 	fmt.Printf("%v bytes written at %v bytes of segment file %v\n", n, loc, *fileIndex)
// }

// LookUp returns the associated value of a given key
// func LookUp(k string, segFiles []*os.File, maps []map[string]int) string {
// 	var loc, index int
// 	for i, m := range maps {
// 		if v, ok := m[k]; ok {
// 			loc = v
// 			index = i
// 			log.Printf("found value for key %s at location %v in file/map %v", k, loc, index)
// 			break
// 		} else {
// 			fmt.Printf("key %s not found in map", k)
// 			return ""
// 		}
// 	}
// 	f := segFiles[index]
// 	//6apple3red
// 	line := make([]byte)
// 	n, err := f.ReadAt(line, int64(loc))
// 	if err != nil {
// 		log.Printf("Failed to lookup value for key %s in segment file at location %v: %s\n", k, loc, err)
// 		return ""
// 	}
// 	fmt.Printf("Read %v bytes from segment file\n", n)
// 	return cleanLine(line)
// }

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
