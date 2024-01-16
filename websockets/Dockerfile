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

FROM golang:1.21 as builder

ARG TAG
ARG COMMIT

COPY websockets/server /go/src/websockets
ADD internal/common /go/src/internal/common

WORKDIR /go/src/websockets

ENV GO111MODULE=on
RUN go mod tidy
RUN go mod download
RUN date +%FT%H:%I:%M+%Z > /tmp/date
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -buildvcs=true -a -gcflags='all="-l"' -ldflags='-s -w -extldflags "-static" -X main.version='${TAG}' -X main.commit='${COMMIT}' -X main.date='$(cat /tmp/date) -o /go/bin/websockets /go/src/websockets/main.go


FROM gcr.io/distroless/static-debian11:nonroot
LABEL org.opencontainers.image.url='https://github.com/srinandan/sample-apps/websockets' \
    org.opencontainers.image.documentation='https://github.com/srinandan/sample-apps' \
    org.opencontainers.image.source='https://github.com/srinandan/sample-apps/websockets' \
    org.opencontainers.image.vendor='Google LLC' \
    org.opencontainers.image.licenses='Apache-2.0' \
    org.opencontainers.image.description='A set of sample microservices'

COPY --from=builder /go/bin/websockets /
EXPOSE 3000
CMD [ "/websockets" ]