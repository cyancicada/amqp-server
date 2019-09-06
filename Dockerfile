FROM scratch
ADD amqpserver /amqpserver
RUN chmod 755 /*
CMD ["/amqpserver","-c","config.json"]