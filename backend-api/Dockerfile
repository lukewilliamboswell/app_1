FROM golang:latest as builder
ARG BIN_NAME=hello-api
COPY . /$BIN_NAME
WORKDIR /$BIN_NAME
RUN cd cmd/$BIN_NAME && CGO_ENABLED=0 GOOS=linux go build -o $BIN_NAME

FROM alpine:latest
ARG BIN_NAME=hello-api
COPY --from=builder /$BIN_NAME/ /$BIN_NAME/
WORKDIR /$BIN_NAME
ENTRYPOINT ./cmd/hello-api/hello-api
