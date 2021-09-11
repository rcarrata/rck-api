FROM golang:1.16 as build
WORKDIR /rck
COPY . /rck
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ./app .

FROM scratch
COPY --from=0 /rck/app .
EXPOSE 8080
CMD ["/app"]