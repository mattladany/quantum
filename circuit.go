package quantum

import (
	"bytes"
	"fmt"
)

// Circuit a quantum circuit that contains qubits and quantum gates.
type Circuit struct {
	dimensions int
	qubits     []*Qubit
	gates      [][2][2]complex128
}

// MakeCircuit creates a Circuit with the provided qubits.
func MakeCircuit(dimensions int, qubits ...*Qubit) (*Circuit, error) {
	circuit := new(Circuit)
	if dimensions != len(qubits) {
		return nil, fmt.Errorf("circuit size: %d does not match number of qubits: %d", dimensions, len(qubits))
	}
	circuit.dimensions = dimensions

	for _, qubit := range qubits {
		circuit.qubits = append(circuit.qubits, qubit)
	}

	return circuit, nil
}

// MakeSingleQubitCircuit creates a Circuit with a single Qubit that has the specified gates.
func MakeSingleQubitCircuit(qubit *Qubit, gates ...[2][2]complex128) (*Circuit, error) {
	circuit := new(Circuit)
	circuit.dimensions = 1
	circuit.qubits = append(circuit.qubits, qubit)
	for _, gate := range gates {
		circuit.gates = append(circuit.gates, gate)
	}
	return circuit, nil
}

// String returns a string representation of the qubits, gates, and their interactions currently loaded in this circuit.
func (circuit Circuit) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(circuit.qubits[0].String())
	//buffer.WriteString(fmt.Sprintf("[%v|%0*s> + ", circuit.qubits[0].basis0, 1, strconv.FormatInt(int64(0), 2)))
	//buffer.WriteString(fmt.Sprintf("%v|%0*s>]---", circuit.qubits[0].basis1, 1, strconv.FormatInt(int64(1), 2)))
	// TODO write function to compare gate slices and return a string representation of the gate
	//    for _, _ = range circuit.gates {
	//        buffer.WriteString(fmt.Sprintf("%s---", "X"))
	//    }

	return buffer.String()[:len(buffer.String())-3]
}

// applyGate applies the provided gate to the provided qubit.
func applyGate(qubit *Qubit, gate [2][2]complex128) {
	oldBasis0 := qubit.basis0
	oldBasis1 := qubit.basis1

	qubit.basis0 = gate[0][0]*oldBasis0 + gate[0][1]*oldBasis1
	qubit.basis1 = gate[1][0]*oldBasis0 + gate[1][1]*oldBasis1
}
