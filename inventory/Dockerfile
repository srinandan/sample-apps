#docker build -t gcr.io/nandanks-151422/inventory .
#docker run -d -p 8080:8080 --name inventory -e DEBUG="true" gcr.io/nandanks-151422/inventory
FROM golang:latest as builder
RUN useradd -U app
ADD . /go/src/inventory 
WORKDIR /go/src/inventory
COPY . /go/src/inventory
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/inventory

#without these certificates, we cannot verify the JWT token
FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/inventory .
COPY --from=builder /go/src/inventory/items.json .
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER app
CMD ["./inventory"]