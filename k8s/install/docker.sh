#!/bin/bash
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum makecache fast
yum -y install docker-ce
sleep 2
#touch  /etc/docker/daemon.json
cat << EOF > /etc/docker/daemon.json
{
  "registry-mirrors": ["https://szko6ggt.mirror.aliyuncs.com"]
}
EOF
systemctl daemon-reload
systemctl enable    docker.service
systemctl restart docker



