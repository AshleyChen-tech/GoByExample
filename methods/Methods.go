// Go supports methods defined on struct types.


package main

import "fmt"

type rect struct {
	width, height int 
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("ared:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("ared:", rp.area())
	fmt.Println("perim:", rp.perim())
}