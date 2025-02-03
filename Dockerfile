# workspace (GOPATH) configured at /go
FROM golang:1.20 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.com/ildambackend/ildam_go_template_service
WORKDIR $GOPATH/src/gitlab.com/ildambackend/ildam_go_template_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/eld_go_template_service /

FROM alpine
COPY --from=builder eld_go_template_service .
RUN mkdir config

ENV ENV_FILE_PATH=/app/.env
RUN apk add --no-cache curl

ENTRYPOINT ["/eld_go_template_service"]
