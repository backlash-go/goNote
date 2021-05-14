#!/bin/bash

wget https://github.com/etcd-io/etcd/releases/download/v3.4.9/etcd-v3.4.9-linux-amd64.tar.gz

mkdir /opt/etcd/{bin,cfg,ssl} -p
tar zxvf etcd-v3.4.9-linux-amd64.tar.gz
mv etcd-v3.4.9-linux-amd64/{etcd,etcdctl} /opt/etcd/bin/

cat >/opt/etcd/cfg/etcd.conf <<EOF
#[Member]
ETCD_NAME="etcd-1"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://172.18.222.42:2380"
ETCD_LISTEN_CLIENT_URLS="https://172.18.222.42:2379"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://172.18.222.42:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://172.18.222.42:2379"
ETCD_INITIAL_CLUSTER="etcd-1=https://172.18.222.42:2380,etcd-2=https://172.18.222.43:2380,etcd-3=https://172.18.222.44:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF



cat > /opt/etcd/cfg/etcd.conf << EOF
#[Member]
ETCD_NAME="etcd-2"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://172.18.222.43:2380"
ETCD_LISTEN_CLIENT_URLS="https://172.18.222.43:2379"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://172.18.222.43:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://172.18.222.43:2379"
ETCD_INITIAL_CLUSTER="etcd-1=https://172.18.222.42:2380,etcd-2=https://172.18.222.43:2380,etcd-3=https://172.18.222.44:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF

cat > /opt/etcd/cfg/etcd.conf << EOF
#[Member]
ETCD_NAME="etcd-3"
ETCD_DATA_DIR="/var/lib/etcd/default.etcd"
ETCD_LISTEN_PEER_URLS="https://172.18.222.44:2380"
ETCD_LISTEN_CLIENT_URLS="https://172.18.222.44:2379"

#[Clustering]
ETCD_INITIAL_ADVERTISE_PEER_URLS="https://172.18.222.44:2380"
ETCD_ADVERTISE_CLIENT_URLS="https://172.18.222.44:2379"
ETCD_INITIAL_CLUSTER="etcd-1=https://172.18.222.42:2380,etcd-2=https://172.18.222.43:2380,etcd-3=https://172.18.222.44:2380"
ETCD_INITIAL_CLUSTER_TOKEN="etcd-cluster"
ETCD_INITIAL_CLUSTER_STATE="new"
EOF

#•ETCD_NAME：节点名称，集群中唯一
#•ETCDDATADIR：数据目录
#•ETCDLISTENPEER_URLS：集群通信监听地址
#•ETCDLISTENCLIENT_URLS：客户端访问监听地址
#•ETCDINITIALADVERTISEPEERURLS：集群通告地址
#•ETCDADVERTISECLIENT_URLS：客户端通告地址
#•ETCDINITIALCLUSTER：集群节点地址
#•ETCDINITIALCLUSTER_TOKEN：集群Token
#•ETCDINITIALCLUSTER_STATE：加入集群的当前状态，new是新集群，existing表示加入已有集群




#3. systemd管理etcd
cat >/usr/lib/systemd/system/etcd.service <<EOF
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target

[Service]
Type=notify
EnvironmentFile=/opt/etcd/cfg/etcd.conf
ExecStart=/opt/etcd/bin/etcd \
--cert-file=/opt/etcd/ssl/server.pem \
--key-file=/opt/etcd/ssl/server-key.pem \
--peer-cert-file=/opt/etcd/ssl/server.pem \
--peer-key-file=/opt/etcd/ssl/server-key.pem \
--trusted-ca-file=/opt/etcd/ssl/ca.pem \
--peer-trusted-ca-file=/opt/etcd/ssl/ca.pem \
--logger=zap
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

# 节点启动顺序 先启动 43 -> 44 ->42
# 报错 etcd 没有按照顺序启动报错原因 https://www.jianshu.com/p/40853445b042
systemctl daemon-reload
systemctl start etcd
systemctl enable etcd
