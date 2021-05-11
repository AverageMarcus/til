---
title: "Named returns in Go functions"
date: 2020-10-05T15:50:00
draft: false
tags:
  - golang
images:
- https://opengraph.cluster.fun/opengraph/?siteTitle=Today%20I%20learnt...&title=Named%20returns%20in%20Go%20functions&tags=golang%2Cprogramming&image=https%3A%2F%2Fmarcusnoble.co.uk%2Fimages%2Fmarcus.jpg&twitter=Marcus_Noble_&github=AverageMarcus&website=www.MarcusNoble.co.uk
---

While debugging some issues I was having with the AWS Golang SDK I discovered it was possible to name your function return values (pointers) and then set them within your function body without needing to explicitly return them at the end.

E.g.

```go
package main

import "fmt"

var greeting = "Hello, world"

func main() {
	fmt.Println(*test())
}

func test() (returnVal *string) {
	returnVal = &greeting
	return
}
```

Note the single `return` at the end of the function.

I'm not really sure if this is a useful feature, feels more like it'd make code harder to read and could lead to some pretty nasty unintended mistakes when someone unfamilure with the code comes to make changes. Interesting either way.

I'd be interested if anyone has any examples of where this kind of thing is beneficial.