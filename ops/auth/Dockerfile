FROM golang:1.19.2 as base
WORKDIR /auth-service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth ./cmd/auth


FROM alpine:3.16.2 as prod
COPY --from=base /auth /usr/bin/auth
CMD ["sh", "-c", "auth migrate && auth run"]
