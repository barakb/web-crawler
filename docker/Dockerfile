FROM golang:1.7.1
MAINTAINER Barak Bar Orion  <barak.bar@gmail.com>

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

ENV app fake-web-tree

RUN apt-get install -y git

RUN curl https://glide.sh/get | sh

RUN mkdir -p /golang/$app/src/github.com/barakb
WORKDIR /golang/$app/src/github.com/barakb

RUN git clone https://github.com/barakb/$app.git


WORKDIR /golang/fake-web-tree/src/github.com/barakb/$app

env PATH /golang/$app/src/github.com/barakb/$app:$PATH

RUN glide install

WORKDIR /golang/$app

COPY serve.sh /golang/$app/

ENTRYPOINT ["/golang/$app/serve.sh"]
