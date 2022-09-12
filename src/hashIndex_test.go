package hashIndex

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	m, dataFiles := Init()
	var segFileIndex *int
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
	fmt.Println(m)
}
