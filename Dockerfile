FROM scratch
ADD amqpserver /amqpserver
ADD config.json /amqpserver
VOLUME /amqpserver
CMD ["/amqpserver","-c","config.json"]