# BUILD_TYPE can have these values: local, test or prod
# if BUILD_TYPE is empty, set to local
ARG BUILD_TYPE=local

FROM golang:stretch as build-env
COPY . ./src
RUN apt update
WORKDIR /go/src
RUN echo "Setup build environnement"
RUN export PATH=$PATH:/go/bin
RUN export GO111MODULE=on
RUN echo "Building Microsevice..."
RUN go build cmd/microservice-achievements/main.go
RUN echo "First Docker build-stage is now done"

FROM gcr.io/distroless/base as prod

FROM golang:stretch as test

FROM golang:stretch as local

FROM ${BUILD_TYPE} AS exit_artefact
COPY --from=build-env /go/src/main /microservice
ENV JAEGER_ENDPOINT=http://jaeger-tracing-collector:14268/api/traces
EXPOSE 8888
EXPOSE 9090
CMD ["/microservice"]
