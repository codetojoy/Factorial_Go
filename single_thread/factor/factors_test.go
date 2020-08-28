
package factor

import (
    "testing"
    "../prime"
)

func TestPut(t *testing.T) {
    primeIndex := prime.New(10)
    factors := NewFactors(10, primeIndex)

    // test
    factors.Put(2,2);
    factors.Put(3,3);

    result := factors.String()
    expected := "2: 2, 3: 3, "

    if result != expected {
        t.Errorf("Factors.Put() == %v, want %v", result, expected)
    }
}

func TestEquals(t *testing.T) {
    cases := []struct {
        inA []int
        inB []int
        expected bool
    }{
        { []int{1,1}, []int{1,1}, true },
        { []int{1,1,0}, []int{1,1}, false },
        { []int{3,4,5}, []int{3,4,5}, true },
        { []int{3,3,3}, []int{3,4,5}, false },
    }

    max := 5

    for _, c := range cases {
        factorsA := NewFactorsForTesting(max, c.inA)
        factorsB := NewFactorsForTesting(max, c.inB)

        // test
        result := factorsA.Equals(factorsB)

        if result != c.expected {
            t.Errorf("[%v].Equals(%v) == %v, expected %v", c.inA, c.inB, result, c.expected)
        }
    }
}

func TestMultiply(t *testing.T) {
    cases := []struct {
        inA []int
        inB []int
        expected []int
    }{
        { []int{1,0,0,0,0}, []int{0,1,0,0,0}, []int{1,1,0,0,0} },
        { []int{1,1,0,0,0}, []int{1,1,0,0,1}, []int{2,2,0,0,1} },
    }

    max := 5

    for _, c := range cases {
        factorsA := NewFactorsForTesting(max, c.inA)
        factorsB := NewFactorsForTesting(max, c.inB)
        factorsExpected := NewFactorsForTesting(max, c.expected)

        // test
        result, e := factorsA.Multiply(factorsB)

        if e == nil && result.Equals(factorsExpected) {
            t.Errorf("[%v].Multiply(%v) == %v, expected %v", c.inA, c.inB, result, c.expected)
        }
    }
}
