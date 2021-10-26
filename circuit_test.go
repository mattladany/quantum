package quantum

import (
    "testing"
    "fmt"
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
