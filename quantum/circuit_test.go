package quantum

import (
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
