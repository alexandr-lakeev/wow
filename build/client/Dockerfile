FROM golang:1.21.1 as build

ENV BIN_FILE /opt/world-of-wisdom/client
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}
COPY .. ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/client/main*

FROM alpine:3.9

ENV BIN_FILE /opt/world-of-wisdom/client
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/world-of-wisdom/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

RUN sleep 1
CMD ${BIN_FILE} -config ${CONFIG_FILE}
