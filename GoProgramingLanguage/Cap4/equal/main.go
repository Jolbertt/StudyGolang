package main

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	// Example usage of the equal function
	x := []string{"apple", "banana", "cherry"}
	y := []string{"apple", "banana", "cherry"}
	z := []string{"apple", "banana"}

	println(equal(x, y)) // Should print true
	println(equal(x, z)) // Should print false
}