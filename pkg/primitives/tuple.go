package primitives

type Tuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

type Negatable interface {
	Negate() Tuple
}
