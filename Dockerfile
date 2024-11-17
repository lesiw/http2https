FROM golang:latest AS builder
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 go build -o app .
FROM scratch
COPY --from=builder /work/app /app
ENTRYPOINT ["/app"]
