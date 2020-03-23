FROM golang:alpine

ADD . $GOPATH/src/spiros
WORKDIR $GOPATH/src/spiros

RUN apk add --update --no-cache git gcc libc-dev tzdata;
#  tzdata wget gcc libc-dev make openssl py-pip;

ENV TZ=${SPIROS_TIMEZONE}

# CMD if [ "${APPENV}" = "staging" ] || [ "${APPENV}" = "production" ] ; then \
#         openssl aes-256-cbc -d -in deploy/conf.enc -out config.yaml -pbkdf2 -pass file:./public.pem ; \
#     elif [ "${APPENV}" = "dev" ] ; then \
#         cp deploy/dev-config.yaml config.yaml ; \
#     fi \
CMD go build -v -o $GOPATH/bin/spiros \
    && spiros run server;
EXPOSE ${SPIROS_PORT}