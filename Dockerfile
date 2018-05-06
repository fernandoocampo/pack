FROM iron/base

EXPOSE 8287
ADD pack-service-linux-amd64 /
ENTRYPOINT ["./pack-service-linux-amd64"]