package quantum

import (
	"testing"
)

func TestMakeCircuit(t *testing.T) {
	qubit2d, err := MakeQubit(0.5, 0.5, 0.5, 0.5)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = MakeCircuit(qubit2d)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
