FROM alpine
VOLUME /tmp/apps/logs
VOLUME /tmp/apps/conf
COPY ./main /tmp/apps/main
COPY ./config.json /tmp/apps/conf/config.json
WORKDIR /tmp/apps
RUN chmod +x main
CMD ["./amqpserver","-c","/tmp/apps/conf/config.json"]
