FROM golang:alpine

# 作者
LABEL maintainer="fizzday <fizzday@yeah.net>" \
        Description="php7.3 alpine3.9"

# ########### 变量 ##############
ENV TZ Asia/Shanghai
ENV GOROSE_VERSION master
ENV GO111MODULE on

# ########## git ###########
RUN apk add --no-cache git

RUN mkdir -p /var/www && cd /var/www
COPY demo/* /var/www/

WORKDIR /var/www