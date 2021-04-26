package hashmap

import (
	"testing"
)

func TestNegativeKey(t *testing.T) {
	m := &Map{}
	bit := 10
	m.Init(bit)
	nums := []int{-2, -7, 11, 15}

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	for i, _ := range nums {
		if _, ok := m.Get(nums[i]); !ok {
			t.Errorf("nums[%d] doesn't exist!", i)
		}
	}
}

func TestNotFound(t *testing.T) {
	m := &Map{}
	bit := 10
	m.Init(bit)
	nums := []int{-2, -7, 11, 15}

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	if _, ok := m.Get(0); ok {
		t.Errorf("Found key 0, but it should not be found.")
	}
}
func ExampleString() {
	m := &Map{}
	bit := 10
	m.Init(bit)
	nums := []int{2, 7, 11, 15}

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	m.String()

	// Output:
	// Map[2]: 0
	// Map[7]: 1
	// Map[11]: 2
	// Map[15]: 3
}
