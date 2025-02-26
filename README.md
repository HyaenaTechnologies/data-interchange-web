[CommonMark]:https://commonmark.org/
[Containerman]: https://podman.io/
[Containerman Pod]: https://docs.podman.io/en/latest/markdown/podman-pod.1.html
[Etcd-IO]: https://etcd.io/
[Fleet]: https://jetbrains.com/fleet
[Greptime Database]: https://greptime.com/
[HTTP]: https://developer.mozilla.org/en-US/docs/Web/HTTP
[JSON]: https://www.json.org/json-en.html
[K8S Control]: https://kubernetes.io/
[K8S Kompose]: https://kompose.io/
[K8S Kube]: https://minikube.sigs.k8s.io/docs/
[K8S Native]: https://knative.dev/docs/
[K8S Ops]: https://kops.sigs.k8s.io/
[K8S Skaffold]: https://skaffold.dev/
[MDN]: https://developer.mozilla.org/en-US/docs/Web/API
[Minio Database]: https://min.io/
[Moby]: http://docker.com
[Moby Compose]: https://docs.docker.com/reference/cli/docker/compose/
[Moby Swarm]: https://docs.docker.com/reference/cli/docker/swarm/
[Rust Language]: https://rust-lang.org
[RustRover]: https://jetbrains.com/rust
[Spice Database]: https://authzed.com/
[Surreal Database]: https://surrealdb.com/
[TOML]: https://toml.io/en/
[VSCode]: https://code.visualstudio.com/docs
[WAI-ARIA]: https://www.w3.org/WAI/ARIA/apg/patterns/
[YAML]: https://yaml.org/

<a href="https://github.com/HyaenaTechnologies/data-interchange-web">
  <h1>
    <picture>
      <img src="https://github.com/HyaenaTechnologies/data-interchange-web/blob/main/assets/di_markdown.png" alt="">
    </picture>
  </h1>
</a>

[![Rust Workflow](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/rust.yml/badge.svg)](https://github.com/HyaenaTechnologies/data-interchange-web/actions/workflows/rust.yml)

# Data Interchange Web

Data Interchange is a Network Application Programming Interface Development Platform

## Features

- Comma-Separated Values
- DNS Protocol
- Environment Variables
- [Hypertext Transfer Protocol][HTTP]
- Internet Protocol
- [JavaScript Object Notation][JSON]
- [Markdown][CommonMark]
- [Tom's Obvious Minimal Language][TOML]
- Transmission Control Protocol
- Transport Layer Security
- User Datagram Protocol
- Web Socket Protocol
- [YAML Ain't Markup Language][YAML]

## Build
- **_Container Engines:_** [Docker][Moby], [Podman][Containerman]

- **_Container Orchestration Engines:_** [Docker Compose][Moby Compose], [Docker Swarm][Moby Swarm], [KOps][K8S Ops], [Minikube][K8S Kube], [Podman Pods][Containerman Pod]

- **_Container Tools:_** [Knative][K8S Native], [Kompose][K8S Kompose], [Kubectl][K8S Control], [Skaffold][K8S Skaffold]

- [Rust][Rust Language]
- [JetBrains Fleet][Fleet]
- [JetBrains RustRover][RustRover]
- [Mozilla Developer Network Web Documentation][MDN]
- [Visual Studio Code][VSCode]
- [Web Accessibility Initiative][WAI-ARIA]

## Databases

- **_Authorization:_** [Spice][Spice Database]

- **_Graph:_** [Surreal][Surreal Database]

- **_Key Value:_** [Etcd][Etcd-IO]

- **_Object Storage:_** [Minio][Minio Database]

- **_Time Series:_** [Greptime][Greptime Database]

## Build Web Server

```shell
git clone

cargo build --release 
mv ./target/release/data-interchange-web ./binary
mv ./binary/data-interchange-web ./binary/htdinet
./binary/htdinet serve
```

OR

```shell
git clone

make build-release
mv ./target/release/data-interchange-web ./binary
mv ./binary/data-interchange-web ./binary/htdinet
./binary/htdinet serve
```

OR

```shell
git clone

./shell/rust.bash
mv ./target/release/data-interchange-web ./binary
mv ./binary/data-interchange-web ./binary/htdinet
./binary/htdinet serve
```

## Containerize Web Server

```shell
git clone

docker build ./ --file ./data-interchange.Dockerfile --tag data-interchange-web:latest
```

OR

```shell
git clone

make docker-build
```

OR

```shell
git clone

./shell/docker.bash
```

## Install Web Server

```shell
echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> ~/.bashrc
sudo echo 'export PATH="$PATH:/usr/local/bin/htdinet"' >> /etc/skel/.bashrc
```
