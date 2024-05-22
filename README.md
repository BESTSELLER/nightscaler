# Nightscaler

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bestseller/nightscaler?sort=semver&style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bestseller/nightscaler?style=flat-square)

Nightscaler is an application designed to annotate Kubernetes namespaces based
on Pub/Sub messages. It serves as an extension to kube-downscaler, providing
additional functionality for managing forced-uptime periods within namespaces.

## Overview

Kubernetes clusters often utilize kube-downscaler to scale down namespaces and
conserve resources. However, this scaling-down process can render services
within a namespace temporarily inaccessible. Nightscaler addresses this
limitation by enabling the annotation of forced-uptime periods, instructing
kube-downscaler to scale up the affected namespaces during specific time frames.

## Features

- Annotations based on Pub/Sub messages: Nightscaler listens for Pub/Sub
  messages and expects specific attributes (`action`, `namespace`, `cluster`,
  and `duration`) to be present in each message.
- Forced-uptime annotations: When a relevant Pub/Sub message is received,
  nightscaler adds an annotation to the targeted namespace, specifying a
  forced-uptime period using the `downscaler/force-uptime` key-value pair.
- Integration with kube-downscaler: The annotations created by nightscaler
  trigger kube-downscaler to scale up the namespace during the annotated
  forced-uptime period, ensuring the availability of services.

## Getting Started

### Prerequisites

- Kubernetes cluster
- Google Cloud Pub/Sub
- Service Account with Pub/Sub permissions and Workload Identity enabled
- Helm

### Installation

To be added.

## Running locally

First you need to start the Pub/Sub emulator, to do so have a look at the
[.hack](./.hack/README.md) folder.

```shell
export PUBSUB_PROJECT_ID=es-standalone-cb21
export PUBSUB_EMULATOR_HOST=localhost:8085

go run .
```

### Usage

1. Create an interface (e.g., a website) that allows users to trigger actions
   that require forced-uptime periods.
2. Configure the interface to send Pub/Sub messages with the required attributes
   (action, namespace, cluster, and duration) to the nightscaler application.
3. Nightscaler will receive the Pub/Sub messages, process them, and add
   annotations to the specified namespaces accordingly.
4. kube-downscaler will respond to the annotations created by nightscaler,
   scaling up the namespaces during the forced-uptime periods.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports,
please open an issue or submit a pull request.
