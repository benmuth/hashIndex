package hashIndex

import (
	"fmt"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	m, dataFiles := Init()
	fmt.Println("test file ", dataFiles)
	segFileIndex := new(int)
	pairs := []struct {
		key string
		val int
	}{
		{
			key: "apple",
			val: 3,
		},
		{
			key: "banana",
			val: 6,
		},
		{
			key: "mango",
			val: 8,
		},
		{
			key: "grape",
			val: 29,
		},
	}
	for i := 0; i < len(pairs); i++ {
		Append(pairs[i].key, pairs[i].val, dataFiles, m, segFileIndex)
	}
	//f := dataFiles[0]
	fs, err := dataFiles[0].Stat()
	if err != nil {
		fmt.Printf("failed to get file info: %s\n", err)
	}

	data, err := os.ReadFile("../files/" + fs.Name())
	//n, err := io.ReadFull(f, data)
	if err != nil {
		fmt.Printf("failed to read test file: %s\n", err)
	}
	fmt.Println(data)

	// cleanup
	if err = os.Remove("../files/file1.csv"); err != nil {
		fmt.Printf("Failed to remove file: %s, %s\n", "file1.csv", err)
	}
}
