FROM ababy/golang-dependency

# ############## author ##############
LABEL maintainer="fizzday <fizzday@yeah.net>" \
        Description="golang gorose alpine | latest"

# ############## env ##############
ENV TZ Asia/Shanghai
ENV GOROSE_VERSION master
ENV GO111MODULE on
ENV GOPRODIR /var/www/

# ############## work dir ###########
RUN mkdir -p $GOPRODIR
COPY demo/* $GOPRODIR

# ############## gorose init & run for once ###########
RUN cd $GOPRODIR && go mod tidy && go run main.go

WORKDIR $GOPRODIR