FROM m.daocloud.io/docker.io/ubuntu:22.04

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
	sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -qy libc++1 ca-certificates

WORKDIR /app

COPY vorker /app/

EXPOSE 8888
EXPOSE 8080

CMD [ "/app/vorker" ]