This directory contains the code for a video I made: [Using Delve to Examine Memory in Go](https://www.youtube.com/watch?v=sV5f1dF8ZU0)

# Agenda

- Delve is a debugger for Go
- Use examinememory to explore raw memory layout of a slice
- Use new -x flag for evaluating expressions, see [1]
- Needs latest version from source, see below

[1] https://github.com/go-delve/delve/pull/2385

# Install delve from source:

git clone https://github.com/go-delve/delve
cd delve
go install github.com/go-delve/delve/cmd/dlv

# Example Code From This Video

https://github.com/felixge/dump/tree/master/delve-examine-memory
