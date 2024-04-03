---
sidebar_position: 7
---

# Working with Hashes

Hashes in Omega are powerful data structures that allow you to store key-value pairs, where each key is unique. This structure is ideal for quick data retrieval, organiation, and manipulation.

## Hash Literals

A hash literal is defined by enclosing key-value pairs in braces `{}`, with keys and values separated by a colon `:` and pairs separated by commas.

### Example

```omega
let my_hash {{ "name": "Omega", "type": "language", 42: "The Answer", true: "yes" }}
```

This creates a hash `my_hash` with string, integer, and boolean keys.

## Accessing Values

Values in a hash can be accessed using their corresponding keys.

### Syntax

```omega
let value { hash[key] }
```

### Example

```omega
let name { my_hash["name"] } // "Omega"
let answer { my_hash[42] } // "The Answer"
```
