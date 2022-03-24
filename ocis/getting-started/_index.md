---
title: "Getting Started"
date: 2020-02-27T20:35:00+01:00
weight: 0
geekdocRepo: https://github.com/owncloud/ocis
geekdocEditPath: edit/master/docs/ocis/getting-started
geekdocFilePath: _index.md
---

{{< toc >}}

## oCIS online demo

We have an oCIS demo instance running on [ocis.owncloud.com](https://ocis.owncloud.com) where you can get a first impression of it.

We also have some more variations of oCIS running and [continuously deployed]({{< ref "../deployment/continuous_deployment" >}}) to reflect different scenarios in that oCIS might be used.

## Run oCIS

We are distributing oCIS as binaries and Docker images.

{{< hint warning >}}
The examples in this document assume that oCIS is accessed from the same host as it is running on (`localhost`). If you would like
to access oCIS remotely please refer to the [Basic Remote Setup]({{< ref "../deployment/basic-remote-setup" >}}) section. Especially
to the notes about setting the `PROXY_HTTP_ADDR` and `OCIS_URL` enviroment variables.
{{< /hint >}}

You can find more deployment examples in the [deployment section]({{< ref "../deployment" >}}).

### Binaries

You can find the latest official release of oCIS at [our download mirror](https://download.owncloud.com/ocis/ocis/) or on [GitHub](https://github.com/owncloud/ocis/releases).
The latest build from the master branch can be found at [our download mirrors testing section](https://download.owncloud.com/ocis/ocis/testing/).

To run oCIS as binary you need to download it first and then run the following commands.
For this example, assuming version 1.16.0 of oCIS running on a Linux AMD64 host:

```console
# download
curl https://download.owncloud.com/ocis/ocis/1.17.0/ocis-1.17.0-linux-amd64 --output ocis

# make binary executable
chmod +x ocis

# run
OCIS_INSECURE=true ./ocis server
```

The default primary storage location is `~/.ocis` or `/var/lib/ocis` depending on the packaging format and your operating system user. You can change that value by configuration.

{{< hint info >}}
When you're using oCIS with self-signed certificates, you need to set the environment variable `OCIS_INSECURE=true`, in order to make oCIS work.
{{< /hint >}}

{{< hint warning >}}
oCIS by default relies on Multicast DNS (mDNS), usually via avahi-daemon. If your system has a firewall, make sure mDNS is allowed in your active zone.
{{< /hint >}}

### Docker

Docker images for oCIS are available on [Docker Hub](https://hub.docker.com/r/owncloud/ocis).

The `latest` tag always reflects the current master branch.

```console
docker pull owncloud/ocis
docker run --rm -ti -p 9200:9200 -e OCIS_INSECURE=true owncloud/ocis
```

{{< hint info >}}
When you're using oCIS with self-signed certificates, you need to set the environment variable `OCIS_INSECURE=true`, in order to make oCIS work.
{{< /hint >}}

{{< hint warning >}}
In order to persist your data, you need to mount a docker volume or create a host bind-mount at `/var/lib/ocis`, for example with: `-v /some/host/dir:/var/lib/ocis`

You cannot use bind mounts on MacOS, since extended attributes are not supported ([owncloud/ocis#182](https://github.com/owncloud/ocis/issues/182), [moby/moby#1070](https://github.com/moby/moby/issues/1070)).
{{< /hint >}}

## Usage

### Login to ownCloud Web

Open [https://localhost:9200](https://localhost:9200) and [login using one of the demo accounts]({{< ref "./demo-users" >}}).

### Basic Management Commands

The oCIS single binary contains multiple extensions and the `ocis` command helps you to manage them. You already used `ocis server` to run all available extensions in the [Run oCIS]({{< ref "#run-ocis" >}}) section. We now will show you some more management commands, which you may also explore by typing `ocis --help` or going to the [docs]({{< ref "../config" >}}).

To start oCIS server:

{{< highlight txt >}}
ocis server
{{< / highlight >}}

The list command prints all running oCIS extensions.
{{< highlight txt >}}
ocis list
{{< / highlight >}}

To stop a particular extension:
{{< highlight txt >}}
ocis kill web
{{< / highlight >}}

To start a particular extension:
{{< highlight txt >}}
ocis run web
{{< / highlight >}}

The version command prints the version of your installed oCIS.
{{< highlight txt >}}
ocis --version
{{< / highlight >}}

The health command is used to execute a health check, if the exit code equals zero the service should be up and running, if the exist code is greater than zero the service is not in a healthy state. Generally this command is used within our Docker containers, it could also be used within Kubernetes.

{{< highlight txt >}}
ocis health --help
{{< / highlight >}}