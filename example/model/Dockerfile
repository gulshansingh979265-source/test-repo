# ✅ Use official Golang image
FROM golang:1.21

# ✅ Set working directory inside container
WORKDIR /app

# ✅ Copy all project files
COPY . .

# ✅ Download dependencies
RUN go mod tidy

# ✅ Build the Go app
RUN go build -o main .

# ✅ Run the executable
CMD ["./main"]
