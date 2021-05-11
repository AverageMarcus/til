FROM debian:buster AS builder
RUN apt-get -qq update \
	&& DEBIAN_FRONTEND=noninteractive apt-get -qq install -y --no-install-recommends libstdc++6 python-pygments git ca-certificates asciidoc curl \
	&& rm -rf /var/lib/apt/lists/*
ENV HUGO_VERSION 0.74.3
ENV HUGO_BINARY hugo_extended_${HUGO_VERSION}_Linux-64bit.deb
RUN curl -sL -o /tmp/hugo.deb \
    https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/${HUGO_BINARY} && \
    dpkg -i /tmp/hugo.deb && \
    rm /tmp/hugo.deb
WORKDIR /app
ADD . .
RUN hugo -d /usr/share/nginx/html/

FROM nginx:latest
COPY --from=builder /usr/share/nginx/html/ /usr/share/nginx/html/
