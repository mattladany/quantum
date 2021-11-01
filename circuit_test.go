package quantum

import (
	"fmt"
	"math"
	"testing"
)

func TestMakeCircuit1Qubit1Size(t *testing.T) {
	qubit, err := MakeQubit(1, 0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = MakeCircuit(1, qubit)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMakeCircuitInvalid1Qubit2Size(t *testing.T) {
	qubit, err := MakeQubit(1, 0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = MakeCircuit(2, qubit)
	if err == nil {
		t.Fatalf("circuit should not have been created, but was")
	}
}

func TestMakeSingleQubitCircuitWith1Gate(t *testing.T) {
	qubit, err1 := MakeQubit(1, 0)
	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	x := MakeXGate()
	circuit, err2 := MakeSingleQubitCircuit(qubit, x)
	if err2 != nil {
		t.Fatalf(err2.Error())
	}

	fmt.Println(circuit)
}

func TestApplyGateX(t *testing.T) {
	qubit, err1 := MakeQubit(1, 0)
	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	x := MakeXGate()

	applyGate(qubit, x)

	if qubit.Basis0() != 0 || qubit.Basis1() != 1 {
		t.Fatalf("X gate was not applied correctly")
	}
}

func TestApplyGateY(t *testing.T) {
	qubit, err1 := MakeQubit(1, 0)
	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	y := MakeYGate()

	applyGate(qubit, y)

	if qubit.Basis0() != 0 || qubit.Basis1() != complex(0, -1) {
		t.Fatalf("Y gate was not applied correctly")
	}
}

func TestApplyGateZ(t *testing.T) {
	qubit, err1 := MakeQubit(1, 0)
	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	z := MakeZGate()

	applyGate(qubit, z)

	if qubit.Basis0() != 1 || qubit.Basis1() != 0 {
		t.Fatalf("Z gate was not applied correctly")
	}
}

func TestApplyGateHadamard(t *testing.T) {
	qubit, err1 := MakeQubit(1, 0)
	if err1 != nil {
		t.Fatalf(err1.Error())
	}

	h := MakeHadamardGate()

	applyGate(qubit, h)

	fmt.Println(qubit)

	if (real(qubit.Basis0())-1.0/math.Sqrt(2)) < (0-float64EqualityThreshold) || (real(qubit.Basis0())-1.0/math.Sqrt(2)) > (0+float64EqualityThreshold) || (real(qubit.Basis1())-1.0/math.Sqrt(2)) < (0-float64EqualityThreshold) || (real(qubit.Basis1())-1.0/math.Sqrt(2)) > (0+float64EqualityThreshold) {
		t.Fatalf("H gate was not applied correctly")
	}
}
