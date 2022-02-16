FROM golang

EXPOSE 8080

RUN apt-get --yes update
RUN apt-get --yes install gcc libc6-dev \
    libglu1-mesa-dev libgl1-mesa-dev \
    libxcursor-dev libxi-dev libxinerama-dev \
    libxrandr-dev libxxf86vm-dev \
    libasound2-dev pkg-config
RUN go install github.com/hajimehoshi/wasmserve@latest

WORKDIR /app

ADD . /app

RUN go get .


CMD [ "/bin/bash", "-c", "echo http://localhost:8080/ && wasmserve ./" ]
