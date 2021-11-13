FROM golang:1.17.3-alpine AS build

COPY serve.go ./
RUN go env -w GO111MODULE=off && CGO_ENABLED=0 go build -o /bin/serv

FROM scratch
COPY --from=build /bin/serv /bin/serv

EXPOSE 8080

CMD ["/bin/serv"]