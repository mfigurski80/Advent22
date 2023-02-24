package utils

func DoForPermutations(a uint, fn func(uint, []uint)) {
	b := make([]uint, a)
	for i := range b {
		b[i] = uint(i)
	}
	fn(0, b)
	for {
		i := a - 1 // find largest index that's in ascending order
		for ; i > 0 && b[i-1] >= b[i]; i-- {
		}
		if i == 0 {
			return
		}
		j := a - 1 // find largest index that's greater than the pivot
		for ; b[j] <= b[i-1]; j-- {
		}
		b[i-1], b[j] = b[j], b[i-1]          // swap pivot with the greater element
		for j := uint(0); j < (a-i)/2; j++ { // reverse the suffix
			b[i+j], b[a-1-j] = b[a-1-j], b[i+j]
		}
		fn(i, b)
	}
}

func DoForSkipPermutations(a uint, fn func([]uint) uint) {
	b := make([]uint, a)
	for i := range b {
		b[i] = uint(i)
	}
	significantPositions := fn(b)
	// Initialize c to be an array of integers from 0 to a-1
	c := make([]int, a)
	for i := range c {
		c[i] = i
	}
	i := 0
	for {
		if i == int(a-1) {
			significantPositions = fn(b)
			i = 0
		} else if i >= 0 {
			c[i], c[int(a)-1-i] = c[int(a)-1-i], c[i]
			i = 0
		}
		// Generate the next permutation
		for i = i + 1; i < int(a); {
			if c[i] < c[i-1] {
				break
			}
			i++
		}
		if i == int(a) {
			return
		}
		j := 0
		if (i % 2) == 0 {
			j = 0
		} else {
			j = c[i-1]
		}
		c[i-1], c[j] = c[j], c[i-1]
		// Copy the permutation from c to b
		for k := range c {
			b[k] = uint(c[k])
		}
		// Skip permutations that are known not to change any of the significant positions
		skip := true
		for k := uint(0); k < significantPositions; k++ {
			if b[k] != b[k+1] {
				skip = false
				break
			}
		}
		if skip {
			i++
		}
	}
}

func DoForElementPermutations[T any](a []T, fn func(int, []T)) {
	b := make([]T, len(a))
	copy(b, a)
	order := make([]uint, len(a))
	for i := range a {
		order[i] = uint(i)
	}

	fn(0, b)
	for {
		i := len(b) - 1
		for ; i > 0 && order[i-1] >= order[i]; i-- {
		}
		if i <= 0 {
			return
		}
		j := len(order) - 1
		for ; order[j] <= order[i-1]; j-- {
		}
		order[i-1], order[j] = order[j], order[i-1]
		b[i-1], b[j] = b[j], b[i-1]
		for j := len(order) - 1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
			order[i], order[j] = order[j], order[i]
		}
		fn(0, b)
	}
}
