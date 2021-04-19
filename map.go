package hashmap

import (
	"fmt"
)

type Map struct {
	bit int
	ht  []htHead
}

type htHead struct {
	first *htNode // points to first element of that bucket
}

type htNode struct {
	key, value int
	next       *htNode
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Size returns the size of hash table.
func (m *Map) Size() int { return 1 << m.bit }

func (m *Map) Init(bit int) {
	m.bit = bit
	size := m.Size()
	m.ht = make([]htHead, size)
}

func (h *htHead) appendHtNode(key, value int) {
	indirect := &h.first
	for ; (*indirect) != nil; indirect = &(*indirect).next {
		if (*indirect).key == key {
			// key exist, no need to append
			return
		}
	}

	// at tail of list, append newNode
	newNode := &htNode{key: key, value: value, next: nil}
	(*indirect) = newNode
}

func (m *Map) Add(key, value int) {
	set := abs(key) % (1 << m.bit)

	m.ht[set].appendHtNode(key, value)
}

func (m *Map) Get(key int) (value int, ok bool) {
	set := abs(key) % (1 << m.bit)
	// access m.ht[set]
	// if m.ht[set] not exist, means key not exist in map
	if len(m.ht) < set {
		return 0, false
	}

	for node := m.ht[set].first; node != nil; node = node.next {
		if node.key == key {
			// key exist, return the value and ok
			return node.value, true
		}
	}
	// key not found
	return 0, false
}

// String shows the whole map with format Map[key:value]
func (m *Map) String() {
	// figure out some method to contain all keys
	for i := 0; i < m.Size(); i++ {
		for node := m.ht[i].first; node != nil; node = node.next {
			fmt.Printf("Map[%d]: %d\n", node.key, node.value)
		}
	}
}
