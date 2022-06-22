package vec

// Some 2D vector utilities

type Vec [2]float64

func Add(a, b Vec) Vec {
	return Vec{a[0] + b[0], a[1] + b[1]}
}

func Sub(a, b Vec) Vec {
	return Vec{a[0] - b[0], a[1] - b[1]}
}

func Mul(a Vec, b float64) Vec {
	return Vec{a[0] * b, a[1] * b}
}
