---
title: "Split on spaces in Go"
date: 2020-09-18
draft: false
tags:
  - go
  - golang
images:
- https://opengraph.cluster.fun/opengraph/?siteTitle=Today%20I%20learnt...&title=Split%20on%20spaces%20in%20Go&tags=golang%2Cprogramming%2Carrays&image=https%3A%2F%2Fmarcusnoble.co.uk%2Fimages%2Fmarcus.jpg&twitter=Marcus_Noble_&github=AverageMarcus&website=www.MarcusNoble.co.uk
---

While looking to split a multiline and space separated string and not having any luck with `strings.Split()` I came across this somewhat oddly names function:

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