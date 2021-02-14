# DO NOT use in production!
# Dockerfile for local development

FROM golang:1.15

# install the watcher
RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /project
COPY ./ /project

ENTRYPOINT ["CompileDaemon", "--build=go build -mod=mod", "-log-prefix=false", "-exclude-dir=.git"]
