# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13.5-alpine AS builder

WORKDIR /go/src/cw3guide
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir /cw3guide

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/cw3guide github.com/yfedoruck/cw3guide/cmd/bot

FROM alpine:3.11
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/cw3guide /go/src/cw3guide
COPY --from=builder /bin/cw3guide /bin/cw3guide

# Run the cw3guide by default when the container starts.
CMD ["/bin/cw3guide"]
EXPOSE 5000