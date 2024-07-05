# README

## Overview

Instructions for deploying an Ethereum Mainnet Full Node using Geth and Prysm, monitored using the Prometheus stack. 

This leverages [ethereum-helm-charts](https://github.com/ethpandaops/ethereum-helm-charts), which is actively maintained and provides support for many clients and configurations.

Additionally, this project includes functionality to create custom metrics or other data formats that can be subscribed to over WebSockets.

### Prerequisites

- Kubernetes cluster up and running
- Helm and kubectl are installed
- An existing storage class with ample disk capacity
- Worker nodes should be labeled with `role=worker` or change the affinity rules as desired
- ghcr creds stored as a secret on the namespace being used

### Features

- Custom exporter written in Go to supply metrics and/or application data from Geth. This can connect via IPC or HTTP. IPC is preferred to avoid enabling the Admin Api over http/ws.
- Real-time feed of custom data via WebSockets, currently showing a list of connected peers, accessible via the `/ui` route on the custom-metrics port (3737).
- Custom peer count metric with labels for additional data such as enode and IP address.
- Ethereum exporter with Grafana dashboard. The dashboard is stored as a ConfigMap and is automatically imported using the grafana dashboard provider.
- Headless and ClusterIP services
- Expandable to include many other clients, configurations, and custom metrics or applications
- Path based storage to enable auto pruning (>= v1.14.0)

## Instructions

```sh
# Update Storage Class
# Before installing the Helm charts, update the storage class to match your environment.
# Edit the values.yaml files for both charts to specify the correct storageClassName.

# Example:
# persistence:
#   storageClassName: "your-storage-class"


# Prometheus Chart
cd charts/prometheus
helm dependency update
helm install my-prometheus-release . -f values.yaml

# Ethereum Chart
cd charts/ethereum
helm dependency update
helm install my-ethereum-release . -f mainnet-prysm.yaml


# Configuration
Defaults can be found here
https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/ethereum-node/values.yaml
https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/geth/values.yaml
https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/prysm/values.yaml
https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml
https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml

# Access
Port forward RPC, Grafana, Prometheus, Custom UI, Custom metrics, or access via Twingate using the cluster domain
The Grafana default credentails are  admin:admin, you will be prompted to change on first login



