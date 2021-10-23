package quantum

import (
    "errors"
    "fmt"
    "math"
)

// Qubit is a representation of a quantum qubit.
type Qubit struct {
    states []complex128 // states are the states of the basis vectors for this qubit.
}

// MakeQubit creates a Qubit from the provided states.
// It returns a Qubit struct if the provided states represent a unit vector.
// It returns an error if the provided states do not represent a unit vector.
func MakeQubit(states ...complex128) (Qubit, error) {
    var qubit Qubit
    var sum float64

    for _, state := range states {
        qubit.states = append(qubit.states, state)
        sum += real(state * state)
    }

    // TODO Figure out why this is required for complex numbers...
    sum = math.Abs(sum)

    // Make sure the states provided represent a unit vector
    if sum - float64EqualityThreshold > 1.0 || sum + float64EqualityThreshold < 1.0 {
        return qubit, errors.New(fmt.Sprintf("qubit must be a unit vector...sum was: %v", sum))
    }

    return qubit, nil
}