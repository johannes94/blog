#!/bin/bash
kubectl create ns backup-ns
kubectl create deploy test-app -n backup-ns --image nginx:1.25.1
kubectl create cm test-cm1 -n backup-ns
kubectl create cm test-cm2 -n backup-ns
kubectl create secret generic test-secret1 -n backup-ns
kubectl create secret generic test-secret2 -n backup-ns
