# gophetch
A **BLAZINGLY FAST** sysfetch written in Go.

## Why?
I wanted to mess around with goroutines and channels and a sysfetch made sense for async.

## How?
gophetch creates a bunch of goroutines with channels and waits on each channel to print the pre-formatted string that each goroutine pushes to it.