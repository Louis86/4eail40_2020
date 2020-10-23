package main

import ( "fmt"
					"math"
				)

// Implement types Rectangle, Circle and Shape
type Shape interface {
   Area() float64
	 fmt.Stringer
}

type Rectangle struct{
   width ,length  float64
}

type Circle struct{
   radius float64
}

func (r Rectangle) Area() float64{
   return r.length * r.width
}
func (c Circle) Area()float64{
	return 3.14*math.Pow(c.radius,2)
}

func (r Rectangle) String() string{
   return fmt.Sprintf("Square of width %.f and length %.f", r.width, r.length)
}

func (c Circle) String() string{
   return fmt.Sprintf("Circle of radius %.f", c.radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
