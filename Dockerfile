# =========================
# 1️⃣ Build stage
# =========================
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY . .

# tidy + build static binary
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o autoscaler ./cmd/main.go

# =========================
# 2️⃣ Runtime stage (distroless)
# =========================
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# copy binary và .env
COPY --from=builder /app/autoscaler .
COPY --from=builder /app/.env .

USER nonroot:nonroot

EXPOSE 2112

ENTRYPOINT ["./autoscaler"]
