FROM golang:alpine

# ############## author ##############
LABEL maintainer="fizzday <fizzday@yeah.net>" \
        Description="golang gorose alpine | latest"

# ############## env ##############
ENV TZ Asia/Shanghai
ENV GOROSE_VERSION master
ENV GO111MODULE on

# ############## golang dependency ###########
RUN apk add gcc
RUN apk add libc-dev
RUN apk add --no-cache git

# ############## work dir ###########
RUN mkdir -p /var/www
COPY demo/* /var/www/

# ############## gorose init & run for once ###########
RUN cd /var/www && go mod tidy && go run main.go

WORKDIR /var/www