# README

## Overview

Instructions for deploying an Ethereum Mainnet Full Node using Geth and Prysm, monitored using the Prometheus stack.

This leverages [ethereum-helm-charts](https://github.com/ethpandaops/ethereum-helm-charts), which is actively maintained and provides support for many clients and configurations.

Additionally, this project includes functionality to create custom metrics or other data formats that can be subscribed to over WebSockets.

### Prerequisites

- Kubernetes cluster up and running
- Helm and kubectl are installed
- An existing storage class with ample disk capacity
- A values file to match your environment
- ghcr credentials stored as a secret in the namespace being used

### Features

- Custom exporter written in Go to supply metrics and application data from Geth. This can connect via IPC or HTTP. IPC is preferred to avoid enabling the Admin API over http/ws.
- Real-time feed of custom data via WebSockets, currently showing a list of connected peers, accessible via the `/ui` route on the custom-metrics port (3737).
- Custom peer count metric with labels for additional data such as enode and IP address.
- Ethereum exporter with a Grafana dashboard. The dashboard is stored as a ConfigMap and is automatically imported using the Grafana dashboard provider.
- Headless and ClusterIP services
- Expandable to include many other clients, configurations, and custom metrics or applications

This project is currently running on RKE2 v1.30.1+rke2r1 utilizing vSphere CPI and CSI. These charts are currently managed by RKE2.

- [Rancher vSphere CSI](https://github.com/rancher/rke2-charts/blob/main/charts/rancher-vsphere-csi/rancher-vsphere-csi/3.1.2-rancher400/values.yaml)
- [Rancher vSphere CPI](https://github.com/rancher/rke2-charts/blob/main/charts/rancher-vsphere-cpi/rancher-vsphere-cpi/1.7.001/values.yaml)

In addition to container volume provisioning, this also enables features such as volume snapshots and restore. This can be further expanded to use features such as vSAN enabling a scalable bare metal architecture with high-performance storage.

### Configuration

Defaults can be found here:
- [Ethereum Node](https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/ethereum-node/values.yaml)
- [Geth](https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/geth/values.yaml)
- [Prysm](https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/prysm/values.yaml)
- [Ethereum Metrics Exporter](https://github.com/ethpandaops/ethereum-helm-charts/blob/master/charts/ethereum-metrics-exporter/values.yaml)
- [Kube Prometheus Stack](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml)
- [Grafana](https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml)

### Access

Port forward RPC, Grafana, Prometheus, Custom UI, Custom metrics, or access via Twingate using the cluster domain.
The Grafana default credentials are admin:admin; you will be prompted to change them on the first login.

## Instructions

```sh
# Prometheus Chart
cd charts/prometheus
helm dependency update
helm install my-prometheus-release . -f my_values.yaml

# Ethereum Chart
cd charts/ethereum
helm dependency update
helm install my-ethereum-release . -f my_values.yaml
