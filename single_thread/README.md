
### Summary

* A basic, single-threaded example that will be migrated later to use channels/goroutines.
* The `run.sh` and `test.sh` scripts are not idiomatic.
* I have tried to set up `GOPATH` etc but I just can't get it to work.
    - e.g. to use a 3rd-party package such as: https://github.com/stretchr/testify
* Since this is functional, I have to move on, as I'm more interested in
  the language than the insufferable module/package system.

### TODO

* idiomatic `go doc`
* idiomatic code format (use `go fmt`)
* refine import so that we don't need `prime.PrimeIndex`
* rename `New()` function for `PrimeIndex` ?
* rename `PrimeIndex` to `Primes`
