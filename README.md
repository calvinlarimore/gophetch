# gophetch
A **BLAZINGLY FAST** sysfetch written in Go.

![Screenshot](https://raw.githubusercontent.com/calvinlarimore/gophetch/main/img/gophetch.png)

## Why?
I wanted to mess around with goroutines and channels and a sysfetch made sense for async.

## How?
gophetch creates a bunch of goroutines with channels and waits on each channel to print the pre-formatted string that each goroutine pushes to it.

## Building and Installing
TL;DR: `go build`, copy to directory of your choice (usually `/usr/local/bin/`)
1. Clone the repo `git clone https://github.com/calvinlarimore/gophetch.git`
2. Move into the repo `cd gophetch`
3. Build binary `go build`
4. Copy binary to `/usr/local/bin/` (or anywhere else) `sudo cp gophetch /usr/local/bin/`
5. That's it!
