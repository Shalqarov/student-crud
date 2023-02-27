FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . ./
RUN go mod tidy
RUN go build -o /school

FROM alpine:latest
WORKDIR /
COPY --from=builder /school /school
COPY --from=builder app/.env .env
CMD ["/school"]