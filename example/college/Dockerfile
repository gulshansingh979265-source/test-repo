# ✅ Step 1: Use Official Golang Image
FROM golang:1.21

# ✅ Step 2: Set Working Directory inside Container
WORKDIR /app

# ✅ Step 3: Copy All Project Files
COPY . .

# ✅ Step 4: Download Dependencies
RUN go mod tidy

# ✅ Step 5: Build the Go App
RUN go build -o main .

# ✅ Step 6: Run the App
CMD ["./main"]
