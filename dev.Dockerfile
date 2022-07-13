# DO NOT use in production!
# Dockerfile for local development

FROM golang:1.18.4

# install the watcher
RUN go install github.com/githubnemo/CompileDaemon@v1.2.1

WORKDIR /project
COPY ./ /project

ENTRYPOINT ["CompileDaemon", "--build=go build -mod=mod", "-log-prefix=false", "-exclude-dir=.git"]
