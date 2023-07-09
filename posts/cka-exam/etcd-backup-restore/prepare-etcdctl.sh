#! /bin/bash
curl -LO https://github.com/etcd-io/etcd/releases/download/v3.5.6/etcd-v3.5.6-linux-amd64.tar.gz 
tar xvzf etcd-v3.5.6-linux-amd64.tar.gz 
export PATH=$PATH:$(pwd)/etcd-v3.5.6-linux-amd64

export cert_path=/var/lib/minikube/certs/etcd

export ETCDCTL_API=3 
export ETCDCTL_CACERT=$cert_path/ca.crt
export ETCDCTL_CERT=$cert_path/server.crt
export ETCDCTL_KEY=$cert_path/server.key

kubectl -n kube-system exec -it etcd-minikube -- sh -c "ETCDCTL_API=3 ETCDCTL_CACERT=$cert_path/ca.crt ETCDCTL_CERT=$cert_path/server.crt ETCDCTL_KEY=$cert_path/server.key etcdctl endpoint health" 
127.0.0.1:2379 is healthy: successfully committed proposal: took = 12.3369ms

kubectl -n kube-system exec etcd-minikube -- sh -c "ETCDCTL_API=3 \ 
ETCDCTL_CACERT=$cert_path/ca.crt ETCDCTL_CERT=$cert_path/server.crt ETCDCTL_KEY=$cert_path/server.key \ 
etcdctl member list" 
aec36adc501070cc, started, minikube, https://192.168.49.2:2380, https://192.168.49.2:2379, false

ubectl -n kube-system exec 

sh -c "ETCDCTL_API=3 ETCDCTL_CACET=$cert_path/ca.crt ETCDCTL_CERT=$cert_path/server.crt ETCDCTL_KEY=$cert_path/server.key etcdctl snapshot restore /tmp/snapshot.db --data-dir /var/lib/minikube/etcd"