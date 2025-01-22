# Go: Unhandled File Errors

This repository demonstrates a common error in Go programs: failing to check for errors returned by functions.  Specifically, it shows how not checking for errors when opening a file can lead to unexpected behavior and program crashes.

## The Bug

The `bug.go` file contains a program that attempts to open and read from a file.  However, it neglects to check the error returned by `os.Open`. If the file doesn't exist or there's a problem accessing it, the program will panic.

## The Solution

The `bugSolution.go` file provides a corrected version of the program. It demonstrates the proper way to handle errors by checking the return value of `os.Open` and using an `if` statement to gracefully handle the case where the file cannot be opened.  This prevents the program from crashing and allows for informative error messages.

## How to Run

1. Clone this repository.
2. Navigate to the repository's directory.
3. Run the buggy program using `go run bug.go`.
4. Observe the panic.
5. Run the corrected program using `go run bugSolution.go`.
6. Observe that the program handles the error correctly. 

This simple example highlights the importance of diligent error handling in robust Go programming.