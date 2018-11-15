# AA-Omise Challenge Solution

## Overview:
This is an attempt to solve the Go-lang Challenge. The solution covers the listed:

**Requirements:**

* [✅] Decrypt the file using a simple [ROT-128][2] algorithm.
* [✅] Make donations by creating a Charge via the [Charge API][0] for each row in the decrypted CSV.
* [✅] Produce a brief summary at the end.
* [✅] Handle errors gracefully, without stopping the entire process.
* [✅] Readable and maintainable code.

**Bonus:**

* [✅] Provides Go package structure.
* [✅] Provides rate limit feature on API calls.
* [✅] Uses Go routines and chanells to run as fast as possible on a multi-core CPU.
* [✅] Allocates as little memory as possible by processing data per line.
* [✅] Completes the entire process without leaving large trace of Credit Card numbers in memory, or on disk.
* [✅] Ensures reproducible builds on your workspace by providing a Makefile.
* [+] Provides test coverage.
* [+] Adds validation for all the provided inputs.
## How to Run:
Download the source repository from Github
```sh
cd $GOPATH/src/
git clone https://github.com/ArmanAA/Go-Challenge.git

```
Install the dependencies and build the project via Make.
```sh
make 
```
Run the provided example. The specified command reads the encrypted file and processes with 100 workers, where every api call has a 0.05 second (500 millisecond) rate limit.
```sh
make example
```

## Note:
A full list of available commands can be found by asking help.
```sh
make help
```

## Test:
Run all tests once.
```sh
make test
```
