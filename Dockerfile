FROM alpine

WORKDIR /app

COPY vorker /app/

EXPOSE 8888
EXPOSE 8080

CMD [ "vorker" ]