FROM golang:1.13 as build_base

# Set the Current Working Directory inside the container
WORKDIR /go-hexagonal

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

FROM build_base as builder

WORKDIR /go-hexagonal

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/go-hexagonal ./cmd/API


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=builder /go/bin/go-hexagonal .

EXPOSE 8080

CMD ["./go-hexagonal"] 