package main // << Go uses packages, `main` is the executable package. A package is a namespace.

import "fmt" // Import uses a package, in this case the `fmt` package has `print` functions, basically cstdio.h equivalent

func main() { // `func` keyword declares a function, `main` is obviously the entry point of an executable.

	// ; is redundant in `golang`, except for special cases
	fmt.Print("Hello World!")

	// << Main does not return a value, use os.Exit(code) for that
}

// Compile with `go build`

// We could do this online, you know. https://play.golang.org/
