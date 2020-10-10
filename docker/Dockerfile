FROM alpine:3.12

COPY ["artifacts/tatanka", "/usr/local/bin"]

ENTRYPOINT ["/usr/local/bin/tatanka"]
CMD ["trade","--paper=true","gdax.BTC-USD"]
