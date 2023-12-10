# KubeBadges
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kubebadges/kubebadges)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/kubebadges/kubebadges?label=version)
![Docker Image Version (latest by date)](https://img.shields.io/docker/v/neosu/kubebadges?label=docker%20version)
![GitHub license](https://img.shields.io/github/license/kubebadges/kubebadges)

## What is KubeBadges
KubeBadges is an open-source Kubernetes status display tool designed to provide concise, intuitive status badges for services within Kubernetes. It offers instant status feedback while reducing external dependencies and complexity.

## Getting Started

### Install with Helm
```bash
helm repo add kubebadges https://kubebadges.github.io/kubebadges
helm install kubebadges kubebadges/kubebadges -n kubebadges --create-namespace
```

### Verify KubeBadges is running
```bash
kubectl get pods -n kubebadges
```

### Access KubeBadges Dashboard
```bash
kubectl port-forward svc/kubebadges 8090:8090 -n kubebadges
```
Open your browser and navigate to http://localhost:8090 to access the KubeBadges management dashboard.
Set up the necessary permissions in the KubeBadges dashboard to allow external access to the badges.

### Set Up External Access for Badges
KubeBadges dashboard runs on port 8090, while the external API uses port 8080. If you need to access badges from outside the cluster, you will need to configure Ingress or other means of exposure for KubeBadges' port 8080.
