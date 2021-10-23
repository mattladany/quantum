package quantum

// Circuit a quantum circuit that contains qubits and quantum gates.
type Circuit struct {
	qubits []Qubit
}

// MakeCircuit creates a Circuit with the provided qubits.
func MakeCircuit(qubits ...Qubit) (*Circuit, error) {
	circuit := new(Circuit)

	for _, qubit := range qubits {
		circuit.qubits = append(circuit.qubits, qubit)
	}

	return circuit, nil
}