# ---------- Build stage ----------
    FROM golang:1.26-alpine AS builder

    WORKDIR /app
    
    RUN apk add --no-cache git
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    COPY migrations ./migrations  
    
    RUN go build -o main cmd/server/main.go
    
    # ---------- Run stage ----------
    FROM alpine:latest
    
    WORKDIR /app
    
    RUN apk add --no-cache ca-certificates
    
    COPY --from=builder /app/main .
    COPY --from=builder /app/migrations ./migrations  
    
    EXPOSE 8080
    
    CMD ["./main"]