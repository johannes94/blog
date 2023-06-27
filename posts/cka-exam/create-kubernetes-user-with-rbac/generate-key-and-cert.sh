set -e
# this generates a private key for the new user we want to create
openssl genrsa -out moon.key 2048

# with the key we can generate a certificate signing requests
openssl req -new -key moon.key -out moon.csr --subj "/CN=moon"

# The certificate expects the spec.requests field to be set with the content
# of moon.csr encoded using base64. Use following command to do that
request=$(cat moon.csr | base64 -w0)

# After that create the CSR resource
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: moon
spec:
  request: $request
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400  # one day
  usages:
  - client auth
EOF

kubectl certificate approve moon
echo "Waiting 10 sec for approval"
kubectl get csr moon -o jsonpath='{.status.certificate}'| base64 -d > moon.crt
