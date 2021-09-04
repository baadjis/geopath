# The base go-image, using latest stable version when this was written
FROM golang:1.17.0-alpine3.14 AS builder

# To fix go get and build with cgo
RUN apk add --no-cache --virtual .build-deps \
    bash \
    gcc \
    git \
    musl-dev

RUN mkdir build


COPY . /build
WORKDIR /build

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o webserver .
RUN adduser -S -D -H -h /build webserver
USER webserver

FROM scratch
COPY --from=builder /build/webserver /app/

# Copy static files
COPY statics app/statics

COPY public app/public


WORKDIR /app

EXPOSE 10000
CMD ["./webserver"]