FROM golang:1.19.2 as base
WORKDIR /gateway-service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway ./cmd/gateway


FROM alpine:3.16.2 as prod
COPY --from=base /gateway /usr/bin/gateway
CMD ["gateway"]
