package quantum

import (
    "math"
    "testing"
)

func TestMakeIGate(t *testing.T) {
    i := MakeIGate()
    if i[0][0] != 1 {
        t.Fatalf("Invalid I gate")
    }
    if i[0][1] != 0 {
        t.Fatalf("Invalid I gate")
    }
    if i[1][0] != 0 {
        t.Fatalf("Invalid I gate")
    }
    if i[1][1] != 1 {
        t.Fatalf("Invalid I gate")
    }
}

func TestMakeXGate(t *testing.T) {
    x := MakeXGate()
    if x[0][0] != 0 {
        t.Fatalf("Invalid X gate")
    }
    if x[0][1] != 1 {
        t.Fatalf("Invalid X gate")
    }
    if x[1][0] != 1 {
        t.Fatalf("Invalid X gate")
    }
    if x[1][1] != 0 {
        t.Fatalf("Invalid X gate")
    }
}

func TestMakeYGate(t *testing.T) {
    y := MakeYGate()
    if y[0][0] != 0 {
        t.Fatalf("Invalid Y gate")
    }
    if y[0][1] != 0+1i {
        t.Fatalf("Invalid Y gate")
    }
    if y[1][0] != 0-1i {
        t.Fatalf("Invalid Y gate")
    }
    if y[1][1] != 0 {
        t.Fatalf("Invalid Y gate")
    }
}

func TestMakeZGate(t *testing.T) {
    z := MakeZGate()
    if z[0][0] != 1 {
        t.Fatalf("Invalid Z gate")
    }
    if z[0][1] != 0 {
        t.Fatalf("Invalid Z gate")
    }
    if z[1][0] != 0 {
        t.Fatalf("Invalid Z gate")
    }
    if z[1][1] != -1 {
        t.Fatalf("Invalid Z gate")
    }
}

func TestMakeHadamardGate(t *testing.T) {
    hadamard := MakeHadamardGate()
    if real(hadamard[0][0])-float64EqualityThreshold > 1.0/math.Sqrt(2) ||
        real(hadamard[0][0])+float64EqualityThreshold < 1.0/math.Sqrt(2) {
        t.Fatalf("Invalid Hadamard gate")
    }
    if real(hadamard[0][1])-float64EqualityThreshold > 1.0/math.Sqrt(2) ||
        real(hadamard[0][1])+float64EqualityThreshold < 1.0/math.Sqrt(2) {
        t.Fatalf("Invalid Hadamard gate")
    }
    if real(hadamard[1][0])-float64EqualityThreshold > 1.0/math.Sqrt(2) ||
        real(hadamard[1][0])+float64EqualityThreshold < 1.0/math.Sqrt(2) {
        t.Fatalf("Invalid Hadamard gate")
    }
    if real(hadamard[1][1])-float64EqualityThreshold > -1.0/math.Sqrt(2) ||
        real(hadamard[1][1])+float64EqualityThreshold < -1.0/math.Sqrt(2) {
        t.Fatalf("Invalid Hadamard gate")
    }
}
