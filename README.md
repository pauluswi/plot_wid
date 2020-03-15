# PLOT Parking Lot
Contain repository to solve parking lot problem. 

There are 3 packages in the project, `main`, `app` and `dao` Dao have been used to seperate the storage logic.

## Installation and Compilation 
`make all  ` command will build and test. It should output a exec file `plot_wid` in the  project directory.

There is one external library used for testing `github.com/stretchr/testify` In order to run it properly this package needs to be downloaded. To download the package golang's internal tool `go get` is used and connected by make file.

## Testing

Both packages have unit tests cases, so `make test` command run the test cases for both the packages. Unit test utilize the internal package `testing` and exteranal package for more complex test cases.

## Run 
Binary can be run in the parent directory with of without file aragumant  by command `./plot_wid` or `./plot_wid {filepath}`