## helm-outdated

![Build status](https://github.com/romnn/helm-outdated-charts/actions/workflows/build.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/romnn/helm-outdated-charts)](https://goreportcard.com/report/github.com/romnn/helm-outdated-charts)

Helm plugin to list and update outdated dependencies of a Helm chart,
forked and rewritten from [UniKnow/helm-outdated](https://github.com/UniKnow/helm-outdated).

##### Features:

- Check for outdated helm chart dependencies
- Support for local (`file://`), OCI (`oci://`), and regular (`https://`) chart repositories

##### Future ideas:

- Check for outdated image versions of public images in helm templates

### Installation

```bash
helm plugin install https://github.com/romnn/helm-outdated-charts --version=main

# update with:
helm plugin update outdated
```

_Alternatively_, you can install and use the `helm-outdated` binary directly:

```bash
go install 'github.com/romnn/helm-outdated-charts/cmd/helm-outdated@main'
```

### List outdated dependencies

```bash
helm outdated list ./path/to/chart/
```

## Development

To use the provided tasks in `taskfile.yaml`, install [task](https://taskfile.dev/).

```bash
# view all available tasks
task --list-all
# install development tools
task dev:tools:install
```

After setup, you can use the following tasks during development:

```bash
task tidy
task run:race
task run:race -- list ./path/to/chart
task build:race
task build:executable # build helm-outdated executable
task test
task lint
task format
```

## Acknowledgements

- The now inactive [UniKnow/helm-outdated](https://github.com/UniKnow/helm-outdated) project.

## License

The project is licensed under the same license as [UniKnow/helm-outdated](https://github.com/UniKnow/helm-outdated).
