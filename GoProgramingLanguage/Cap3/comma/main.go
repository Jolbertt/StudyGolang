package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3{
		return s
	}
	return comma(s[:n-3])+","+s[n-3:]
}

func main(){
	fmt.Println(comma("07786466000103"))
}
