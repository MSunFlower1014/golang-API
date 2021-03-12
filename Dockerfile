# 此docker 通过dlv启动，用来进行远程代码调试专用
FROM golang:1.14.3-alpine3.11

ENV CGO_ENABLED 0
# 设置代理
ENV GOPROXY https://goproxy.cn,direct

# Allow Go to retreive the dependencies for the build step
RUN apk add  git
# 此步骤参考官网安装说明：https://github.com/go-delve/delve/blob/master/Documentation/installation/README.md
WORKDIR /
RUN git clone  https://github.com.cnpmjs.org/go-delve/delve.git
WORKDIR /delve
RUN go install github.com/go-delve/delve/cmd/dlv

# 复制web执行文件和配置文件，打包程序环境 GOOS=linux
COPY ./app /app
COPY ./config  /config
WORKDIR /

EXPOSE 8000 40000

#CMD ["dlv", "--continue", "--accept-multiclient","--listen=:40000", "--headless=true", "--api-version=2","exec",  "/app"]
CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2","exec",  "/app"]