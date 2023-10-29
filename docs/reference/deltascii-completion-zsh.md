## `deltascii completion zsh`

<sub><sup>Last updated on 2023-10-29</sup></sub>

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(deltascii completion zsh)

To load completions for every new session, execute once:

#### Linux:

	deltascii completion zsh > "${fpath[1]}/_deltascii"

#### macOS:

	deltascii completion zsh > $(brew --prefix)/share/zsh/site-functions/_deltascii

You will need to start a new shell for this setup to take effect.


```shell
deltascii completion zsh [flags]
```

### Options

```shell
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### See also

- [deltascii completion](deltascii-completion.md) - Generate the autocompletion script for the specified shell
