ARG BUILDPLATFORM=linux/amd64
ARG TARGETPLATFORM=linux/arm64

## Build
FROM --platform=${BUILDPLATFORM} golang:1 AS build

ARG GO_BUILDER_GITHUB_TOKEN=missing_token
RUN git config --global "url.https://$GO_BUILDER_GITHUB_TOKEN@github.com/".insteadOf https://github.com/
ENV GOPRIVATE=github.com/teyz

WORKDIR /app

COPY . .

RUN go mod download \
    && GOARCH="arm64" go build cmd/main.go

## Deploy
FROM --platform=${TARGETPLATFORM} golang:1-alpine

USER nobody:nobody
WORKDIR /app
COPY --from=build /app/main ./main

## ??
EXPOSE 50052
ENTRYPOINT ["./main"]
