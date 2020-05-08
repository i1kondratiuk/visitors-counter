################################################################
# Create a Docker image for GO.
################################################################
FROM golang:1.14.2-alpine3.11

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/visitors-counter/

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY config.yml /app/

# Build the Go app
RUN go build -o /app/visitors-counter .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go build`
CMD /app/visitors-counter