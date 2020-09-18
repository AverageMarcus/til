---
title: "Split on spaces in Gp"
date: 2020-09-18
draft: false
tags:
  - go
  - golang
images:
- /images/golang-split-by-space.gif
---

While looking to split a multiline and space separated string and not having any look with `strings.Split()` I came across this somewhat oddly names function:

```go
import (
    "fmt"
    "strings"
)

func main() {
    input := `This is 
a multiline, space 
separated string`

    output := strings.Fields(input)

    fmt.Println(output) // ["This", "is", "a", "multiline,", "space", "separated", "string"]
}
```