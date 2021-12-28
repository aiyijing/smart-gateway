#!/bin/bash
set -x

docker ps -a | grep gate && docker rm gate
ip route del local 0.0.0.0/0 dev lo table 100
ip rule del fwmark 0x1 table 100
iptables -t mangle -F
iptables -t mangle -X
iptables -t mangle -Z