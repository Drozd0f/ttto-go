FROM golang:1.19.2 as base
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ttto ./cmd/monolith


FROM alpine:3.16.2 as prod

COPY --from=base /ttto /usr/bin/ttto
CMD ["sh", "-c", "ttto migrate && ttto run"]
