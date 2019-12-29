#docker build -t gcr.io/nandanks-151422/orders .
#docker run -d -p 8080:8080 --name orders -e DEBUG="true" gcr.io/nandanks-151422/orders
FROM golang:latest as builder
RUN useradd -U app
ADD . /go/src/orders 
WORKDIR /go/src/orders
COPY . /go/src/orders
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o /go/bin/orders

#without these certificates, we cannot verify the JWT token
FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/orders .
COPY --from=builder /go/src/orders/orders.json .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER app
CMD ["./orders"]