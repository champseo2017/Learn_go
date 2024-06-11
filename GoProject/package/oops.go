package oops

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (sq Square) Area() float64 {
	return sq.Side * sq.Side
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
