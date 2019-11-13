# apbs-rest-detect

<p align="left">
  <a href="https://github.com/Eo300/apbs-rest-detect/actions?query=branch%3Amaster+workflow%3A%22Build+All+Platforms%22"><img alt="GitHub Actions status" src="https://github.com/Eo300/apbs-rest-detect/workflows/Build%20All%20Platforms/badge.svg?branch=master"></a>
</p>

<!-- ![](https://github.com/Eo300/apbs-rest-detect/workflows/.github/workflows/build.yml/badge.svg?branch=master) -->

Tool to detect installation requirements for APBS-REST

Based on the software already installed on your machine, we will give our best recomendation for which method to get [ABPS-REST](https://github.com/Electrostatics/apbs-rest) up and running.  Installation paths include Docker Desktop, Minikube (via Virtualbox), or Minikube (via KVM).

## Usage
Run the binary/executable for your respective system
```shell
./detect
```

You should see output with similar format to the following, respective to your system:

### Linux
```shell
$ ./detect.exe
Target: Linux

Recommended Path:
  Minikube (via VirtualBox)

Required software:
  - VirtualBox
  - Minikube

Needed software...
  - VirtualBox - get from https://www.virtualbox.org/wiki/Downloads
  - Minikube   - get from https://kubernetes.io/docs/tasks/tools/install-minikube/
```

### Windows
Note the additional line denoting Windows edition:
```shell
$ ./detect.exe
Target: Windows

Recommended Path:
  Docker Desktop (w/ Kubernetes)

Windows Edition:
  Microsoft Windows 10 Enterprise

Required software:
  - Docker Desktop

Needed software...
  - Docker Desktop - get from https://docs.docker.com/docker-for-windows/install/
```

### Mac
```shell
$ ./detect.exe
Target: macOS

Recommended Path:
  Docker Desktop (w/ Kubernetes)

Required software:
  - Docker Desktop

Needed software...
  - Docker Desktop - get from https://docs.docker.com/docker-for-mac/install/
```
<br>

If your CPU does not support virtualization for 64-bit systems, you should receive output such as the following:
```shell
$ ./detect
Target: Linux

Unfortunately, your CPU does not support virtualization.
```

## Build
To build from scratch, install the Go compiler, clone this repo, and from its top directory enter the following into your favorite shell:

```
$ go build -o <binary_name> ./src/
```