# Go image for building the project
FROM golang:alpine3.16 as builder
RUN apk --no-cache add git dep ca-certificates

ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

RUN mkdir -p $GOPATH/src/github.com/creepzed/url-shortener-service

WORKDIR $GOPATH/src/github.com/creepzed/url-shortener-service

COPY go.mod .
COPY app/ app/

RUN go mod tidy
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./app/main.go

# Runtime image with scratch container
FROM scratch
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /app/

ENTRYPOINT ["/app/main"]