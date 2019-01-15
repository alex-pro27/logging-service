FROM golang:1.11.4
COPY root /
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
#RUN yum -y install epel-release \
#    && yum -y update \
#    && yum -y install golang \
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && chmod +x /entrypoint.sh \
    && go get github.com/derekparker/delve/cmd/dlv \
    && go get github.com/kshvakov/clickhouse \
    && go get github.com/gin-gonic/gin \
    && go get github.com/jmoiron/sqlx \
    && go get github.com/Masterminds/squirrel
CMD ["/entrypoint.sh"]
