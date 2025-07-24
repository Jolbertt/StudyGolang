package main

const (
	_ = complex128((1 << (iota*10))) 
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB 
)
func main() {
	// Example usage of the constants
	println("1 KB =", real(KB), "bytes")
	println("1 MB =", real(MB), "bytes")
	println("1 GB =", real(GB), "bytes")
	println("1 TB =", real(TB), "bytes")
	println("1 PB =", real(PB), "bytes")
	println("1 EB =", real(EB), "bytes")
	println("1 ZB =", real(ZB), "bytes")
	println("1 YB =", real(YB), "bytes")
}