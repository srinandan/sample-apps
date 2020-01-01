#docker build -t gcr.io/nandanks-151422/orders-gql .
#docker run -d -p 8080:8080 --name orders-gql -e DEBUG="true" gcr.io/nandanks-151422/orders-gql
FROM golang:latest as builder
RUN useradd -U app
ADD . /go/src/orders-gql 
WORKDIR /go/src/orders-gql
COPY . /go/src/orders-gql
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' -o /go/bin/orders-gql server/server.go

#without these certificates, we cannot verify the JWT token
FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/orders-gql .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER app
CMD ["./orders-gql"]