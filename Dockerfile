FROM golang:alpine
LABEL author="wwhai"
LABEL email="cnwwhai@gmail.com"
LABEL homepage="rulex.ezlinker.cn"
ENV APP_NAME="rulex"
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPROXY="https://goproxy.cn,direct"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add build-base

ADD . /rulex/
WORKDIR /rulex/
RUN make

EXPOSE 2580
EXPOSE 2581
EXPOSE 2582
EXPOSE 2583
EXPOSE 2584
EXPOSE 2585

ENTRYPOINT ["./rulex", "run"]

