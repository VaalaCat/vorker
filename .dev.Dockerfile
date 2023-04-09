FROM ubuntu:22.10

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
	sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

RUN apt update && apt install -qy tini libc++1 zsh wget curl git vim htop tmux p7zip-full

RUN wget https://mirrors.ustc.edu.cn/golang/go1.19.4.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz

ENV GOPROXY https://proxy.golang.com.cn,direct
ENV GOPRIVATE git.vaala.cloud
RUN git config --global url."ssh://git@git.vaala.cloud".insteadOf "https://git.vaala.cloud"