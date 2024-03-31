---
sidebar_position: 4
---

# Variables

## Overview

Omega supports variables as a fundamental aspect of the language, allowing you to store and manipulate data. Variables in Omega are typed, ensuring that operations on these variables are safe and predictable.

## Variable Initialization

In Omega, you initialize variables using the `let` and `const` keywords. The `let` keyword is used for mutable variables, whereas `const` is used for immutable variables.

### Syntax

Mutable variable initialization:

```omega
let variable_name: type { value };
```

Immutable variable initialization:

```omega
const variable_name: type { value };
```

### Types

Currently, Omega supports two basic types for variables:

- `int`: A signed integer.
- `uint`: An unsigned integer.


### Examples

Declaring a mutable integer variable with an initial value:

```omega
let my_var: int { 5 };
```

Declaring an immutable unsigned integer variable with an initial value:

```omega
const my_const: uint { 10 };
```
