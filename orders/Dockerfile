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

FROM golang:1.21 as builder

ARG TAG
ARG COMMIT

ADD orders/apis /go/src/orders/apis
ADD orders/app /go/src/orders/app
ADD orders/data /go/src/orders/data
ADD orders/db /go/src/orders/db
ADD internal/common /go/src/internal/common
ADD internal/datatypes /go/src/internal/datatypes

COPY orders/main.go /go/src/orders/main.go
COPY orders/go.mod orders/go.sum /go/src/orders/

WORKDIR /go/src/orders

ENV GO111MODULE=on
RUN go mod tidy
RUN go mod download
RUN date +%FT%H:%I:%M+%Z > /tmp/date
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -buildvcs=true -a -gcflags='all="-l"' -ldflags='-s -w -extldflags "-static" -X main.version='${TAG}' -X main.commit='${COMMIT}' -X main.date='$(cat /tmp/date) -o /go/bin/orders /go/src/orders/main.go

FROM gcr.io/distroless/static-debian11:nonroot
LABEL org.opencontainers.image.url='https://github.com/srinandan/sample-apps/orders' \
    org.opencontainers.image.documentation='https://github.com/srinandan/sample-apps' \
    org.opencontainers.image.source='https://github.com/srinandan/sample-apps/orders' \
    org.opencontainers.image.vendor='Google LLC' \
    org.opencontainers.image.licenses='Apache-2.0' \
    org.opencontainers.image.description='A set of sample microservices'

COPY --from=builder /go/bin/orders /
EXPOSE 8080
CMD [ "/orders" ]