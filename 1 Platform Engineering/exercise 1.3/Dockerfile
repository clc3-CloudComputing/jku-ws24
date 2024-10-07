FROM golang:1.13-alpine

# Set /src as working directory
WORKDIR /src

# Copy source code to working directory
COPY main.go go.* ./

# Build the GO app as myapp binary and move it to /usr/
RUN CGO_ENABLED=0 go build -o /usr/myapp

# Mark port for port publishing
EXPOSE 8888

# Run the service myapp when a container of this image is launched
CMD ["/usr/myapp"]