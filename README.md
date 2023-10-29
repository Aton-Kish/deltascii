# ΔSCII

ΔSCII is a tool for easily editing [asciinema](https://asciinema.org/) [asciicast](https://github.com/asciinema/asciinema/blob/main/doc/asciicast-v2.md) files.

## Installation

### Manually

Download the pre-compiled binaries from the [releases page](https://github.com/Aton-Kish/deltascii/releases).

### `go install`

```shell
go install github.com/Aton-Kish/deltascii@latest
```

## Usage

```shell
: delta
deltascii Δ -i ascii.cast -o deltascii.cast

: accumulate
deltascii Σ -i deltascii.cast -o ascii.cast
```

If you want to learn more, check out the [user guide](docs/README.md).

## Troubleshooting

If you think you've found a bug, or something isn't behaving the way you think it should, please raise an [issue](https://github.com/Aton-Kish/deltascii/issues/new/choose) on GitHub.

## Changelog

Refer to the [CHANGELOG](./CHANGELOG.md).

## License

ΔSCII is licensed under the MIT License, see [LICENSE](./LICENSE).
