FROM scratch
ADD amqpserver /amqpserver
ADD config.json /amqpserver
VOLUME /amqpserver/config.json
VOLUME /amqpserver/logs
CMD ["/amqpserver","-c","config.json"]