FROM alpine
VOLUME /tmp/apps/logs
VOLUME /tmp/apps/conf
COPY ./amqpserver /tmp/apps/amqpserver
COPY ./config.json /tmp/apps/conf/config.json
WORKDIR /tmp/apps
RUN chmod +x amqpserver
CMD ["./amqpserver","-c","/tmp/apps/conf/config.json"]
