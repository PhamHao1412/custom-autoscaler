FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o autoscaler ./cmd/main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /root/
COPY --from=builder /app/autoscaler .
EXPOSE 2112
USER nonroot:nonroot
ENTRYPOINT ["./autoscaler"]
