#!/bin/bash

kubectl create ns webserver
kubectl create ns internal-client
kubectl create ns external-client

kubectl run -n webserver webserver1 --image nginx
kubectl run -n webserver webserver2 --image nginx

kubectl run -n internal-client client1 --image busybox -- sh -c "while true; do echo Running; sleep 1; done;"
kubectl run -n external-client client2 --image busybox -- sh -c "while true; do echo Running; sleep 1; done;"
