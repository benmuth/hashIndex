package hashIndex

import (
	"fmt"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	maps, dataFiles := Init()
	fmt.Println("test file ", dataFiles)
	fileIndex := new(int)
	pairs := []struct {
		key string
		val string
	}{
		{
			key: "apple",
			val: "red",
		},
		{
			key: "banana",
			val: "yellow",
		},
		{
			key: "mango",
			val: "orange",
		},
		{
			key: "grape",
			val: "purple",
		},
	}
	for _, p := range pairs {
		t.Run(p.key, func(t *testing.T) {
			Append(p.key, p.val, dataFiles, maps, fileIndex)
		})
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
	fmt.Println(string(data))

	// cleanup
	t.Cleanup(func() {
		if err = os.Remove("../files/file1.csv"); err != nil {
			fmt.Printf("Failed to remove file: %s, %s\n", "file1.csv", err)
		}
	})
}

func TestLookUp(t *testing.T) {
	//setup
	maps, dataFiles := Init()
	fmt.Println("test file ", dataFiles)
	fileIndex := new(int)
	pairs := []struct {
		key  string
		val  string
		want string
	}{
		{
			key:  "apple",
			val:  "red",
			want: "red",
		},
		{
			key:  "banana",
			val:  "yellow",
			want: "yellow",
		},
		{
			key:  "mango",
			val:  "orange",
			want: "",
		},
		{
			key:  "grape",
			val:  "purple",
			want: "",
		},
	}
	Append(pairs[0].key, pairs[0].val, dataFiles, maps, fileIndex)
	Append(pairs[1].key, pairs[1].val, dataFiles, maps, fileIndex)

	//run tests
	for _, p := range pairs {
		t.Run(p.key, func(t *testing.T) {
			got := LookUp(p.key, dataFiles, maps)
			if got != p.want {
				t.Fatalf("Failed to lookup value: got %s, want %s\n", got, p.want)
			}
		})
	}

	fs, err := dataFiles[0].Stat()
	if err != nil {
		fmt.Printf("failed to get file info: %s\n", err)
	}

	data, err := os.ReadFile("../files/" + fs.Name())
	if err != nil {
		fmt.Printf("failed to read test file: %s\n", err)
	}
	fmt.Println(string(data))

	// teardown
	t.Cleanup(func() {
		if err = os.Remove("../files/file1.csv"); err != nil {
			fmt.Printf("Failed to remove file: %s, %s\n", fs, err)
		}
	})
}
