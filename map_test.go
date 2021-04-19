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
		_, ok := m.Get(nums[i])
		if !ok {
			t.Errorf("nums[%d] doesn't exist!", i)
		}
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
