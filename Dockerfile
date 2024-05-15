FROM ubuntu:latest
LABEL authors="neko"

ENTRYPOINT ["top", "-b"]