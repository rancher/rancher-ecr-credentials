FROM ubuntu:16.04
RUN apt-get update && apt-get install -y ca-certificates
COPY rancher-ecr-credentials /usr/bin/
CMD ["rancher-ecr-credentials"]
