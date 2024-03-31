---
sidebar_position: 2
---

# Getting Started

## Installation

Currently, Omega compiles to Rust, so you'll need to have rustc installed on your machine. The compiler is built in JavaScript, so you're also going to need Node. Please refer to the [Rust installation guide](https://www.rust-lang.org/tools/install) and the [Node installation guide](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) for instructions.

After everything is installed, you can clone the Omega repository from GitHub:

```bash
git clone https://github.com/novohort/omega.git
cd omega
```

## Your First Omega Program

To get started with Omega, let's write a simple program that prints "Hello, world!" to the console.

Create a file named `hello_world.o` and add the following Omega code:

```omega
fn main(): void {
  out("Hello, world!");
}
```

To compile your Omega program, run:

```bash
node alpha.js <omega_file.o> <output_program.exe>
```

This will compile it into an executable. Run the executable to see "Hello, world!" printed to the console.