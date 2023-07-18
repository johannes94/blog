# Create a Network Policy

Create a network policy in namespace `webserver` named `webserver-policy`.

The network policy should allow:

- TCP using port 80 from pods in the namespace `internal-client`
- TCP using port 80 from pods in the same namespace

The policy should be assigned to all pods in the namespace.

Furthermore, ensure that Pods in the namespace `webserver` can only connect to other pods in the same namespace using TCP on port 80.

You can use existing Pods in namespaces `webserver` `internal-client` and `external-client` to verify your policy.

