FROM golang:1.24.1 AS builder

WORKDIR /gnss_corr

COPY go.* ./
RUN echo "запуск сборки приложения"

RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg

WORKDIR /gnss_corr/cmd

ARG BUILD_WITH_RACE
ARG VERSION

RUN go build -ldflags="-X main.Version=$VERSION" $BUILD_WITH_RACE -o /bin/gnss_corr

FROM ubuntu:24.04

RUN apt update && apt install ca-certificates -yq

COPY --from=builder /bin/gnss_corr /bin/gnss_corr

ENV TZ=Etc/UTC

ENTRYPOINT ["/bin/gnss_corr"]