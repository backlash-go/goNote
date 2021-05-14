#!/bin/bash

#1 install csffl tool
curl -L https://github.com/cloudflare/cfssl/releases/download/v1.5.0/cfssl_1.5.0_linux_amd64 -o cfssl
chmod +x cfssl
curl -L https://github.com/cloudflare/cfssl/releases/download/v1.5.0/cfssljson_1.5.0_linux_amd64 -o cfssljson
chmod +x cfssljson
curl -L https://github.com/cloudflare/cfssl/releases/download/v1.5.0/cfssl-certinfo_1.5.0_linux_amd64 -o cfssl-certinfo
chmod +x cfssl-certinfo

mv cfss* /usr/local/bin
mkdir cert
cd cert



cat >ca-config.json <<EOF
{
  "signing": {
    "default": {
      "expiry": "87600h"
    },
    "profiles": {
      "www": {
         "expiry": "87600h",
         "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ]
      }
    }
  }
}
EOF

cat >ca-csr.json <<EOF
{
    "CN": "etcd CA",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "Beijing",
            "ST": "Beijing"
        }
    ]
}
EOF

# 生成ca.pem和ca-key.pem文件。
cfssl gencert -initca ca-csr.json | cfssljson -bare ca

# hosts：etcd cluster node IP
cat >server-csr.json <<EOF
{
    "CN": "etcd",
    "hosts": [
    "172.18.222.42",
    "172.18.222.43",
    "172.18.222.44"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "BeiJing",
            "ST": "BeiJing"
        }
    ]
}
EOF

# -profile   ca-config.json 文件里面的www
cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \
  --config=ca-config.json -profile=www \
  server-csr.json | cfssljson -bare server
