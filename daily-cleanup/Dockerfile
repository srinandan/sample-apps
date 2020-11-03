FROM alpine

RUN apk add --update \
    curl openssl jq bash \
    && rm -rf /var/cache/apk/

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl && chmod +x ./kubectl && mv ./kubectl /usr/local/bin/kubectl

COPY entrypoint.sh /

CMD ["./entrypoint.sh"]
