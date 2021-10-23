package quantum

import (
	"math"
	"testing"
)

func TestMakeQubit0(t *testing.T) {
	qubit, err := MakeQubit(1, 0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if qubit.basis0 != 1 || qubit.basis1 != 0 {
		t.Fatalf("qubit did not initialize correctly")
	}
}

func TestMakeQubit1(t *testing.T) {
	qubit, err := MakeQubit(0, 1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if qubit.basis0 != 0 || qubit.basis1 != 1 {
		t.Fatalf("qubit did not initialize correctly")
	}

}

func TestMakeQubitSuperposition(t *testing.T) {
	_, err := MakeQubit(complex(1.0/math.Sqrt(2), 0), complex(1.0/math.Sqrt(2), 0))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeQubitInvalid(t *testing.T) {
	_, err := MakeQubit(0.5, 0.5)
	if err == nil {
		t.Fatalf("qubit should not have been created, but was")
	}
}

func TestMakeTensoredQubitsSimpleFloats(t *testing.T) {
	_, err := MakeTensoredQubits(0.5, 0.5, 0.5, 0.5)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeTensoredQubitsRationalFloats(t *testing.T) {
	_, err := MakeTensoredQubits(complex(1.0/math.Sqrt(2), 0), complex(1.0/math.Sqrt(2), 0))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeTensoredQubitsWithEntangledState(t *testing.T) {
	_, err := MakeTensoredQubits(complex(1.0/math.Sqrt(2), 0), 0, 0, complex(1.0/math.Sqrt(2), 0))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeTensoredQubitsWithComplexStates(t *testing.T) {
	_, err := MakeTensoredQubits(0, 0.5i, 0+0.5i, 0+0.5i, 0+0.5i)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeTensoredQubitsWithInvalidState(t *testing.T) {
	_, err := MakeTensoredQubits(0, 0.5i, 0+0.5i, 0.5)
	if err == nil {
		t.Fatalf("Successfully created qubit when it should have failed")
	}
}
