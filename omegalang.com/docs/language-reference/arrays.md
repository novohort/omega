---
sidebar_position: 6
---

# Working with Arrays

Arrays in Omega provide a flexible way to store and manipulate ordered collections of data. They can hold any type of value, including numbers, strings, and even other arrays.

## Array Literals

An array is defined by enclosing a comma-separated list of values within square brackets `[]`.

### Example

```omega
let my_arr { [1, "two", 420, true] }
```

This creates an array `my_arr` that contains an integer, a string, another integer, and a boolean value.

## Accessing Elements

Elements in an array can be accessed using their index. Omega arrays are zero-indexed, meaning the first element has an index of `0`.

### Syntax

```omega
let element { array[index] }
```

### Example

```omega
let my_arr { [1, "two", 420, true] }
let first_element { my_arr[0] } // 1
let second_element { my_arr[1] } // "two"
```

## Array Operations

Omega supports several built-in functions to manipulate array, such as `append`, `prepend`, `tail`, and more.

### append()

Adds an element to the end of an array.

```omega
let new_arr { append(my_arr, "last element!") }
```

### prepend()

Adds an element to the beginning of an array.

```omega
let new_arr { prepend(my_arr, "first element!") }
```

### tail()

Returns a new array containing all elements except the first.

```omega
let tail_arr { tail(my_arr) }
```

### head()

Returns a new array containing all elements except the last.

```omega
let head_arr { head(my_arr) }
```

### first()

Returns the first element of the array.

```omega
let first { first(my_arr) }
```

### last()

Returns the last element of the array.

```omega
let last { last(my_arr) }
```

### len()

Returns the number of elements in the array.

```omega
let length { len(my_arr) }
```
