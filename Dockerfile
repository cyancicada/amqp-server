FROM scratch
ADD amqpserver /amqpserver
CMD ["/amqpserver"]