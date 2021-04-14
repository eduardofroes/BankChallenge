FROM golang:1.15-alpine

RUN mkdir -p /opt/bankchallenge
WORKDIR /opt/bankchallenge

COPY src/bankchallenge /opt/bankchallenge

EXPOSE 8080

RUN chmod -R 777 /opt/bankchallenge
CMD ["/opt/bankchallenge/bankchallenge"]
