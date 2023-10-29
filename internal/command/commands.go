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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Aton-Kish/deltascii/internal/asciinema"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var (
	version = "unknown"
)

func SetVersion(v string) {
	version = v
}

func Register() *xcommand {
	rootCmd := newRootCommand()
	deltaCmd := newDeltaCommand()
	accCmd := newAccumulateCommand()

	rootCmd.AddCommand(deltaCmd.Command, accCmd.Command)

	return rootCmd
}

func newRootCommand(optFns ...func(o *options)) *xcommand {
	opts := newOptions(optFns...)

	cmd := newCommand(&cobra.Command{
		Use:     "deltascii",
		Short:   "ΔSCII",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			return nil
		},
		SilenceUsage: true,
	})

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.err)

	cmd.InitDefaultVersionFlag()
	cmd.InitDefaultCompletionCmd()

	return cmd
}

type deltaFlags struct {
	input  string
	output string
}

func newDeltaCommand(optFns ...func(o *options)) *xcommand {
	opts := newOptions(optFns...)

	flags := new(deltaFlags)

	cmd := newCommand(&cobra.Command{
		Use:     "Δ",
		Aliases: []string{"delta"},
		Short:   "ΔSCII(n) = ASCII(n) - ASCII(n-1)",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var r io.Reader
			if flags.input == "-" {
				r = cmd.InOrStdin()
			} else {
				data, err := os.ReadFile(flags.input)
				if err != nil {
					return err
				}

				r = bytes.NewReader(data)
			}

			buf := new(bytes.Buffer)
			if err := convertASCIICast(r, buf, deltaFn); err != nil {
				return err
			}

			if flags.output == "-" {
				if _, err := fmt.Fprint(cmd.OutOrStdout(), buf.String()); err != nil {
					return err
				}
			} else {
				if err := os.WriteFile(flags.output, buf.Bytes(), 0o644); err != nil {
					return err
				}
			}

			return nil
		},
		SilenceUsage: true,
	})

	cmd.Flags().StringVarP(&flags.input, "input", "i", "", `input asciicast v2 file or "-" (read from stdin)`)
	_ = cmd.MarkFlagRequired("input")

	cmd.Flags().StringVarP(&flags.output, "output", "o", "", `output Δ-asciicast v2 file or "-" (write to stdout)`)
	_ = cmd.MarkFlagRequired("output")

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.out)
	cmd.SetErr(opts.stdio.err)

	return cmd
}

type accumulateFlags struct {
	input  string
	output string
}

func newAccumulateCommand(optFns ...func(o *options)) *xcommand {
	opts := newOptions(optFns...)

	flags := new(accumulateFlags)

	cmd := newCommand(&cobra.Command{
		Use:     "Σ",
		Aliases: []string{"accumulate"},
		Short:   "ASCII(n) = ΣΔSCII(n)",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var r io.Reader
			if flags.input == "-" {
				r = cmd.InOrStdin()
			} else {
				data, err := os.ReadFile(flags.input)
				if err != nil {
					return err
				}

				r = bytes.NewReader(data)
			}

			buf := new(bytes.Buffer)
			if err := convertASCIICast(r, buf, accumulateFn); err != nil {
				return err
			}

			if flags.output == "-" {
				if _, err := fmt.Fprint(cmd.OutOrStdout(), buf.String()); err != nil {
					return err
				}
			} else {
				if err := os.WriteFile(flags.output, buf.Bytes(), 0o644); err != nil {
					return err
				}
			}

			return nil
		},
		SilenceUsage: true,
	})

	cmd.Flags().StringVarP(&flags.input, "input", "i", "", `input Δ-asciicast v2 file or "-" (read from stdin)`)
	_ = cmd.MarkFlagRequired("input")

	cmd.Flags().StringVarP(&flags.output, "output", "o", "", `output asciicast v2 file or "-" (write to stdout)`)
	_ = cmd.MarkFlagRequired("output")

	cmd.SetIn(opts.stdio.in)
	cmd.SetOutput(opts.stdio.out)
	cmd.SetErr(opts.stdio.err)

	return cmd
}

type calcFn func(acc, val float64) (newAcc, newVal float64)

var (
	deltaFn calcFn = func(acc, val float64) (newAcc, newVal float64) {
		delta := decimal.NewFromFloat(val).Sub(decimal.NewFromFloat(acc)).InexactFloat64()
		return val, delta
	}

	accumulateFn calcFn = func(acc, val float64) (newAcc, newVal float64) {
		sum := decimal.NewFromFloat(val).Add(decimal.NewFromFloat(acc)).InexactFloat64()
		return sum, sum
	}
)

func convertASCIICast(r io.Reader, w io.Writer, fn calcFn) error {
	dec := json.NewDecoder(r)
	enc := json.NewEncoder(w)

	var h asciinema.V2Header
	if err := dec.Decode(&h); err != nil {
		return err
	}

	if err := enc.Encode(&h); err != nil {
		return err
	}

	acc := 0.0
	for dec.More() {
		var e asciinema.V2Event
		if err := dec.Decode(&e); err != nil {
			return err
		}

		acc, e.Time = fn(acc, e.Time)

		if err := enc.Encode(&e); err != nil {
			return err
		}
	}

	return nil
}
