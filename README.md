# The Omega Programming Language

This is the main source code repository for Omega. It contains the compiler, standard library, and documentation.

## About Omega (and its compiler, Alpha)

Omega is (obviously) in its extreme infancy stage; it hasn't even attached to the uterus lining yet. Be advised that everything has the potential to change as time goes on, but we're planning to prevent any breaking changes as much as possible so we can maintain a singular backwards-compatible spec as opposed to Rust's tendency to break previous versions.

Omega is designed to be a general-purpose programming language. Sure, it'd be great if we can set it up so it can accept multiple programming paradigms (functional, OOP, etc) in order to ease adoption from developers of all flavors, but that's wildly ambitious especially at this stage. I mean, fuck, we barely got "hello world" to work.

Alpha, the Omega compiler, is written with JavaScript and utilizes rustc for the compilation step.

## Getting Started

You'll need two things installed in order to work with Omega: Node, and rustc.

When you're ready to compile, simply run `node alpha.js <omega_file.o> <output_program.exe>`, and Alpha will tokenize your Omega file, parse it, convert it to valid Rust code, and then compile it to an executable. Once that's done, Alpha automatically cleans up the leftover files that were generated during Rust's compilation step.

## A Brief Introduction

Omega's syntax is partly inspired by Rust's syntax, with the goal being to provide a more consistent experience. Below is an example of Omega code, compared to its valid Rust counterpart.

```
fn main(): void {
  out("Hello, world!");
}
```

```
fn main() {
  println!("Hello, world!");
}
```

As you can see, there are a few key differences.

First, everything relevant has an explicit type declaration. This is to prevent any confusion about what may or may not be. In Rust, a function's return type is implied unless otherwise specified (because in some cases you may not return anything). This is all fine and dandy, but we're aiming for consistency with Omega, so if your function doesn't return anything then its return type is `void`.

Likewise, all variables will have explicit type declarations.

Additionally, the `println!` macro is simplified to `out` (with its input counterpart planned to be `in`).
