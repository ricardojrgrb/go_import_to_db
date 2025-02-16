# use official Golang image
FROM golang:1.22.5-alpine3.20

# set working directory
WORKDIR /app

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o app .

#EXPOSE the port
EXPOSE 8000

# Run the executable
CMD ["./app"]