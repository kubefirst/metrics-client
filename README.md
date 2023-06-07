# metrics-client

This client can be used to transmit metrics in isolation from other processes. This is useful for jobs, hooks, etc.

## Usage

```bash
transmit a metric

Usage:
  metrics-client transmit [flags]

Flags:
      --cluster-id string            the kubefirst cluster id (required)
      --cluster-type string          the kubefirst cluster type (required)
  -h, --help                         help for transmit
      --install-method string        the installation method for the cluster
      --kubefirst-team string        kubefirst team [true/false]
      --kubefirst-team-info string   kubefirst team info
      --type string                  the type of metric to transmit [cluster-zero] (required)
```

## Example

```bash
metrics-client transmit --type cluster-zero --cluster-id cluster --cluster-type mgmt --install-method helm
```
