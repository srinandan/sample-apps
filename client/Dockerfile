FROM alpine

RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/

COPY entrypoint.sh /

CMD ["./entrypoint.sh"]