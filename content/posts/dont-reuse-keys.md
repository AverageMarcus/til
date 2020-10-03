---
title: "Don't Reuse Keys"
date: 2020-10-03T12:49:37+01:00
draft: false
tags:
  - cli
  - credentials
images:
- /images/dont-reuse-keys.gif
---

Not a technical post today, more of a reminder to myself not to reuse API keys for different purposes. In this instance I reset the credentials I had labelled "Terraform" which I just so happened to also be using In [Harbor](https://goharbor.io/) to connect to my S3 bucket. Que 2 hours of me later trying to figure out why I couldn't pull or push any images.
