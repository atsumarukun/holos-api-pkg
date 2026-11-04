FROM golang:1.27.3@sha256:2920bb124f1a2e12a1d39f0cab756ae0068441ee1da7655dec7f5e374944095b

ARG UID=1000 USERNAME=developer
RUN useradd -u $UID -m $USERNAME
USER $USERNAME

RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0
