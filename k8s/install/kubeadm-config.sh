#!/bin/bash
cat >kubeadm-config.yaml <<EOF
apiVersion: kubeadm.k8s.io/v1beta2
bootstrapTokens:
- groups:
  - system:bootstrappers:kubeadm:default-node-token
  token: 9037x2.tcaqnpaqkra9vsbw
  ttl: 24h0m0s
  usages:
  - signing
  - authentication
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: 172.18.222.42
  bindPort: 6443
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: km1
  taints:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
---
apiServer:
  certSANs:  # 包含所有Master/LB/VIP IP，一个都不能少！为了方便后期扩容可以多写几个预留的IP
  - km1
  - km2
  - 172.18.222.42
  - 172.18.222.43
  - 172.18.222.44
  - 127.0.0.1
  extraArgs:
    authorization-mode: Node,RBAC
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta2
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controlPlaneEndpoint: 172.18.222.63:16443 # 负载均衡虚拟IP（VIP）和端口
controllerManager: {}
dns:
  type: CoreDNS
etcd:
  external:  # 使用外部etcd
    endpoints:
    - https://172.18.222.42:2379 # etcd集群3个节点
    - https://172.18.222.43:2379
    - https://172.18.222.44:2379
    caFile: /opt/etcd/ssl/ca.pem # 连接etcd所需证书
    certFile: /opt/etcd/ssl/server.pem
    keyFile: /opt/etcd/ssl/server-key.pem
imageRepository: registry.aliyuncs.com/google_containers # 由于默认拉取镜像地址k8s.gcr.io国内无法访问，这里指定阿里云镜像仓库地址
kind: ClusterConfiguration
kubernetesVersion: v1.20.0 # K8s版本，与上面安装的一致
networking:
  dnsDomain: cluster.local
  podSubnet: 10.244.0.0/16  # Pod网络，与下面部署的CNI网络组件yaml中保持一致
  serviceSubnet: 10.96.0.0/12  # 集群内部虚拟网络，Pod统一访问入口
scheduler: {}
EOF



#
kubeadm init --config kubeadm-config.yaml

kubeadm join 172.18.222.42:6443 --token 9037x2.tcaqnpaqkra9vsbw \
    --discovery-token-ca-cert-hash sha256:83b4d1a58286e712d9347c2e2c36f3abd187f086a1e4b4d3574d265ee17f0c54 --control-plane


 scp -r /etc/kubernetes/pki/ km2:/etc/kubernetes/