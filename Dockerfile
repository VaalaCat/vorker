FROM ubuntu:22.04

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
	sed -i s@/security.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
	sed -i 's/ports.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

COPY litefs.yml /etc/litefs.yml
COPY --from=flyio/litefs:0.5 /usr/local/bin/litefs /usr/local/bin/litefs

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -qy libc++1 ca-certificates

COPY bin/* /bin/

RUN chmod +x /bin/*

WORKDIR /app

COPY vorker /app/

EXPOSE 8888
EXPOSE 8080

CMD [ "/app/vorker" ]