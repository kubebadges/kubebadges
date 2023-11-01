FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y curl unzip zip git wget
RUN apt-get install -y xz-utils

RUN wget https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.13.9-stable.tar.xz
RUN tar xf flutter_linux_3.13.9-stable.tar.xz -C /home
ENV PATH="$PATH:/home/flutter/bin"
Run git config --global --add safe.directory /home/flutter