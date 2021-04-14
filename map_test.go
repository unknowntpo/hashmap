package hashmap

import (
	"fmt"
)

func Example() {
	m := &Map{}
	bit := 10
	m.Init(bit)
	nums := []int{2, 7, 11, 15}

	// add element of nums as key, and the index of element in nums as value
	for i, v := range nums {
		m.Add(v, i)
	}

	for i, _ := range nums {
		v, ok := m.Get(nums[i]) // Output: 0, true
		if ok {
			fmt.Printf("nums[%d]: %d\n", i, v)
		} else {
			fmt.Printf("nums[%d] doesn't exist!", i)
		}
	}

	// Output:
	// nums[0]: 0
	// nums[1]: 1
	// nums[2]: 2
	// nums[3]: 3
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
