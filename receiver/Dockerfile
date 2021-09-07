# Start from the latest golang base image
FROM golang:latest

# Set Working DIR
WORKDIR /app

# Copy GO modules and expected hashes file
COPY . .

# Install dependencies
RUN go mod download

EXPOSE 8080

RUN go build -o receiver app/*.go

# Start the listening process via TCP
ENTRYPOINT [ "./receiver" ]
