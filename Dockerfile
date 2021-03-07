FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/MSunFlower1014/golang-API
COPY . $GOPATH/src/github.com/MSunFlower1014/golang-API
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./golang-API"]