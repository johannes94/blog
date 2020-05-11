FROM golang:alpine as buildstage

COPY ./src /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/webservice

FROM scratch
COPY --from=buildstage /src/bin/webservice /webservice

ENTRYPOINT [ "./webservice" ]
