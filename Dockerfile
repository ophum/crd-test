FROM alpine:3.7

COPY main /app/main
ENTRYPOINT ["/app/main"]
