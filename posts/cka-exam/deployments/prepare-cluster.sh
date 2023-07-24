#!/bin/bash

# Prepare for Q1: Create a Deployment
kubectl create ns moon

# Prepare for Q2: Scale a Deployment
kubectl create ns mars
kubectl create deploy mars-server -n mars --image nginx:1.25.1 --replicas 3 

# Prepare for Q3: Set image of a Deployment
kubectl create ns saturn
kubectl create deploy saturn-server -n saturn --image nginx:1.24.0 --replicas 2

# Prepare for Q4: Mount secret in Deployment
kubectl create ns uranus
kubectl create secret generic db-user -n uranus --from-literal=user=testuser --from-literal=pass=1234

# Prepare for Q5: Environment configuration for a Deployment
kubectl create ns sun
kubectl create configmap app-config -n sun --from-literal=name=sun-server

# Prepare for Q6: Rollout and Rollback a Deployment
kubectl create ns earth
kubectl create deploy earth-server -n earth --image nginx:1.25.1 --replicas 5