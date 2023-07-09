#! /bin/bash
curl -LO https://github.com/etcd-io/etcd/releases/download/v3.5.6/etcd-v3.5.6-linux-amd64.tar.gz 
tar xvzf etcd-v3.5.6-linux-amd64.tar.gz 
export PATH=$PATH:$(pwd)/etcd-v3.5.6-linux-amd64

export cert_path=/var/lib/minikube/certs/etcd

export ETCDCTL_API=3 
export ETCDCTL_CACERT=$cert_path/ca.crt
export ETCDCTL_CERT=$cert_path/server.crt
export ETCDCTL_KEY=$cert_path/server.key

etcdctl endpoint health