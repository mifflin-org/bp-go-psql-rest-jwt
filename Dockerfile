# The Docker Image will be constructred in two stages
# 1: Building the Golang App and creating a binary on a Golang based Ubuntu Image
# 2: Copying the binary to a very simple Alpine OS and running it

# Stage 1

# Apline Forms are lighter versions of the Ubuntu
# golang:alpine gives us a base image with golang installed
# Assigning it as builder will help to copy content from this image
FROM golang:alpine as builder

LABEL maintainer="Aayush Joglekar <aayushjog@gmail.com>"

# Updates the OS and installs git
RUN apk update && apk add --no-cache git

# Sets the working directory as /app
# All the commands hereafter will be run in the /app folder
WORKDIR /app

# Copies the go mod & sum files to the /app folder
# Go Mod contains all the dependencies of the project
COPY go.mod go.sum ./

# Downloads the dependencies to appropriate location
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2

# Fetches the latest build of alpine
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Sets the working directory to /root/ for this container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 5000 to the outside world
EXPOSE 5000

#Command to run the executable
CMD ["./main"]