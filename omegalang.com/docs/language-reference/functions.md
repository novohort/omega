---
sidebar_position: 2
---

# Functions

Functions are central to structuring code in Omega, allowing you to define reusable blocks of code. This page guides you through the syntax and usage of functions in Omega.

## Defining Functions

In Omega, functions are defined using the `fn` keyword. This is followed by the function name, a pair of parentheses `()` indicating parameters (currently, Omega does not support parameters), a colon `:`, the return type, and the function body enclosed in braces `{}`.

### Syntax

The basic syntax for defining a function in Omega is as follows:

```omega
fn function_name(): return_type {
  // function body
}
```

Since Omega is in its early development stages, it currently supports functions without parameters and return values. The return type is specified as `void` to indicate the function does not return a value.

### Example

Here is a simple example of a function in Omega:

```omega
fn main(): int {
  return 1 + 5;
}
```

This example defines a function named `main` that outputs "Hello, world!" to the console. The `main` function serves as the entry point of an Omega program.
