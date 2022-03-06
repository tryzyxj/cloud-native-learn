FROM golang:1.16-alpine AS build
COPY ./ /go/src/project/
WORKDIR /go/src/project/
RUN go build -o /bin/project

FROM alpine:3.15
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]