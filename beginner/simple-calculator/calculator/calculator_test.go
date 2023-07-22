package calculator

import "testing"

func assert(t testing.TB, got any, want any) {
    t.Helper()

    if got != want {
        t.Errorf("Got %v, Want %v", got, want)
    }
}

func TestCalculator(t *testing.T) {
    operands := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
    c := Calculator{operands: operands}

    t.Run("should equal 15 on `Add`", func(t *testing.T) {
        got := c.Add()
        want := 15.0
        assert(t, got, want)
    })

    t.Run("should equal -13 on `Sub`", func(t *testing.T) {
        got := c.Sub()
        want := -13.0
        assert(t, got, want)
    })

    t.Run("should equal 120 on `Mul`", func(t *testing.T) {
        got := c.Mul()
        want := 120.0
        assert(t, got, want)
    })

    t.Run("should equal  on `Sub`", func(t *testing.T) {
        got := c.Div()
        want := 1.0 / 2.0 / 3.0 / 4.0 / 5.0
        assert(t, got, want)
    })
}
