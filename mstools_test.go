//tests for mstools
package mstools

import "testing"

func TestRound(t *testing.T) {
    rounded:= Round(1.230001121135135,2)
    if rounded != 1.23 {
        t.Errorf("Rounded value is wrong. Wanted: %f, got: %f.", 1.23, rounded)
    }
}
