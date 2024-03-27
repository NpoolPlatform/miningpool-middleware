
FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY miningpool-middleware /app

ENTRYPOINT ["/app/miningpool-middleware"]

EXPOSE 50530 50531 50532 50533

CMD ["run"]
