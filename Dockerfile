FROM m.daocloud.io/docker.io/ubuntu:22.04

WORKDIR /app

COPY vorker /app/

EXPOSE 8888
EXPOSE 8080

CMD [ "/app/vorker" ]