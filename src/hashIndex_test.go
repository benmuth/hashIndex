package hashIndex

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// func TestAppend(t *testing.T) {
// 	maps, dataFiles := Init()
// 	fmt.Println("test file ", dataFiles)
// 	fileIndex := new(int)
// 	pairs := []struct {
// 		key string
// 		val string
// 	}{
// 		{
// 			key: "apple",
// 			val: "red",
// 		},
// 		{
// 			key: "banana",
// 			val: "yellow",
// 		},
// 		{
// 			key: "mango",
// 			val: "orange",
// 		},
// 		{
// 			key: "grape",
// 			val: "purple",
// 		},
// 	}
// 	for _, p := range pairs {
// 		t.Run(p.key, func(t *testing.T) {
// 			Append(p.key, p.val, dataFiles, maps, fileIndex)
// 		})
// 	}
// 	//f := dataFiles[0]
// 	fs, err := dataFiles[0].Stat()
// 	if err != nil {
// 		fmt.Printf("failed to get file info: %s\n", err)
// 	}

// 	data, err := os.ReadFile("../files/" + fs.Name())
// 	//n, err := io.ReadFull(f, data)
// 	if err != nil {
// 		fmt.Printf("failed to read test file: %s\n", err)
// 	}
// 	fmt.Println(string(data))

// 	// cleanup
// 	t.Cleanup(func() {
// 		if err = os.Remove("../files/file1.csv"); err != nil {
// 			fmt.Printf("Failed to remove file: %s, %s\n", "file1.csv", err)
// 		}
// 	})
// }

// func TestLookUp(t *testing.T) {
// 	//setup
// 	maps, dataFiles := Init()
// 	fmt.Println("test file ", dataFiles)
// 	fileIndex := new(int)
// 	pairs := []struct {
// 		key  string
// 		val  string
// 		want string
// 	}{
// 		{
// 			key:  "apple",
// 			val:  "red",
// 			want: "red",
// 		},
// 		{
// 			key:  "banana",
// 			val:  "yellow",
// 			want: "yellow",
// 		},
// 		{
// 			key:  "mango",
// 			val:  "orange",
// 			want: "",
// 		},
// 		{
// 			key:  "grape",
// 			val:  "purple",
// 			want: "",
// 		},
// 	}
// 	Append(pairs[0].key, pairs[0].val, dataFiles, maps, fileIndex)
// 	Append(pairs[1].key, pairs[1].val, dataFiles, maps, fileIndex)

// 	//run tests
// 	for _, p := range pairs {
// 		t.Run(p.key, func(t *testing.T) {
// 			got := LookUp(p.key, dataFiles, maps)
// 			if got != p.want {
// 				t.Fatalf("Failed to lookup value: got %s, want %s\n", got, p.want)
// 			}
// 		})
// 	}

// 	fs, err := dataFiles[0].Stat()
// 	if err != nil {
// 		fmt.Printf("failed to get file info: %s\n", err)
// 	}

// 	data, err := os.ReadFile("../files/" + fs.Name())
// 	if err != nil {
// 		fmt.Printf("failed to read test file: %s\n", err)
// 	}
// 	fmt.Println(string(data))

// 	// teardown
// 	t.Cleanup(func() {
// 		if err = os.Remove("../files/file1.csv"); err != nil {
// 			fmt.Printf("Failed to remove file: %s, %s\n", fs, err)
// 		}
// 	})
// }

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []byte
	}{
		{
			name: "hello",
			text: "hello",
			want: []byte{5, 104, 101, 108, 108, 111},
		},
		{
			name: "more than 256 characters",
			text: "uniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYork",
			want: []byte{255, 2, 1, 17, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107},
		},
	}

	for _, test := range tests {
		got := encode(test.text)
		if !cmp.Equal(got, test.want) {
			t.Errorf("%s: got != want, diff: %v\n", test.name, cmp.Diff(got, test.want))
		}
	}
}

type testByte []byte

func (t testByte) ReadAt(p []byte, off int64) (n int, err error) {
	n = len(p)

	end := int(off) + n
	if end > len(t) {
		n = end - len(t)
		fmt.Println(n, end, len(t))
		err = fmt.Errorf("didn't reach end of file")
	}
	// fmt.Println(t[off:])
	for i := 0; i < n; i++ {
		p[i] = t[i+int(off)]
	}
	return n, err
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "hello",
			data: []byte{5, 104, 101, 108, 108, 111},
			want: "hello",
		},
		{
			name: "more than 256 characters",
			data: []byte{255, 2, 1, 17, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107, 117, 110, 105, 113, 117, 101, 78, 101, 119, 89, 111, 114, 107},
			want: "uniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYork",
		},
	}

	for _, test := range tests {
		got, err := decode(0, testByte(test.data))
		if err != nil {
			// t.Errorf("got error: %v\n", err)
			fmt.Printf("got error: %v\n", err)
		}
		if !cmp.Equal(got, test.want) {
			t.Errorf("%s: got != want, diff: %v\n", test.name, cmp.Diff(got, test.want))
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "hello",
			text: "hello",
		},
		{
			name: "more than 256 characters",
			text: "uniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYorkuniqueNewYork",
		},
	}

	for _, test := range tests {
		encoded := encode(test.text)
		got, err := decode(0, testByte(encoded))
		if err != nil {
			// t.Errorf("got error: %v\n", err)
			fmt.Printf("got error: %v\n", err)
		}
		if !cmp.Equal(got, test.text) {
			t.Errorf("%s: got != want, diff: %v\n", test.name, cmp.Diff(got, test.text))
		}
	}
}