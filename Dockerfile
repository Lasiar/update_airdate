FROM amd64/golang as builder
RUN mkdir -p /src/kre_air_update
RUN GOPATH=/ go get github.com/denisenkom/go-mssqldb
ADD ./main.go /src/kre_air_update
ADD ./sys /src/kre_air_update/sys
ADD ./model /src/kre_air_update/model
ADD ./web /src/kre_air_update/web
RUN cd /src/kre_air_update
WORKDIR /src/kre_air_update
RUN  GOPATH=/ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

FROM amd64/alpine
RUN mkdir /app
COPY --from=builder /src/kre_air_update/main /app
ADD prod.config.json /app/config.json
ADD assets /app/assets
WORKDIR /app
EXPOSE 80
CMD ["/app/main"]

#save -o