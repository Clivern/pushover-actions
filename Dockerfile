FROM golang:1.13.8

LABEL "com.github.actions.name"="pushover-actions"
LABEL "com.github.actions.description"="Pushover A Notifications for Github Repository Changes"
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="red"

LABEL "repository"="https://github.com/Clivern/pushover-actions"
LABEL "homepage"="http://github.com/clivern"
LABEL "maintainer"="Clivern <hello@clivern.com>"

ENV PUSHOVER_TOKEN=
ENV PUSHOVER_USER=

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o pushover-actions .

ENTRYPOINT ["./pushover-actions"]