package quantum

import (
    "math"
    "testing"
)

func TestMakeQubitSimpleFloats(t *testing.T) {
    _, err := MakeQubit(0.5, 0.5, 0.5, 0.5)
    if err != nil {
        t.Fatalf(err.Error())
    }
}

func TestMakeQubitRationalFloats(t *testing.T) {
    _, err := MakeQubit(complex(1.0 / math.Sqrt(2), 0), complex(1.0 / math.Sqrt(2), 0))
    if err != nil {
        t.Fatalf(err.Error())
    }
}

func TestMakeQubitWithEntangledState(t *testing.T) {
    _, err := MakeQubit(complex(1.0 / math.Sqrt(2), 0), 0, 0, complex(1.0 / math.Sqrt(2), 0))
    if err != nil {
        t.Fatalf(err.Error())
    }
}

func TestMakeQubitWithComplexStates(t *testing.T) {
    _, err := MakeQubit(0, 0.5i, 0 + 0.5i, 0 + 0.5i, 0 + 0.5i)
    if err != nil {
        t.Fatalf(err.Error())
    }
}

func TestMakeQubitWithInvalidState(t *testing.T) {
    _, err := MakeQubit(0, 0.5i, 0 + 0.5i, 0.5)
    if err == nil {
        t.Fatalf("Successfully created qubit when it should have failed")
    }
}