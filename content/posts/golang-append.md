---
title: "Golang's append mutates the provided array"
date: 2020-10-30T12:49:37+01:00
draft: false
tags:
  - golang
images:
- /images/golang-append.gif
---

A word of warning when using `append()` in Golang...

When using a slice of an array as the first parameter when calling `append` if the length of the resulting array is less than that of the initial array then the values in the initial array will be overridden by the new values being appended.

For example, if we use `append` to take the first two values from the array called `first` and all the values from the array called `second` we can see that `first` is being mutated unexpectedly.

```go
package main

import "fmt"

func main() {
  first := []int{0, 1, 2, 3, 4, 5, 6}
  second := []int{4, 5, 6}

  fmt.Println(first)
  // -> [0 1 2 3 4 5 6]
  fmt.Println(append(first[:2], second...))
  // -> [0 1 4 5 6]
  fmt.Println(first)
  // -> [0 1 4 5 6 4 5 6]
}
```

This is _only_ an issue when the resulting array is shorter than the array passed in as the first parameter.
