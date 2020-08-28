
package main

import "testing"

func TestIsPrime(t *testing.T) {
    cases := []struct {
        in int
        expected bool
    }{
        {2,true}, {3,true}, {4,false}, {5,true}, {6,false},
        {7,true}, {8,false}, {9,false}, {10,false},
    }
    primeIndex := PrimeIndex{}
    primeIndex.Initialize(11);

    for _, c := range cases {
        // test
        result := primeIndex.IsPrime(c.in)

        if result != c.expected {
            t.Errorf("GetPrimeForIndex(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}

func TestGetIndexForPrime(t *testing.T) {
    cases := []struct {
        in, expected int
    }{
        {2,0}, {3,1}, {5,2}, {7,3}, {11,4},
    }
    primeIndex := PrimeIndex{}
    primeIndex.Initialize(8);

    for _, c := range cases {
        // test
        result := primeIndex.GetIndexForPrime(c.in)

        if result != c.expected {
            t.Errorf("GetPrimeForIndex(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}

func TestGetPrimeForIndex(t *testing.T) {
    cases := []struct {
        in, expected int
    }{
        {0,2}, {1,3}, {2,5}, {3,7}, {4,11},
    }
    primeIndex := PrimeIndex{}
    primeIndex.Initialize(8);

    for _, c := range cases {
        // test
        result := primeIndex.GetPrimeForIndex(c.in)

        if result != c.expected {
            t.Errorf("GetPrimeForIndex(%v) == %v, expected %v", c.in, result, c.expected)
        }
    }
}
