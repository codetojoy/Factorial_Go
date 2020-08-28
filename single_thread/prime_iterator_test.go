
package main

import "testing"

func TestNext(t *testing.T) {
    primeIndex := PrimeIndex{}
    primeIndex.Initialize(10)
    primeIterator := NewPrimeIterator(primeIndex)

    // test
    primeIterator.Next()
    primeIterator.Next()
    result := primeIterator.Next()

    expected := 5

    if result != expected {
        t.Errorf("Factors.Put() == %v, want %v", result, expected)
    }
}
