FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o main -a

# second stage #
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/BackupFiles ./BackupFiles
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]