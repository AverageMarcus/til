---
title: "Don't Reuse API Keys"
date: 2020-10-03T12:49:37+01:00
draft: false
tags:
  - cli
  - credentials
images:
- https://opengraph.cluster.fun/opengraph/?siteTitle=Today%20I%20learnt...&title=Don't%20Reuse%20API%20Keys&tags=cli%2Ccredentials&image=https%3A%2F%2Fmarcusnoble.co.uk%2Fimages%2Fmarcus.jpg&twitter=Marcus_Noble_&github=AverageMarcus&website=www.MarcusNoble.co.uk
---

Not a technical post today, more of a reminder to myself not to reuse API keys for different purposes. In this instance I reset the credentials I had labelled "Terraform" which I just so happened to also be using In [Harbor](https://goharbor.io/) to connect to my S3 bucket. Que 2 hours of me later trying to figure out why I couldn't pull or push any images.
