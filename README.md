# KubeBadges

## What is KubeBadges
KubeBadges is an open-source Kubernetes status display tool designed to provide concise, intuitive status badges for services within Kubernetes. It offers instant status feedback while reducing external dependencies and complexity.

![screenshot](docs/images/index.png)

## Features

- **Status Badge Generation:** Easily generate status badges for various resources within Kubernetes, such as services, Pods, and Deployments.
- **Badge Customization:** Users can flexibly customize badges through KubeBadges' control panel, including attributes like color and name.
- **Custom Probes:** Users can extend the system's functionality by defining and configuring their probes to display specific service data.
- **Custom Service Names:** Provide easily understandable aliases or display names for services in Kubernetes.
- **Minimum Privilege Policy:** Use Kubernetes' RBAC to ensure that KubeBadges only accesses the data it needs with minimal permissions.

## Getting Started

## Advantages

- **Simplified Workflow:** KubeBadges reduces the complexity of viewing service status, providing a centralized, easy-to-manage solution.
- **No External Dependencies:** All data is stored within the Kubernetes cluster, eliminating the need for external databases or other services.
- **Secure and Reliable:** KubeBadges uses a minimum privilege policy to ensure data security.
- **Highly Customizable:** From the appearance of badges to custom probe data points, everything can be tailored to meet specific needs.

## License
[Apache License Version 2.0](./LICENSE)
