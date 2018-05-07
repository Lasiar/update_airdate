FROM amd64/golang as builder
MAINTAINER Никониров Григорий mrgbh007@gmail.com
RUN go get github.com/denisenkom/go-mssqldb
RUN mkdir /src
COPY *.go /src/
RUN cd /src
WORKDIR /src
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main

FROM amd64/alpine
MAINTAINER Никониров Григорий mrgbh007@gmail.com
RUN mkdir /app
COPY --from=builder /src/main /app
ADD template.xlsx /app
ADD config.json /app
ADD template.html /app
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
