<img src="imgs/logo.png" width="450">

[![GitHub release](https://img.shields.io/github/release/goodwithtech/dockle.svg)](https://github.com/goodwithtech/dockle/releases/latest)
[![CircleCI](https://circleci.com/gh/goodwithtech/dockle.svg?style=svg)](https://circleci.com/gh/goodwithtech/dockle)
[![Go Report Card](https://goreportcard.com/badge/github.com/goodwithtech/dockle)](https://goreportcard.com/report/github.com/goodwithtech/dockle)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

> Dockle - Simple Security Auditing and helping build the Best Docker Images  

`Dockle` helps you
1) Build secure Docker images
    - Checkpoints includes [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)
2) Build [Best Practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) Docker images

You can check a docker image only run... 
```bash
$ brew install goodwithtech/dockle/dockle
$ dockle [YOUR_IMAGE_NAME]
``` 

<img src="imgs/usage_pass_light.png" width="800">
<img src="imgs/usage_fail_light.png" width="800">


# TOC

- [Features](#features)
- [Comparison](#comparison)
- [Installation](#installation)
  - [Linuxbrew](#linuxbrew)
  - [RHEL/CentOS](#rhelcentos)
  - [Debian/Ubuntu](#debianubuntu)
  - [Mac OS X / Homebrew](#mac-os-x--homebrew)
  - [Windows](#windows)
  - [Binary](#binary)
  - [From source](#from-source)
- [Checkpoint Summary](#checkpoint-summary)
- [Quick Start](#quick-start)
  - [Basic](#basic)
  - [Docker](#docker)
- [Examples](#examples)
  - [Scan an image](#scan-an-image)
  - [Scan an image file](#scan-an-image-file)
  - [Specify exit code](#specify-exit-code)
  - [Ignore the specified checkpoints](#ignore-the-specified-checkpoints)
  - [Clear image caches](#clear-image-caches)
- [Continuous Integration](#continuous-integration-ci)
  - [Travis CI](#travis-ci)
  - [CircleCI](#circleci)
  - [Authorization for Private Docker Registry](#authorization-for-private-docker-registry)
- [Checkpoint Detail](#checkpoint-detail)
  - [CIS's Docker Image Checkpoints](#docker-image-checkpoints)
  - [Dockle Checkpoints for Docker](#dockle-checkpoints-for-docker)
  - [Dockle Checkpoints for Linux](#dockle-checkpoints-for-linux)
- [Credits](#credits)
- [Roadmap](#roadmap)


# Features

- Detect container's vulnerabilities
- Helping build Dockerfile best-practice
- Simple
  - Specify only an image name
  - See [Quick Start](#quick-start) and [Examples](#examples)
- Support CIS Benchmarks
  - High accuracy
- DevSecOps
  - Suitable for CI such as Travis CI, CircleCI, Jenkins, etc.
  - See [CI Example](#continuous-integration-ci)


# Comparison



|  | [Dockle](https://github.com/goodwithtech/dockle) | [Hadolint](https://github.com/hadolint/hadolint) | [Docker Bench for Security](https://github.com/docker/docker-bench-security) |
|--- |---:|---:|---:|
| Target |  BuildImage | Dockerfile | Host<br/>DockerDaemon<br/>BuildImage<br/>ContainerRuntime |
| How to run | Binary | Binary | ShellScript |
| Dependency | No | No | Some dependencies |
| Suitable CI | ✓ | ✓ | x |
| Purpose |SecurityAudit<br/>DockerfileLint| DockerfileLint | SecurityAudit<br />DockerfileLint |
| Covered CIS Benchmarks(Docker Image and Build File) | 7 | 3 | 5 |


<details>
<summary>Detail of CIS Benchmark</summary>

|  | [Dockle](https://github.com/goodwithtech/dockle) | [Docker Bench for Security](https://github.com/docker/docker-bench-security) | [Hadolint](https://github.com/hadolint/hadolint) | 
|---|:---:|:---:|:---:|
| 1.  Create a user for the container | ✓ | ✓ | ✓ |
| 2.  Use trusted base images for containers | - | – | - |
| 3.  Do not install unnecessary packages in the container | - | - | - | 
| 4.  Scan and rebuild the images to include security patches | - | - | - | 
| 5.  Enable Content trust for Docker | ✓ | ✓ | - |
| 6.  Add HEALTHCHECK instruction to the container image | ✓ | ✓ | - |
| 7.  Do not use update instructions alone in the Dockerfile | ✓ | ✓ | ✓|
| 8.  Remove setuid and setgid permissions in the images | ✓ | - | - |
| 9.  Use COPY instead of ADD in Dockerfile | ✓ | ✓ | ✓|
| 10. Do not store secrets in Dockerfiles | ✓ | - | - |
| 11. Install verified packages only | -  |  - | - |
| |7|5|3|

All checkpoints [here](#checkpoint-summary)!
</details>


# Installation

## Linuxbrew

```bash
$ brew install goodwithtech/dockle/dockle
```
## RHEL/CentOS

```bash
$ VERSION=$(
 curl --silent "https://api.github.com/repos/goodwithtech/dockle/releases/latest" | \
 grep '"tag_name":' | \
 sed -E 's/.*"v([^"]+)".*/\1/' \
) && rpm -ivh https://github.com/goodwithtech/dockle/releases/download/v${VERSION}/dockle_${VERSION}_Linux-64bit.rpm
```

## Debian/Ubuntu

```bash
$ VERSION=$(
 curl --silent "https://api.github.com/repos/goodwithtech/dockle/releases/latest" | \
 grep '"tag_name":' | \
 sed -E 's/.*"v([^"]+)".*/\1/' \
) && curl -L -o dockle.deb https://github.com/goodwithtech/dockle/releases/download/v${VERSION}/dockle_${VERSION}_Linux-64bit.deb
$ sudo dpkg -i dockle.deb && rm dockle.deb
```

## Mac OS X / Homebrew

You can use homebrew on Mac OS.

```
$ brew install goodwithtech/dockle/dockle
```

## Windows

```bash
$ VERSION=$(
 curl --silent "https://api.github.com/repos/goodwithtech/dockle/releases/latest" | \
 grep '"tag_name":' | \
 sed -E 's/.*"v([^"]+)".*/\1/' \
) && curl -L -o dockle.zip https://github.com/goodwithtech/dockle/releases/download/v${VERSION}/dockle_${VERSION}_Windows-64bit.zip
$ unzip dockle.zip && rm dockle.zip
$ ./dockle.exe [IMAGE_NAME]
```


## Binary

Get the latest version from [this page](https://github.com/goodwithtech/dockle/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH` (on UNIX-y systems, /usr/local/bin or the like). Make sure it has execution bits turned on.


## From source

```sh
$ go get -u github.com/goodwithtech/dockle
```

# Quick Start

## Basic

Simply specify an image name (and a tag).

```bash
$ dockle [YOUR_IMAGE_NAME]
```

<details>
<summary>Result</summary>

```
FATAL   - Create a user for the container
        * Last user should not be root
WARN    - Enable Content trust for Docker
        * export DOCKER_CONTENT_TRUST=1 before docker pull/build
FATAL   - Add HEALTHCHECK instruction to the container image
        * not found HEALTHCHECK statement
FATAL   - Do not use update instructions alone in the Dockerfile
        * Use 'Always combine RUN apt-get update with apt-get install' : /bin/sh -c apt-get update && apt-get install -y git
PASS    - Remove setuid and setgid permissions in the images
FATAL   - Use COPY instead of ADD in Dockerfile
        * Use COPY : /bin/sh -c #(nop) ADD file:81c0a803075715d1a6b4f75a29f8a01b21cc170cfc1bff6702317d1be2fe71a3 in /app/credentials.json
FATAL   - Do not store secrets in ENVIRONMENT variables
        * Suspicious ENV key found : MYSQL_PASSWD
FATAL   - Do not store secret files
        * Suspicious filename found : app/credentials.json
PASS    - Avoid sudo command
FATAL   - Avoid sensitive directory mounting
        * Avoid mounting sensitive dirs : /usr
PASS    - Avoid apt-get/apk/dist-upgrade
PASS    - Use apk add with --no-cache
FATAL   - Clear apt-get caches
        * Use 'apt-get clean && rm -rf /var/lib/apt/lists/*' : /bin/sh -c apt-get update && apt-get install -y git
PASS    - Avoid latest tag
FATAL   - Avoid empty password
        * No password user found! username : nopasswd
PASS    - Be unique UID
PASS    - Be unique GROUP
```
</details>

## Docker

Replace [YOUR_CACHE_DIR] with the cache directory on your machine.

```
$ docker run --rm -v [YOUR_CACHE_DIR]:/root/.cache/ goodwithtech/dockle [YOUR_IMAGE_NAME]
```

Example for macOS:

```
$ docker run --rm -v $HOME/Library/Caches:/root/.cache/ goodwithtech/dockle [YOUR_IMAGE_NAME]
```

If you would like to scan the image on your host machine, you need to mount `docker.sock`.

```
$ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock ...
```

Please re-pull latest `goodwithtech/dockle` if an error occured.

# Checkpoint Summary

| CODE | DESCRIPTION | LEVEL[※](#level) |
|---|---|:---:|
| | [CIS's Docker Image Checkpoints](#docker-image-checkpoints) | |
| [CIS-DI-0001](#cis-di-0001-create-a-user-for-the-container) | Create a user for the container | WARN |
| [CIS-DI-0002](#cis-di-0002-use-trusted-base-images-for-containers) | Use trusted base images for containers | FATAL
| [CIS-DI-0003](#cis-di-0003-do-not-install-unnecessary-packages-in-the-container) | Do not install unnecessary packages in the container | FATAL
| [CIS-DI-0004](#cis-di-0004-scan-and-rebuild-the-images-to-include-security-patches) | Scan and rebuild the images to include security patches | FATAL
| [CIS-DI-0006](#cis-di-0006-add-healthcheck-instruction-to-the-container-image) | Add HEALTHCHECK instruction to the container image | WARN
| [CIS-DI-0007](#cis-di-0007-do-not-use-update-instructions-alone-in-the-dockerfile) | Do not use update instructions alone in the Dockerfile | FATAL
| [CIS-DI-0008](#cis-di-0008-remove-setuid-and-setgid-permissions-in-the-images) | Remove setuid and setgid permissions in the images | INFO
| [CIS-DI-0009](#cis-di-0009-use-copy-instead-of-add-in-dockerfile) | Use COPY instead of ADD in Dockerfile | FATAL
| [CIS-DI-0010](#cis-di-0010-do-not-store-secrets-in-dockerfiles) | Do not store secrets in Dockerfiles | FATAL
| [CIS-DI-0011](#cis-di-0011-install-verified-packages-only) | Install verified packages only | INFO
|| [Dockle Checkpoints for Docker](#dockle-checkpoints-for-docker) |
| [DKL-DI-0001](#dkl-di-0001-avoid-sudo-command) | Avoid `sudo` command | FATAL
| [DKL-DI-0002](#dkl-di-0002-avoid-sensitive-directory-mounting) | Avoid sensitive directory mounting | FATAL
| [DKL-DI-0003](#dkl-di-0003-avoid-apt-get-upgrade-apk-upgrade-dist-upgrade) | Avoid `apt-get upgrade`, `apk upgrade`, `dist-upgrade` | FATAL
| [DKL-DI-0004](#dkl-di-0004-use-apk-add-with---no-cache) | Use apk add with `--no-cache` | FATAL
| [DKL-DI-0005](#dkl-di-0005-clear-apt-get-caches) | Clear apt-get caches | FATAL
| [DKL-DI-0006](#dkl-di-0006-avoid-latest-tag) | Avoid `latest` tag | WARN
|| [Dockle Checkpoints for Linux](#dockerdockle-checkpoints-for-linux) |
| [DKL-LI-0001](#dkl-li-0001-avoid-empty-password) | Avoid empty password | FATAL
| [DKL-LI-0002](#dkl-li-0002-be-unique-uidgroups) | Be unique UID/GROUPs | FATAL

#### Level

`Dockle` has 5 check levels

| LEVEL | DESCRIPTION |
|:---:|---|
| FATAL | Be practical and prudent |
| WARN | Be practical and prudent, but limited uses (official docker image ) |
| INFO | May negatively inhibit the utility or performance |
| SKIP | Not found target files |
| PASS | Not found any problems |

# Examples

## Scan an image

Simply specify an image name (and a tag).

```bash
$ dockle goodwithtech/test-image:v1
```
<details>
<summary>Result</summary>

```
FATAL   - CIS-DI-0001: Create a user for the container
        * Last user should not be root
WARN    - CIS-DI-0005: Enable Content trust for Docker
        * export DOCKER_CONTENT_TRUST=1 before docker pull/build
FATAL   - CIS-DI-0006: Add HEALTHCHECK instruction to the container image
        * not found HEALTHCHECK statement
FATAL   - CIS-DI-0007: Do not use update instructions alone in the Dockerfile
        * Use 'Always combine RUN 'apt-get update' with 'apt-get install' : /bin/sh -c apt-get update && apt-get install -y git
FATAL   - CIS-DI-0008: Remove setuid and setgid permissions in the images
        * Found setuid file: etc/passwd grw-r--r--
        * Found setuid file: usr/lib/openssh/ssh-keysign urwxr-xr-x
        * Found setuid file: app/hoge.txt ugrw-r--r--
        * Found setuid file: app/hoge.txt ugrw-r--r--
        * Found setuid file: etc/shadow urw-r-----
FATAL   - CIS-DI-0009: Use COPY instead of ADD in Dockerfile
        * Use COPY : /bin/sh -c #(nop) ADD file:81c0a803075715d1a6b4f75a29f8a01b21cc170cfc1bff6702317d1be2fe71a3 in /app/credentials.json
FATAL   - CIS-DI-0010: Do not store secrets in ENVIRONMENT variables
        * Suspicious ENV key found : MYSQL_PASSWD
FATAL   - CIS-DI-0010: Do not store secret files
        * Suspicious filename found : app/credentials.json
PASS    - DKL-DI-0001: Avoid sudo command
FATAL   - DKL-DI-0002: Avoid sensitive directory mounting
        * Avoid mounting sensitive dirs : /usr
PASS    - DKL-DI-0003: Avoid apt-get/apk/dist-upgrade
PASS    - DKL-DI-0004: Use apk add with --no-cache
FATAL   - DKL-DI-0005: Clear apt-get caches
        * Use 'apt-get clean && rm -rf /var/lib/apt/lists/*' : /bin/sh -c apt-get update && apt-get install -y git
PASS    - DKL-DI-0006: Avoid latest tag
FATAL   - DKL-LI-0001: Avoid empty password
        * No password user found! username : nopasswd
PASS    - DKL-LI-0002: Be unique UID
PASS    - DKL-LI-0002: Be unique GROUP
```
</details>

## Scan an image file

```bash
$ docker save alpine:latest -o alpine.tar
$ dockle --input alpine.tar
```

## Specify exit code
By default, `Dockle` exits with code 0 even if there are some problems.
Use the --exit-code option if you may want to exit with a non-zero exit code.

```bash
$ dockle  -exist-code 1 [IMAGE_NAME]
```

## Ignore the specified checkpoints

Use `.dockleignore`.

```bash
$ cat .dockleignore
# set root to default user because we want to run nginx
CIS-DI-0001
# Use latest tag because only check for image inside
DKL-DI-0006
```

### Clear image caches

The `--clear-cache` option removes image caches. This option is useful if the image which has the same tag is updated (such as when using `latest` tag).

```
$ dockle --clear-cache python:3.7
```

# Continuous Integration (CI)

Scan your image built in Travis CI/CircleCI. 
The test will fail with if a error is found.
You can skip target checkpoint if you use `.dockleignore`.
 
When you don't would like to fail the test, specify `--exit-code 0`.


## Travis CI

```yaml
services:
  - docker

env:
  global:
    - COMMIT=${TRAVIS_COMMIT::8}

before_install:
  - docker build -t dockle-ci-test:${COMMIT} .
  - export VERSION=$(curl --silent "https://api.github.com/repos/goodwithtech/dockle/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
  - wget https://github.com/goodwithtech/dockle/releases/download/v${VERSION}/dockle_${VERSION}_Linux-64bit.tar.gz
  - tar zxvf dockle_${VERSION}_Linux-64bit.tar.gz
script:
  - ./dockle dockle-ci-test:${COMMIT}
  - ./dockle --exit-code 1 dockle-ci-test:${COMMIT}
```

Example: https://travis-ci.org/goodwithtech/dockle-ci-test<br/>
Repository: https://github.com/goodwithtech/dockle-ci-test

## CircleCI

```yaml
jobs:
  build:
    docker:
      - image: docker:18.09-git
    steps:
      - checkout
      - setup_remote_docker
      - run:
          name: Build image
          command: docker build -t dockle-ci-test:${CIRCLE_SHA1} .
      - run:
          name: Install dockle
          command: |
            apk add --update curl
            VERSION=$(
                curl --silent "https://api.github.com/repos/goodwithtech/dockle/releases/latest" | \
                grep '"tag_name":' | \
                sed -E 's/.*"v([^"]+)".*/\1/'
            )
            wget https://github.com/goodwithtech/dockle/releases/download/v${VERSION}/dockle_${VERSION}_Linux-64bit.tar.gz
            tar zxvf dockle_${VERSION}_Linux-64bit.tar.gz
            mv dockle /usr/local/bin
      - run:
          name: Scan the local image with dockle
          command: dockle --exit-code 1 dockle-ci-test:${CIRCLE_SHA1}
workflows:
  version: 2
  release:
    jobs:
      - build
```

Example : https://circleci.com/gh/goodwithtech/dockle-ci-test<br/>
Repository: https://github.com/goodwithtech/dockle-ci-test

## Authorization for Private Docker Registry

`Dockle` can download images from private registry, without installing `Docker` and any 3rd party tools.
That's because it's easy to run in a CI process.

All you have to do is install `Dockle` and set ENVIRONMENT variables.
But, I can't recommend using ENV vars in your local machine to you.

### Docker Hub

Docker Hub needs `DOCKLE_AUTH_URL`, `DOCKLE_USERNAME` and `DOCKLE_PASSWORD`.
You don't need to set ENV vars when download from public repository.

```bash
export DOCKLE_AUTH_URL=https://registry.hub.docker.com
export DOCKLE_USERNAME={DOCKERHUB_USERNAME}
export DOCKLE_PASSWORD={DOCKERHUB_PASSWORD}
```

### Amazon ECR (Elastic Container Registry)

`Dockle` uses AWS SDK. You don't need to install `aws` CLI tool.
You can use [AWS CLI's ENVIRONMENT variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html).

### GCR (Google Container Registry)

`Dockle` uses Google Cloud SDK. So, you don't need to install `gcloud` command.

If you would like to use target project's repository, you can settle via `GOOGLE_APPLICATION_CREDENTIAL`. 
```bash
# must set DOCKLE_USERNAME empty char
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credential.json
```

### Self Hosted Registry (BasicAuth)

BasicAuth server needs `DOCKLE_USERNAME` and `DOCKLE_PASSWORD`.

```bash
export DOCKLE_USERNAME={USERNAME}
export DOCKLE_PASSWORD={PASSWORD}

# if you'd like to use 80 port, use NonSSL
export DOCKLE_NON_SSL=true
```

# Checkpoint Detail

## Docker Image Checkpoints

These checkpoints refered to [CIS Docker 1.13.0 Benchmark v1.0.0](https://www.cisecurity.org/benchmark/docker/).

### CIS-DI-0001: Create a user for the container

> Create a non-root user for the container in the Dockerfile for the container image.<br/>
> It is a good practice to run the container as a non-root user, if possible. 
```
# Dockerfile
RUN useradd -d /home/dockle -m -s /bin/bash dockle
USER dockle

or 

RUN addgroup -S dockle && adduser -S -G dockle dockle
USER dockle

```


### CIS-DI-0002: Use trusted base images for containers

Dockle checks Content Trust. 
Please check with [Trivy](https://github.com/knqyf263/trivy).

### CIS-DI-0003: Do not install unnecessary packages in the container

Not supported.

### CIS-DI-0004: Scan and rebuild the images to include security patches

Not supported.
Please check with [Trivy](https://github.com/knqyf263/trivy).

### CIS-DI-0005: Enable Content trust for Docker

> Content trust is disabled by default. You should enable it.

```bash
$ export DOCKER_CONTENT_TRUST=1 
```

https://docs.docker.com/engine/security/trust/content_trust/#about-docker-content-trust-dct
> Docker Content Trust (DCT) provides the ability to use digital signatures for data sent to and received from remote Docker registries.
> Engine Signature Verification prevents the following:
> - `$ docker container run` of an unsigned image.
> - `$ docker pull` of an unsigned image.
> - `$ docker build` where the FROM image is not signed or is not scratch.

### CIS-DI-0006: Add HEALTHCHECK instruction to the container image

> Add `HEALTHCHECK` instruction in your docker container images to perform the health check on running containers.<br/>
> Based on the reported health status, the docker engine could then exit non-working containers and instantiate new ones.
```
# Dockerfile
HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost/ || exit 1
```

### CIS-DI-0007: Do not use update instructions alone in the Dockerfile

> Do not use update instructions such as apt-get update alone or in a single line in the Dockerfile.<br/>
> Adding the update instructions in a single line on the Dockerfile will cache the update layer.
```bash
RUN apt-get update && apt-get install -y package-a
```

### CIS-DI-0008: Remove setuid and setgid permissions in the images

> Removing setuid and setgid permissions in the images would prevent privilege escalation attacks in the containers.<br/>
> setuid and setgid permissions could be used for elevating privileges.
```bash
chmod u-s setuid-file
chmod u-g setgid-file
```

### CIS-DI-0009: Use COPY instead of ADD in Dockerfile

> Use COPY instruction instead of ADD instruction in the Dockerfile.<br/>
> `ADD` instruction introduces risks such as adding malicious files from URLs without scanning and unpacking procedure vulnerabilities.
```
# Dockerfile
ADD test.json /app/test.json
↓
COPY test.json /app/test.json
```

### CIS-DI-0010: Do not store secrets in Dockerfiles

> Do not store any secrets in Dockerfiles.<br/>
> the secrets within these Dockerfiles could be easily exposed and potentially be exploited.

`Dockle` checks ENVIRONMENT variables and credential files.

### CIS-DI-0011: Install verified packages only

Not supported.
It's better to use [Trivy](https://github.com/knqyf263/trivy).

## Dockle Checkpoints for Docker

These checkpoints refered to [Docker Best Practice](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) and so on.

### DKL-DI-0001: Avoid `sudo` command

https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#user
> Avoid installing or using sudo as it has unpredictable TTY and signal-forwarding behavior that can cause problems.

### DKL-DI-0002: Avoid sensitive directory mounting

A volume mount makes weakpoints. 
This depends on mounting volumes.
Currently, `Dockle` checks following directories.

`/boot`,`/dev`,`/etc`,`/lib`,`/proc`,`/sys`, `/usr`

`dockle` only checks `VOLUME` statements. We can't check `docker run -v /lib:/lib ...`.


### DKL-DI-0003: Avoid `apt-get upgrade`, `apk upgrade`, `dist-upgrade`

https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#apt-get
 
> Avoid RUN apt-get upgrade and dist-upgrade, as many of the “essential” packages from the parent images cannot upgrade inside an unprivileged container.

### DKL-DI-0004: Use apk add with `--no-cache`

https://github.com/gliderlabs/docker-alpine/blob/master/docs/usage.md#disabling-cache

> As of Alpine Linux 3.3 there exists a new --no-cache option for apk. It allows users to install packages with an index that is updated and used on-the-fly and not cached locally:
> This avoids the need to use --update and remove /var/cache/apk/* when done installing packages.

### DKL-DI-0005: Clear apt-get caches

Use “apt-get clearn && rm -rf /var/lib/apt/lists/*` if use apt-get install

https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#apt-get
> In addition, when you clean up the apt cache by removing /var/lib/apt/lists it reduces the image size, since the apt cache is not stored in a layer. Since the RUN statement starts with apt-get update, the package cache is always refreshed prior to apt-get install.


### DKL-DI-0006: Avoid `latest` tag

https://vsupalov.com/docker-latest-tag/

> Docker images tagged with :latest have caused many people a lot of trouble.

## Dockle Checkpoints for Linux

These checkpoints refered to [Linux Best Practices](https://www.cyberciti.biz/tips/linux-security.html) and so on.

### DKL-LI-0001: Avoid empty password 

https://blog.aquasec.com/cve-2019-5021-alpine-docker-image-vulnerability

> CVE-2019-5021: Alpine Docker Image ‘null root password’ Vulnerability

### DKL-LI-0002: Be unique UID/GROUPs

http://www.linfo.org/uid.html

> Contrary to popular belief, it is not necessary that each entry in the UID field be unique. However, non-unique UIDs can cause security problems, and thus UIDs should be kept unique across the entire organization.

# Credits

Special Thanks to [@knqyf263](https://github.com/knqyf263) (Teppei Fukuda) and [Trivy](https://github.com/knqyf263/trivy)

# License

AGPLv3

# Author

[@tomoyamachi](https://github.com/tomoyamachi) (Tomoya Amachi)

# Roadmap
- [ ] JSON output
- [ ] Check php.ini file
- [ ] Check nginx.conf file
- [ ] create CI badges
- Check /etc/hosts
  - [ ] duplicates
  - [ ] hostname
  - [ ] localhost
- Packages
  - [ ] Package managers
- File Permissions
  - [ ] Insecure permission
- Image Size
  - [ ] check large size container
  
  
## if running docker daemon...
- Networking
  - [ ] `docker port container` if docker running
  - [ ] by file
    - /proc/1/net/tcp : openning port (if running)
- Volume mount
  - mount dangerous 
    - /boot, /dev, /etc, /lib

