# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#docker build -t gcr.io/project-id/orders .
#docker run -d -p 8080:8080 --name orders -e DEBUG="true" gcr.io/project-id/orders
FROM golang:latest as builder
ADD . /go/src/orders 
WORKDIR /go/src/orders
COPY . /go/src/orders
RUN groupadd -r -g 20000 app && useradd -M -u 20001 -g 0 -r -c "Default app user" app && chown -R 20001:0 /go
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -a -ldflags '-s -w -extldflags "-static"' -o /go/bin/orders

#without these certificates, we cannot verify the JWT token
FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /
COPY --from=builder /go/bin/orders .
COPY --from=builder /go/src/orders/orders.json .
COPY --from=builder /etc/passwd /etc/group /etc/shadow /etc/
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
USER 20001
EXPOSE 8080
EXPOSE 8090
CMD ["./orders"]