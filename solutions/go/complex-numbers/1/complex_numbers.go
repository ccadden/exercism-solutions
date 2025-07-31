package complexnumbers

import "math"

type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		a: (n1.a + n2.a),
		b: (n1.b + n2.b),
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		a: (n1.a - n2.a),
		b: (n1.b - n2.b),
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	// (a + i * b) * (c + i * d) = (a * c - b * d) + (b * c + a * d) * i
	return Number{
		a: (n1.a * n2.a) - (n1.b * n2.b),
		b: (n1.b * n2.a) + (n1.a * n2.b),
	}
}

func (n Number) Times(factor float64) Number {
	return n.Multiply(Number{a: factor, b: 0})
}

func (n1 Number) Divide(n2 Number) Number {
	// (a + i * b) / (c + i * d) = (a * c + b * d)/(c^2 + d^2) + (b * c - a * d)/(c^2 + d^2) * i
	return Number{
		a: ((n1.a * n2.a) + (n1.b * n2.b))/(n2.a * n2.a + n2.b * n2.b),
		b: ((n1.b * n2.a) - (n1.a * n2.b))/(n2.a * n2.a + n2.b * n2.b),
	}
}

func (n Number) Conjugate() Number {
	return Number{
		a: n.a,
		b: -1 * n.b,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a * n.a + n.b * n.b)
}

func (n Number) Exp() Number {
	e := Number{a: math.Exp(n.a)}

	return e.Multiply(Number{a: math.Cos(n.b), b: math.Sin(n.b)})
}
