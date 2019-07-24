FROM golang:alpine

# 作者
LABEL maintainer="fizzday <fizzday@yeah.net>" \
        Description="golang1.12 gorose2.0 alpine"

# ########### 变量 ##############
ENV TZ Asia/Shanghai
ENV GOROSE_VERSION master
ENV GO111MODULE on

# ########## git ###########
RUN apk add --no-cache git

RUN mkdir -p /var/www
COPY demo/* /var/www/

RUN cd /var/www && go mod tidy

WORKDIR /var/www