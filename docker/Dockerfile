FROM alpine:3.12

COPY ["artifacts/tatonka", "/usr/local/bin"]

ENTRYPOINT ["/usr/local/bin/tatonka"]
CMD ["trade","--paper=true","gdax.BTC-USD"]
