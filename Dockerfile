# Using multistage 

# Build Project
FROM golang:1.20-alpine as go-builder

LABEL maintainer="Maryam Sanaei <sanaei25272@gmail.com>"

WORKDIR /app

COPY . . 

# Both build and runtime container are the same, so we can keep enable CGO
RUN apk add -u -t curl git && \
    go build -o server *.go && \
    rm -rf /vat/cache/apk/*


# Runtime Container
# We can use scratch image instead of alpine to have smaller image
FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/server /app/server
EXPOSE 8080
ENTRYPOINT ["/app/server"]
