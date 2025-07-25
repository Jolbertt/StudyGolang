package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // Create a slice with zero length but same capacity
	for _, s := range strings {
		if s != "" {
			out = append(out, s) // Append non-empty strings
		}
	}
	return out
}
func main() {
	data := []string{"Hello", "", "World", "", "!"}
	fmt.Printf("%q\n", nonempty2(data))
	fmt.Printf("%q\n", data) // Check if the original slice is modified
}