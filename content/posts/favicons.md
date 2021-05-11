---
title: "How to get the favicon of any site"
date: 2020-11-10T09:49:37+01:00
draft: false
tags:
- 
images:
- https://opengraph.cluster.fun/opengraph/?siteTitle=Today%20I%20learnt...&title=How%20to%20get%20the%20favicon%20of%20any%20site&tags=favicon&image=https%3A%2F%2Fmarcusnoble.co.uk%2Fimages%2Fmarcus.jpg&twitter=Marcus_Noble_&github=AverageMarcus&website=www.MarcusNoble.co.uk
---

If you ever find yourself needing to display a small icon for a 3rd party URL but don't want to have to crawl the site to pull out the favicon URL then you can make use of a Google CDN:

```
https://s2.googleusercontent.com/s2/favicons?domain_url=https://www.bbc.co.uk/
```

Example: ![](https://s2.googleusercontent.com/s2/favicons?domain_url=https://www.bbc.co.uk/)

You can even provide any page, not just the root URL.

e.g. `https://s2.googleusercontent.com/s2/favicons?domain_url=https://www.bbc.co.uk/news/newsbeat-54838856`: ![](https://s2.googleusercontent.com/s2/favicons?domain_url=https://www.bbc.co.uk/news/newsbeat-54838856)
