FROM golang:1.8

RUN go get github.com/kardianos/govendor #eedb81532d50829397bc91628e48c48438433e90

ENV PROJECTPATH=/go/src/github.com/replicatedcom/loadis

ENV LOG_LEVEL DEBUG

WORKDIR $PROJECTPATH

CMD ["/bin/bash"]
