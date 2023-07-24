kubectl set image deploy -n earth earth-server nginx=nginx:1.25.3000

# Write deployment status to file
kubectl rollout status deploy -n earth earth-server -w=false > /tmp/earth-status.txt

# Rollback the deployment
kubectl rollout undo deployment -n earth earth-server

# Get Pods
kubectl get pods -n earth

# Verify image
kubectl get deploy -n earth -o yaml | grep image:

# Write deployments rollout history to file 
kubectl rollout history deploy -n earth earth-server > /tmp/earth-history.txt