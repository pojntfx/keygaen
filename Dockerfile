FROM golang:1.20.4-alpine3.18 as build

WORKDIR /usr/app

RUN apk update && apk add make g++ gcc npm

COPY ./cmd /usr/app/cmd
COPY ./docs /usr/app/docs
COPY ./pkg /usr/app/pkg
COPY ./web /usr/app/web
COPY ./package-lock.json /usr/app/package.json
COPY ./package.json /usr/app/package.json
COPY ./go.mod /usr/app/go.mod
COPY ./go.sum /usr/app/go.sum
COPY ./Makefile /usr/app/Makefile

RUN make depend
RUN make build

FROM golang:1.20.4-alpine3.18

WORKDIR /usr/app

COPY --from=build /usr/app /usr/app
COPY --from=build /go/pkg/mod /go/pkg/mod

ENTRYPOINT [ "go", "run", "/usr/app/cmd/keygaen-pwa/main.go" ,"-serve" ]
