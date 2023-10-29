# Developer guide

## Setting up

requirements:

- [Go](https://go.dev/) (v1.21.1 or higher)
- [Task](https://taskfile.dev/) (v3.30.1 or higher)

install dependencies:

```shell
task install
```

## Workflow

> [!NOTE]
> You can discover all available tasks by running `task -l`.

Here, we introduce the main tasks.

### Linting

```shell
: linting
task lint
: formatting
task format
```

### Generating code

Some documents, such as the command reference (`docs/reference/*.md`), are automatically generated.

For automatic generation, we use [`go generate`](https://go.dev/blog/generate).
You can find hints about the content of automatic generation by looking for code comments that start with `//go:generate`.

```shell
task generate
```

### Testing

```shell
task test
```

### Build

```shell
: build executable binary only
task build
: build and archive executable binary
task archive
```

### Release

> [!WARNING]
> The specified version will be released, so double-check before executing.

The release task is available only on **main** branch.
`NEXT_VERSION` environment variable is required and must follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

```shell
git checkout main
NEXT_VERSION=vx.x.x task release
```
