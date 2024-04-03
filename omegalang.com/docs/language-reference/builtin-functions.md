---
sidebar_position: 4
---

# Builtin Functions

Omega offers a variety of built-in functions that are integral to handling common operations across different data types, such as arrays and strings.

## Array Functions

Omega includes several functions specifically designed for array manipulation, making it easier to work with collections of data.

### tail()

Removes the first element of an array and returns the remaining elements.

```omega
let my_arr { [1, 2, 3] };
let result { tail(my_arr) }; // returns [2, 3]
```

### head()

Removes the last element of an array and returns the remaining elements.

```omega
let my_arr { [1, 2, 3] };
let result { head(my_arr) }; // returns [1, 2]
```

### prepend()

Adds an element to the beginning of an array.

```omega
let my_arr { [1, 2, 3] };
let result { prepend(my_arr, 42) }; // returns [42, 1, 2, 3]
```

### append()

Adds an element to the end of an array.

```omega
let my_arr { [4, 2, 0] };
let result { append(my_arr, 69) }; // returns [4, 2, 0, 69]
```

### last()

Returns the last element of an array without modifiying the array.

```omega
let my_arr { [1, 2, 3] };
let result { last(my_arr) }; // returns 3
```

### first()

Returns the first element of an array without modifying the array.

```omega
let my_arr { [1, 2, 3] };
let result { first(my_arr) }; // returns 1
```

## String and Array Length

The `len()` function can be used to find the length of strings and arrays.

```omega
let arr_len { len([4, 2, 0, 6, 9]) }; // returns 5
let str_len { len("Hello, Omega!") }; // returns 13
```

## Output Function

Omega provides the `out` function to print values to the console, aiding in debugging and interactive development, while also returning `null`.

```omega
let print { out("Hello, Omega!") };
```
