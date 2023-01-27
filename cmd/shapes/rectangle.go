package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	r := Rectangle{Width: width, Height: height}
	return &r
}

func (r Rectangle) getArea() float64 {
	return r.Width * r.Height
}

func (r Rectangle) getPerimeter() float64 {
	return r.Width*2 + r.Height*2
}
