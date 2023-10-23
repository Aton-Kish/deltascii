// Copyright (c) 2023 Aton-Kish
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package command

import (
	"context"
	"sync"

	"github.com/spf13/cobra"
)

type deltaCommand struct {
	options *options

	cmd  *cobra.Command
	once sync.Once
}

func NewDeltaCommand(optFns ...func(o *options)) command {
	return &deltaCommand{
		options: newOptions(optFns...),
	}
}

func (c *deltaCommand) Execute(ctx context.Context, args ...string) error {
	cmd := c.command()
	cmd.SetArgs(args)

	if err := cmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}

func (c *deltaCommand) AddCommand(cmds ...command) {
	subs := make([]*cobra.Command, 0, len(cmds))
	for _, cmd := range cmds {
		subs = append(subs, cmd.command())
	}

	c.command().AddCommand(subs...)
}

func (c *deltaCommand) command() *cobra.Command {
	c.once.Do(func() {
		c.cmd = &cobra.Command{
			Use:     "delta",
			Aliases: []string{"Δ"},
			Short:   "ASCII(t+1)-ASCII(t)",
			RunE: func(cmd *cobra.Command, args []string) error {
				if err := cmd.Help(); err != nil {
					return err
				}

				return nil
			},
			SilenceUsage: true,
		}

		c.cmd.SetIn(c.options.stdio.in)
		c.cmd.SetOutput(c.options.stdio.err)
	})

	return c.cmd
}
