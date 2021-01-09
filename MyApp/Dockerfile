FROM golang:1.14 as builder


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/mathapp

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o mathapp

FROM debian:buster

ENV APP_USER app
ENV APP_HOME /go/src/mathapp

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME


COPY --chown=0:0 --from=builder $APP_HOME/mathapp $APP_HOME

EXPOSE 5555
USER $APP_USER
CMD ["./mathapp"]