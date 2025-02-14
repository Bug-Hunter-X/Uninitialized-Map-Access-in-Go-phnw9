# Uninitialized Map Access in Go

This repository demonstrates a common but subtle error in Go: accessing elements of an uninitialized map.  Attempting to access a key in an uninitialized map does *not* result in a runtime panic. Instead, it returns the zero value for the map's value type (often 0 for integers, "" for strings, etc.). This silent failure can be difficult to detect and debug.

The `bug.go` file shows the problematic code.  The `bugSolution.go` provides a corrected version using a safe access pattern.

## How to reproduce the bug
1. Clone the repository.
2. Run the `bug.go` file using `go run bug.go`.  Observe the output. The output should be 0,0,<nil> pointer dereference panic. 

## Solution
The solution lies in checking if the map exists before accessing its elements using the `if val, ok := m["key"]; ok` syntax. See `bugSolution.go` for the correct approach.