# The Omega Programming Language

This is the main source code repository for Omega. It contains the interpreter, Alpha.

## About Omega

Omega is (obviously) in its extreme infancy stage; it hasn't even attached to the uterus lining yet. Be advised that everything has the potential to change as time goes on, but we're planning to prevent any breaking changes as much as possible so we can maintain a singular backwards-compatible spec as opposed to Rust's tendency to break previous versions.

Omega is designed to be a general-purpose programming language. Sure, it'd be great if we can set it up so it can accept multiple programming paradigms (functional, OOP, etc) in order to ease adoption from developers of all flavors, but that's wildly ambitious especially at this stage.

## About Alpha, the Omega interpreter

Alpha originally started as a compiler written in Javascript and leveraging rustc, but has since changed direction.

Starting in v0.3.0, Alpha is a live interpreter (REPL) for Omega, and is written in Go.

## Getting Started

You'll need to install Go in order to make use of Alpha (and thereby, Omega).

Once that's done, you may launch the REPL by running `go run alpha.go` in your terminal just like you would with any other REPL. Be sure that you're in the correct directory (`interpreter`)
