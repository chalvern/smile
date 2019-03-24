FROM golang:1.12

## vim
RUN apt-get update \
  && apt-get install -y vim \
  && rm -rf /var/lib/apt/lists/*

# vim setting
COPY vimrc /root/.vimrc

RUN go get -u github.com/derekparker/delve/cmd/dlv

WORKDIR $GOPATH
