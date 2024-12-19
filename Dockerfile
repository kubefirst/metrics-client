FROM alpine:3.17.2

COPY metrics-client /usr/bin/

ENTRYPOINT ["/usr/bin/metrics-client"]
