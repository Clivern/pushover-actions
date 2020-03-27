FROM golang:1.14.1

LABEL "com.github.actions.name"="pushover-actions"
LABEL "com.github.actions.description"="Push notifications for github repository changes through pushover"
LABEL "com.github.actions.icon"="package"
LABEL "com.github.actions.color"="red"

LABEL "repository"="https://github.com/Clivern/pushover-actions"
LABEL "homepage"="http://github.com/clivern"
LABEL "maintainer"="Clivern <hello@clivern.com>"

ARG PO_VERSION=0.0.6

ENV GO111MODULE=on

RUN mkdir -p /app

RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/pushover-actions/releases/download/${PO_VERSION}/pushover-actions_${PO_VERSION}_Linux_x86_64.tar.gz | tar xz

RUN rm LICENSE
RUN rm README.md

COPY entrypoint.sh /app

RUN chmod +x /app/entrypoint.sh

RUN chmod +x /app/pushover-actions

ENTRYPOINT ["/app/entrypoint.sh"]