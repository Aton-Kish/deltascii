## `deltascii completion fish`

<sub><sup>Last updated on 2023-10-29</sup></sub>

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	deltascii completion fish | source

To load completions for every new session, execute once:

	deltascii completion fish > ~/.config/fish/completions/deltascii.fish

You will need to start a new shell for this setup to take effect.


```shell
deltascii completion fish [flags]
```

### Options

```shell
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### See also

- [deltascii completion](deltascii-completion.md) - Generate the autocompletion script for the specified shell
