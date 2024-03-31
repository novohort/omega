---
sidebar_position: 3
---

# Output

Omega offers a straightforward mechanism for displaying output to the console, making it easy for developers to output messages. This functionality is provided through the built-in `out` function.

## `out` Function

The `out` function is designed to output text to the console. It simplifies the process of displaying messages, aiding in debugging and user interaction within Omega programs.

### Usage

To use the `out` function, you simply call it with a string literal as its argument. The syntax is as follows:

```omega
out("text to print");
```

### Example

Here's a basic example demonstrating the `out` function in action:

```omega
fn main(): void {
  out("Hello, Omega!");
}
```

This code snippet defines a `main` function, which is the entry point of an Omega program. Within the `main` function, the `out` function is called with the string `"Hello, Omega!"` as its argument. When run, this program will print the message `Hello, Omega!` to the console.
