# syntax=docker/dockerfile:1
FROM golang:1.22-alpine as builder
RUN apk update && \
    apk upgrade && \
    apk --no-cache add git upx wget ca-certificates
RUN mkdir /build
WORKDIR /build
RUN wget https://tranco-list.eu/top-1m.csv.zip

ADD . /build/
ARG COMMIT
ARG LASTMOD
RUN echo "INFO: building for $COMMIT on $LASTMOD"
RUN \
    CGO_ENABLED=0 GOOS=linux go build \
    -a \
    -installsuffix cgo \
    -ldflags "-s -w -X main.COMMIT=$COMMIT -X main.LASTMOD=$LASTMOD -extldflags '-static'" \
    -o siterank \
    *.go \
    && upx siterank
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/top-1m.csv.zip /app/
COPY --from=builder /build/siterank /app/
ENV PORT 4000
ENTRYPOINT ["/app/siterank"]
