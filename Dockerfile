FROM golang:1.14-alpine as go-builder

ENV GO111MODULE=on
WORKDIR /app

ADD social-image-gen .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o social-image-gen main.go

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

COPY --from=go-builder /app/social-image-gen /social-image-gen

WORKDIR /app
ADD . .

ADD ./social-image-gen/fonts/ /fonts
RUN /social-image-gen --mode posts
RUN hugo -d /usr/share/nginx/html/

FROM nginx:latest
COPY --from=builder /usr/share/nginx/html/ /usr/share/nginx/html/
