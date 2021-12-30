# SmartGateway

SmartGateway is smart gateway.It uses iptables to forward traffic. The traffic outbounds is tproxy of v2ray.
SmartGateway run in docker, So it is very easy to install and uninstall.


## Dependency

* Docker
* Linux

## QuickStart

* Easy to install

```bash
docker run -d --name gate --privileged --net host \
       -v "/var/run/docker.sock:/var/run/docker.sock" \
       -v "/var/lib/docker/volumes:/var/lib/docker/volumes" \
       aiyijing/gw:latest 
```

* uninstall

```bash
docker stop gate ;docker rm gate
# clear ip rule and iptables
ip route del local 0.0.0.0/0 dev lo table 100
ip rule del fwmark 0x1 table 100
iptables -F
iptables -X
iptables -Z
```

## Feature

* SmartDNS
* Smart forward traffic
* Simple to install and uninstall
