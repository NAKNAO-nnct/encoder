FROM golang:1.19-alpine

RUN apk update && apk --no-cache add ca-certificates curl bash xz-libs git wget python3 py3-pip
RUN pip3 install awscli
WORKDIR /tmp
RUN wget https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz
RUN tar -xf ffmpeg-git-amd64-static.tar.xz && cd ff* && mv ff* /usr/local/bin

WORKDIR /app

ENTRYPOINT ["/bin/bash"]
