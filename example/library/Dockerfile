# ✅ Use official Go image
FROM golang:1.21

# ✅ Set working directory
WORKDIR /app

# ✅ Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod tidy

# ✅ Copy all source code
COPY . .

# ✅ Build the Go application
RUN go build -o library-app

# ✅ Run the app
CMD ["./library-app"]
