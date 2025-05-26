FROM golang:1.21

RUN apt update && apt install ca-certificates -yq
RUN echo "запуск сборки мигратора"

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

WORKDIR /migrations

COPY ./postgres/ ./postgres/

ENV TZ=Etc/UTC

ENTRYPOINT ["sh"]
