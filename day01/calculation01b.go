package day01

type Calc01b struct {
	naive Calc01
}

// So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just
// calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative.
func (m *Calc01b) RequiredFuel(mass int) int {
	x := m.naive.RequiredFuel(mass)
	if x <= 0 || mass <= 0 {
		return 0
	}

	return x + m.RequiredFuel(x)
}
