package main

type Circle struct {
	x, y, r float64
}

func main() {
	var c0 Circle // Set c0 to 0. Where 0 means all fields with default value

	// Other forms to initialize:
	c1 := Circle{x: 157, y: 157, r: 122}
	c2 := Circle{157, 157, 122}
	c3 := &Circle{157, 157, 122}
}
