package quantum

import (
    "fmt"
    "testing"
)

func TestCreateQubit(t *testing.T) {
    var a Qubit = Qubit{
        states: []complex128{0.5, 0.5, 0.5, 0.5},
    }
    fmt.Println(a)
}
