# miniDevPod

A lightweight CLI tool for managing development pods in Kubernetes. Built with Go and the Cobra CLI framework.

## Features

- **Create** dev pods with custom resource allocation
- **Connect** to existing pods via SSH
- **List** all managed pods
- **Delete** pods when no longer needed
- **Forward** ports between host and pod
- **Sync** files from host to pod

## Demo

[Screencast from 2026-02-07 15-29-51.webm](https://github.com/user-attachments/assets/890e5b05-6cec-40b9-9715-c2cbdd23c390)


## Installation

```bash
go build -o minidevpod main.go
```

## Usage

### Create a new dev pod
```bash
./minidevpod create --name my-pod --repo https://github.com/user/repo.git --branch main --cpu 500m --memory 1Gi
```

### Connect to an existing pod
```bash
./minidevpod connect my-pod
# or use the alias
./minidevpod ssh my-pod
```

### List all pods
```bash
./minidevpod list
# or use the alias
./minidevpod ls
```

### Delete a pod
```bash
./minidevpod delete my-pod
# or use the alias
./minidevpod rm my-pod
```

### Forward ports
```bash
./minidevpod forward my-pod 8080:8080
```

### Sync files
```bash
./minidevpod sync my-pod /path/to/local/files
```

## Requirements

- Go 1.25.3 or higher
- Access to a Kubernetes cluster
- kubectl configured

## License

See LICENSE file for details.
