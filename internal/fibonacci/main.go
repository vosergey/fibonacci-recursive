package fibonacci

func Recursive(n int64, m map[int64]int64) int64 {
	if n <= 1 {
		return n
	}

	a := n - 1
	b := n - 2

	if _, exists := m[a]; !exists {
		m[a] = Recursive(a, m)
	}

	if _, exists := m[b]; !exists {
		m[b] = Recursive(b, m)
	}

	return m[a] + m[b]
}

func Iterative(number int64) int64 {

	var a int64 = 1
	var b int64 = 1
	var i int64 = 3

	for i <= number {
		c := a + b

		a = b
		b = c

		i++
	}

	return b
}
