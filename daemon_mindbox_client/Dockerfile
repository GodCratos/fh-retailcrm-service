## Latest version
#ARG GO_VERSION=1.15
#
#FROM golang:${GO_VERSION}-alpine
#
#WORKDIR /src
#
#COPY ./go.mod ./go.sum ./
#RUN go mod download
#
#COPY ./ ./
#
#RUN CGO_ENABLED=0 go build \
#   -o /main main.go
#
#WORKDIR /srv
#COPY /main /srv/main
#
#EXPOSE 80
#
## Run the compiled binary.
#ENTRYPOINT ["/srv/main"]

FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]