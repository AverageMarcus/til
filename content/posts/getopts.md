---
title: "CLI flag handling in Bash using getopts"
date: 2021-08-04T20:49:37+01:00
draft: false
tags:
  - bash
images:
- https://opengraph.cluster.fun/opengraph/?siteTitle=Today%20I%20learnt...&title=CLI%20flag%20handling%20in%20Bash%20using%20getopts&tags=bash&image=https%3A%2F%2Fmarcusnoble.co.uk%2Fimages%2Fmarcus.jpg&twitter=Marcus_Noble_&github=AverageMarcus&website=www.MarcusNoble.co.uk
---

I'm not sure how I've never come across this before but while looking through the [Scaleway Kosmos](https://www.scaleway.com/en/betas/#kuberneteskosmos) multi-cloud init script I dicovered the [`getopts`](https://www.man7.org/linux/man-pages/man1/getopts.1p.html) utility.

`getopts` makes it easier to parse arguments passed to a shell script by defining which letters your script supports. It supports both boolean and string style arguments but only supports single letter flags. (e.g. `-h` and not `--help`)

Example usage:

```sh
#!/bin/bash

NAME="World"
FORCE=false

showHelp() {
    echo "Usage: example.sh [args]"
    exit 0
}

while getopts 'hfn:' FLAG
do
  case $FLAG in
    h) showHelp ;;
    f) FORCE=true ;;
    n) NAME=$OPTARG ;;
    *) echo "Unsupported argument flag passed" ;;
  esac
done

echo "Hello, $NAME"

```

Notice the `:` following the `n`? That indicates that a value should follow the argument flag (`n` in this example) and will be made available as the `OPTARG` variable.