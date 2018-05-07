FROM amd64/golang as builder
RUN go get github.com/denisenkom/go-mssqldb
RUN mkdir /kre_air_update
COPY ./ /kre_air_update
RUN cd /   kre_air_update
WORKDIR /kre_air_update
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

FROM amd64/alpine
RUN mkdir /app
COPY --from=builder /air/main /app
ADD config.json /app
ADD assets /app
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
