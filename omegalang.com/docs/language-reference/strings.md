---
sidebar_position: 5
---

# Working with Strings

Strings in Omega are essential for handling text data. They are denoted by double quotes (`"`) and can be manipulated and queried through built-in functions.

## String Literals

A string literal in Omega is any sequence of characters in double quotes. For example:

```omega
let greeting { "Hello, Omega!" };
```

This creates a string variable `greeting` that contains the text `Hello, Omega!`

### Concatenation

Strings can be conctenated using the `+` operator, allowing you to join two or more strings into one.

### Syntax

```omega
let concatenated_string { string_1 + string_2 };
```

### Example

```omega
let hello { "Hello, " };
let world { "Omega!" };
let greeting { hello + world } // "Hello, Omega!"
```

This combines `hello` and `world` into a single string `greeting`.

## Length

To get the length of a string, use the `len` function, which returns the number of characters in the string.

### Syntax

```omega
let length { len(string) };
```

### Example

```omega
let message { "Omega" };
let msg_length { len(message) }; // returns 5
```
