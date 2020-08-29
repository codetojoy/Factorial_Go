
package prime

import "testing"

func TestNext(t *testing.T) {
    primeIndex := New(10)
    primeIterator := NewPrimeIterator(primeIndex)

    // test
    primeIterator.Next()
    primeIterator.Next()
    result := primeIterator.Next()

    expected := 5

    if result != expected {
        t.Errorf("PrimeIterator.Next() == %v, want %v", result, expected)
    }
}

func TestReset(t *testing.T) {
    primeIndex := New(10)
    primeIterator := NewPrimeIterator(primeIndex)

    primeIterator.Next()
    primeIterator.Next()
    primeIterator.Next()

    // test
    primeIterator.Reset()
    result := primeIterator.Next()

    expected := 2

    if result != expected {
        t.Errorf("PrimeIterator.Reset() == %v, want %v", result, expected)
    }
}
