FROM scratch
ADD amqpserver /amqpserver
VOLUME /amqpserver/config.json
VOLUME /amqpserver/logs
CMD ["/amqpserver","-c","config.json"]