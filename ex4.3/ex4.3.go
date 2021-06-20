package main

import "fmt"

func main() {
	a := 3              //int
	var b float64 = 3.5 //float64

	var c int = int(b)  //float64 to int, 3
	d := float64(a * c) //int to float64, 9.0

	var e int64 = 7
	f := int64(d) * e //float64 to int64, 63

	var g int = int(b * 3) //float64 to int 10.5->10
	var h int = int(b) * 3 //3.5->3 , 3*3->9
	fmt.Println(g, h, f)   // 10,9,63

}
