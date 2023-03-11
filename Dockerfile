FROM golang:1.20-alpine AS go-builder

RUN apk update && apk add ca-certificates

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /auth-service .

# Final image
FROM scratch AS final-image

COPY --from=go-builder /auth-service /auth-service

COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 80 443 8080

CMD [ "/auth-service" ]
