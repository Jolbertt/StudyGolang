package main

import (
	"fmt"
	"bytes"
)

func comma(s string) string {
	n := 0
	var buf bytes.Buffer
	for i := 0; i < len(s); i++{
		if n >2{
			buf.WriteString(",")
			n=0
		}
		fmt.Fprintf(&buf, "%c", s[(len(s)-1)-i])
		n++
	}
	var buf2 bytes.Buffer
	for i := 0; i < len(buf.String()); i++ {
		fmt.Fprintf(&buf2, "%c", buf.String()[(len(buf.String())-1)-i])
	}
	return buf2.String() 
}

func main(){
	fmt.Println(comma("07786466000103"))
}
