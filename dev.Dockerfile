FROM golang:1.13

# Set the Current Working Directory inside the container
WORKDIR /go-hexagonal

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -exclude-dir=.git -build="go build -o go-hexagonal ./cmd/api" -command="./go-hexagonal -logtostderr"
