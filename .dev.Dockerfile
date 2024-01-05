FROM ubuntu:22.04

USER root

LABEL maintainer me@vaala.cat

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
	sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -qy libc++1

RUN apt-get update && DEBIAN_FRONTEND="noninteractive" apt-get install -y\
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg \
	zsh \
	fish \
	lsb-release \
	wget \
	tmux git \
	build-essential \
	sudo \
	rsync \
	ssh \
	vim \
	unzip \
	p7zip-full \
	nodejs \
	npm \
	bash \
	inetutils-ping \
	net-tools \
	pgcli \
	htop \
	locales \
	man \
	python3 \
	python3-pip \
	software-properties-common \
	systemd \
	systemd-sysv \
	fuse3 \
	sqlite \
	--no-install-recommends 

RUN wget https://mirrors.ustc.edu.cn/golang/go1.21.1.linux-arm64.tar.gz && \
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.1.linux-arm64.tar.gz

WORKDIR /app

RUN useradd dev \
	--create-home \
	--shell=/usr/bin/fish \
	--uid=1000 \
	--user-group && \
	echo "dev ALL=(ALL) NOPASSWD:ALL" >>/etc/sudoers.d/nopasswd

USER dev

RUN pip config set global.index-url http://pypi.douban.com/simple/ && \
	pip config set install.trusted-host pypi.douban.com && \
	npm config set registry https://registry.npm.taobao.org/

COPY --from=flyio/litefs:0.5 /usr/local/bin/litefs /usr/local/bin/litefs

ENV GOPROXY https://proxy.golang.com.cn,direct
ENV PATH /usr/local/go/bin:$PATH
ENV GOROOT /usr/local/go

RUN go install github.com/cweill/gotests/gotests@latest 		&& \
	go install github.com/fatih/gomodifytags@latest     		&& \
	go install github.com/josharian/impl@latest             	&& \
	go install github.com/haya14busa/goplay/cmd/goplay@latest 	&& \
	go install github.com/go-delve/delve/cmd/dlv@latest     	&& \
	go install honnef.co/go/tools/cmd/staticcheck@latest    	&& \
	go install golang.org/x/tools/gopls@latest

EXPOSE 8888
EXPOSE 8080

CMD [ "/bin/bash" ]