package quantum

import (
	"errors"
	"fmt"
)

// Circuit a quantum circuit that contains qubits and quantum gates.
type Circuit struct {
	size   int
	qubits []*Qubit
}

// MakeCircuit creates a Circuit with the provided qubits.
func MakeCircuit(size int, qubits ...*Qubit) (*Circuit, error) {
	circuit := new(Circuit)
	if size != len(qubits) {
		return nil, errors.New(fmt.Sprintf("circuit size: %d does not match number of qubits: %d", size, len(qubits)))
	}
	circuit.size = size

	for _, qubit := range qubits {
		circuit.qubits = append(circuit.qubits, qubit)
	}

	return circuit, nil
}
