package hashMap

const mapSize = 128

type hashMap [mapSize]bucket

type bucket *bucketNode

type bucketNode struct {
	s    string
	val  int
	next *bucketNode
}

// init initializes a hashMap of size mapSize
func (h *hashMap) init() {
	for i := range *h {
		var b bucket
		h[i] = b
	}
}

// hash hashes a string into an integer
func hash(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c)
	}
	return sum
}

// Search looks up a string and if it exists, returns its corresponding value and true, or if it doesn't exist, returns 0 and false
func (h *hashMap) Search(s string) (int, bool) {
	loc := hash(s) % mapSize
	for node := h[loc]; node != nil; node = node.next {
		if node.s == s {
			return node.val, true
		}
	}
	return 0, false
}

// Add adds an entry to the map or updates the existing entry
func (h *hashMap) Add(s string, val int) {
	loc := hash(s) % mapSize
	if h[loc] == nil {
		h[loc] = &bucketNode{s: s, val: val}
	} else {
		for node := h[loc]; node.next != nil; node = node.next {
			if node.next == nil {
				node.next = &bucketNode{s: s, val: val}
			}
		}
	}
}
