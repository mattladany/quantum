package quantum

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

// Qubit is a representation of a qubit.
type Qubit struct {
	basis0 complex128
	basis1 complex128
}

// MakeQubit creates a qubit from the provided basis states.
// It returns a Qubit struct if the provided states represent a unit vector.
// It returns an error if the provided states do not represent a unit vector.
func MakeQubit(basis0 complex128, basis1 complex128) (*Qubit, error) {

	// Make sure the provided basis states represent a unit vector
	if !checkIfUnitVector(basis0, basis1) {
		return nil, fmt.Errorf("a qubit must be a unit vector")
	}

	qubit := new(Qubit)
	qubit.basis0 = basis0
	qubit.basis1 = basis1

	return qubit, nil
}

// Basis0 retrieves the value of basis0 for this Qubit.
func (qubit *Qubit) Basis0() complex128 {
	return qubit.basis0
}

// Basis1 retrieves the value of basis1 for this Qubit.
func (qubit *Qubit) Basis1() complex128 {
	return qubit.basis1
}

// Update updates this Qubit's fields with the provided values.
// Will return nil if the provided values make a unit vector.
// Will return an error if the provided values do not make a unit vector.
func (qubit *Qubit) Update(basis0 complex128, basis1 complex128) error {

	// Make sure the provided basis states represent a unit vector
	if !checkIfUnitVector(basis0, basis1) {
		return fmt.Errorf("a qubit must be a unit vector")
	}

	qubit.basis0 = basis0
	qubit.basis1 = basis1
	return nil
}

// String writes the provided Qubit as a string in Dirac-esque notation.
func (qubit *Qubit) String() string {
	var stateBuffer bytes.Buffer
	stateBuffer.WriteString(fmt.Sprintf("%v|%0*s> + ", qubit.basis0, 1, strconv.FormatInt(int64(0), 2)))
	stateBuffer.WriteString(fmt.Sprintf("%v|%0*s> + ", qubit.basis1, 1, strconv.FormatInt(int64(1), 2)))
	return stateBuffer.String()
}

// checkIfUnitVector returns true if the sum of squared magnitutes of
//  the parameters is 1, false otherwise.
func checkIfUnitVector(states ...complex128) bool {
	var sum float64

	for _, state := range states {
		sum += math.Abs(real(state * state))
	}

	// Make sure the provided basis sates represent a unit vector
	if sum-float64EqualityThreshold > 1.0 || sum+float64EqualityThreshold < 1.0 {
		return false
	}
	return true
}

// TensoredQubits is a representation of a tensor product of qubits.
type TensoredQubits struct {
	states []complex128 // states are the states of the basis vectors for this qubit.
}

// MakeTensoredQubits creates a tensor product of qubits from the provided states.
// It returns a TensoredQubits struct if the provided states represent a unit vector.
// It returns an error if the provided states do not represent a unit vector.
func MakeTensoredQubits(states ...complex128) (*TensoredQubits, error) {
	tensoredQubits := new(TensoredQubits)

	// Make sure the states provided represent a unit vector
	if !checkIfUnitVector(states...) {
		return nil, fmt.Errorf("a tensor product of qubits must be a unit vector")
	}

	return tensoredQubits, nil
}

// String writes the provided TensoredQubits as a string in Dirac-esque notation.
func (tensoredQubits TensoredQubits) String() string {
	var stateBuffer bytes.Buffer
	binaryDigits := int(math.Sqrt(float64(len(tensoredQubits.states))))
	for i, state := range tensoredQubits.states {
		stateBuffer.WriteString(fmt.Sprintf("%v|%0*s> + ", state, binaryDigits, strconv.FormatInt(int64(i), 2)))
	}
	stateString := stateBuffer.String()
	return stateString[:len(stateString)-2]
}
