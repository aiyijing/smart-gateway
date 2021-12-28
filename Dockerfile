FROM ubuntu:latest
RUN apt-get update && apt-get install -y iptables bash iproute2 && rm -rf /var/lib/apt/lists/*
ADD ./gw /gw
ADD ./client/dist /static
RUN mkdir /etc/gateway/
ENTRYPOINT ["/gw"]