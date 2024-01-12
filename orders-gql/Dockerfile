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

#docker build -t gcr.io/project-id/orders-gql .
#docker run -d -p 8080:8080 --name orders-gql -e DEBUG="true" gcr.io/project-id/orders-gql
FROM golang:1.21 as builder

ARG TAG
ARG COMMIT

ADD orders-gql/app /go/src/orders-gql/app
ADD orders-gql/data /go/src/orders-gql/data
ADD orders-gql/gql /go/src/orders-gql/gql
ADD orders-gql/models /go/src/orders-gql/models
ADD orders-gql/odr /go/src/orders-gql/odr
ADD orders-gql/resolvers /go/src/orders-gql/resolvers
ADD orders-gql/server /go/src/orders-gql/server
ADD internal/common /go/src/internal/common
ADD internal/datatypes /go/src/internal/datatypes

COPY orders-gql/go.mod orders-gql/go.sum /go/src/orders-gql/

WORKDIR /go/src/orders-gql

ENV GO111MODULE=on
RUN go mod tidy
RUN go mod download
RUN date +%FT%H:%I:%M+%Z > /tmp/date
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -buildvcs=true -a -gcflags='all="-l"' -ldflags='-s -w -extldflags "-static" -X main.version='${TAG}' -X main.commit='${COMMIT}' -X main.date='$(cat /tmp/date) -o /go/bin/orders-gql /go/src/orders-gql/server/server.go

FROM gcr.io/distroless/static-debian11:nonroot
LABEL org.opencontainers.image.url='https://github.com/srinandan/sample-apps/orders-gql' \
    org.opencontainers.image.documentation='https://github.com/srinandan/sample-apps' \
    org.opencontainers.image.source='https://github.com/srinandan/sample-apps/orders-gql' \
    org.opencontainers.image.vendor='Google LLC' \
    org.opencontainers.image.licenses='Apache-2.0' \
    org.opencontainers.image.description='A set of sample microservices'

COPY --from=builder /go/bin/orders-gql /
EXPOSE 8080
CMD [ "/orders-gql" ]