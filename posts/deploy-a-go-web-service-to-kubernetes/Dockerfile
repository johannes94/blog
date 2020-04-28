FROM golang:1.14.2 as buildstage

COPY ./src /src
WORKDIR /src
RUN go build -o ./bin/webservice

FROM ubuntu:latest
COPY --from=buildstage /src/bin/webservice /usr/bin/
ENTRYPOINT [ "/usr/bin/webservice" ]
