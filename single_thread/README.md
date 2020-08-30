
### Summary

* The `run.sh` and `test.sh` scripts are not idiomatic.
* I have tried to set up `GOPATH` etc but I just can't get it to work.
    - e.g. to use a 3rd-party package such as: https://github.com/stretchr/testify
* Since this is functional, I have to move on, as I'm more interested in
  the language than the insufferable module/package system.

### TODO

* break loop into chunks for prep re: multiple threads 
* refine import so that we don't need `prime.PrimeIndex`
* rename `New()` function for `PrimeIndex` ?
* rename `PrimeIndex` to `Primes`

x * use reference to PrimeIndex
x * `Config` struct to consolidate params
x * test edge-cases
x * use `GOPATH` effectively ?
x * code coverage in tests?
