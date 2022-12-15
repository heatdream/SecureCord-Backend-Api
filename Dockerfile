FROM golang:1.19-alpine3.16 as api_builder

WORKDIR /go/src/github.com/heatdream/securecord-backend-api
ADD . /go/src/github.com/heatdream/securecord-backend-api/
RUN go mod download && go mod verify
RUN go build -o api-exec .


FROM alpine:latest
RUN mkdir application
COPY --from=api_builder /go/src/github.com/heatdream/securecord-backend-api/api-exec /application/
WORKDIR /application
CMD ["/application/api-exec"]