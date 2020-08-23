---
title: "Test Example"
linkTitle: "Test Example"
weight: 15
description: >
  Eat your Own DogFoot, described How this blog will be deliverd.
---

## Testing the Ci/Cd Blog

Example of a CI/CD Process for this project.

### Artefacts

* Markdown, and generated HTML
* Dockerfile
* Helm Chart

### Tooling

* [hadolint/hadolint](https://github.com/hadolint/hadolint)
* [igorshubovych/markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli)
* [terratest](https://github.com/gruntwork-io/terratest)
* [skaffold.dev](hhttps://skaffold.dev)

## Test Elements

### Unit Tests

#### Static Tests

##### Dockerfile

Check the Dockerfile from `pitch/Dockerfile` with [hadolint/hadolint](https://github.com/hadolint/hadolint).

```sh
docker run --rm -i hadolint/hadolint < pitch/Dockerfile
```

###### Docker Terratest E2E

```sh
cd ./tests
go test -v -run TestDockerElevatorPitch
```

##### Helm Chart

###### Helm Terratest E2E

```sh
cd ./tests
go test -v -run TestHelmCiCdPitchDeployment
```

##### Markdown Content

```sh
markdownlint -i '{**/pitch/themes/**,**/node_modules/**}' .
```

## Development

For Production near local development use a combination of [Kind](https://kind.sigs.k8s.io/) and [skaffold](https://skaffold.dev), 
so you will get a Real Time Hot deployment with port Forward.

```sh
skaffold dev --no-prune=false --cache-artifacts=false --port-forward
```

Checking the Running deployment with helm test.

```sh
helm test my-release
```
