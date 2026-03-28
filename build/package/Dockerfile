FROM golang:1.24

ARG UID=1000
ARG USERNAME=docker

RUN apt-get update && \
    apt-get -y install sudo && \
    useradd -u $UID -G sudo -m $USERNAME && \
    echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers

USER $USERNAME

RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8
