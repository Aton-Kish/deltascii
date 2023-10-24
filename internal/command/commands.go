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
	"github.com/spf13/cobra"
)

func NewRootCommand(optFns ...func(o *options)) *cobra.Command {
	opts := newOptions(optFns...)

	cmd := &cobra.Command{
		Use:   "deltascii",
		Short: "ΔSCII",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			return nil
		},
		SilenceUsage: true,
	}

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.err)

	return cmd
}

func NewDeltaCommand(optFns ...func(o *options)) *cobra.Command {
	opts := newOptions(optFns...)

	cmd := &cobra.Command{
		Use:     "delta",
		Aliases: []string{"Δ"},
		Short:   "ΔSCII(t) = ASCII(t+1) - ASCII(t)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			return nil
		},
		SilenceUsage: true,
	}

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.err)

	return cmd
}

func NewSummationCommand(optFns ...func(o *options)) *cobra.Command {
	opts := newOptions(optFns...)

	cmd := &cobra.Command{
		Use:     "summation",
		Aliases: []string{"Σ"},
		Short:   "ASCII(t) = ΣΔSCII(t)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			return nil
		},
		SilenceUsage: true,
	}

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.err)

	return cmd
}
