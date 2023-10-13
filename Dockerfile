# Use a minimal base image for your Go application
FROM golang:1.21 AS build

# Set the working directory within the container
WORKDIR /dba

# Copy the Go application source code and assets (HTML, CSS, JS)
COPY . .

# Build the Go application
RUN go build -o dba

# Copy static assets (HTML and JavaScript files) to the container
COPY static/ /dba/static/

# Expose the port your Go application is listening on
EXPOSE 8080

# Set the command to run your Go application
CMD ["./dba"]