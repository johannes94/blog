# stop control plane pods so they don't write to etcd
# in the shell connected to your control plane
mv /etc/kubernetes/manifests/kube-apiserver.yaml /etc/kubernetes/
mv /etc/kubernetes/manifests/etcd.yaml /etc/kubernetes/

# get the data directory of etcd: --data-dir=/var/lib/minikube/etcd
cat /etc/kubernetes/etcd.yaml | grep data-dir
# clear the data dir from old data
rm -r /var/lib/minikube/etcd/*
etcdctl snapshot restore snapshot.db --data-dir=/var/lib/minikube/etcd

# restart control plane pods
mv /etc/kubernetes/etcd.yaml /etc/kubernetes/manifests/
mv /etc/kubernetes/kube-apiserver.yaml /etc/kubernetes/manifests/