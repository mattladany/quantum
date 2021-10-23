package quantum

import "math"

// MakeIGate create a single qubit identity matrix operator.
func MakeIGate() [2][2]complex128 {
	return [2][2]complex128{
		{1, 0},
		{0, 1},
	}
}

// MakeXGate create a single qubit Pauli-X matrix operator.
func MakeXGate() [2][2]complex128 {
	return [2][2]complex128{
		{0, 1},
		{1, 0},
	}
}

// MakeYGate create a single qubit Pauli-Y matrix operator.
func MakeYGate() [2][2]complex128 {
	return [2][2]complex128{
		{0, 0 + 1i},
		{0 - 1i, 0},
	}
}

// MakeZGate create a single qubit Pauli-Z matrix operator.
func MakeZGate() [2][2]complex128 {
	return [2][2]complex128{
		{1, 0},
		{0, -1},
	}
}

// MakeHadamardGate create a single qubit Hadamard matrix operator.
func MakeHadamardGate() [2][2]complex128 {
	return [2][2]complex128{
		{complex(1.0/math.Sqrt(2), 0), complex(1.0/math.Sqrt(2), 0)},
		{complex(1.0/math.Sqrt(2), 0), complex(-1.0/math.Sqrt(2), 0)},
	}
}
