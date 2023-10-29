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
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_xcommand_GenerateReadme(t *testing.T) {
	type fields struct {
		Command *cobra.Command
	}

	type args struct {
		dir string
	}

	type expected struct {
		content string
		errIs   error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "happy path: has children",
			fields: fields{
				Command: func() *cobra.Command {
					cmd := &cobra.Command{
						Use:   "Use",
						Short: "Short",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}

					sub := &cobra.Command{
						Use:   "SubUse",
						Short: "SubShort",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}

					cmd.AddCommand(sub)

					return cmd
				}(),
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				content: `# Command reference

<sub><sup>Last updated on ` + time.Now().Format("2006-01-02") + `</sup></sub>

- [Use](Use.md) - Short
- [Use SubUse](Use-SubUse.md) - SubShort
`,
				errIs: nil,
			},
		},
		{
			name: "happy path: orphan",
			fields: fields{
				Command: &cobra.Command{
					Use:   "Use",
					Short: "Short",
				},
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				content: `# Command reference

<sub><sup>Last updated on ` + time.Now().Format("2006-01-02") + `</sup></sub>

- [Use](Use.md) - Short
`,
				errIs: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			c := &xcommand{
				Command: tt.fields.Command,
			}

			// Act
			err := c.GenerateReadme(tt.args.dir)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.NoError(t, err)

				data, err := os.ReadFile(filepath.Join(tt.args.dir, fileNameReadme))
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.content, string(data))
			} else {
				assert.Error(t, err)

				if tt.expected.errIs != nil {
					assert.ErrorIs(t, err, tt.expected.errIs)
				}
			}
		})
	}
}

func Test_xcommand_GenerateReference(t *testing.T) {
	type fields struct {
		Command *cobra.Command
	}

	type args struct {
		dir string
	}

	type expected struct {
		name    string
		content string
		errIs   error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "happy path: long",
			fields: fields{
				Command: func() *cobra.Command {
					root := &cobra.Command{
						Use:   "RootUse",
						Short: "RootShort",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}
					root.PersistentFlags().String("PersistentFlag", "", "Usage of PersistentFlag")

					cmd := &cobra.Command{
						Use:     "Use",
						Aliases: []string{"Alias1", "Alias2"},
						Short:   "Short",
						Long: `Long
`,
						Example: "Example",
						Version: "Version",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}
					cmd.Flags().String("Flag1", "", "Usage of Flag1")
					cmd.Flags().String("Flag2", "DefaultValueOfFlag2", "Usage of Flag2")

					sub := &cobra.Command{
						Use:   "SubUse",
						Short: "SubShort",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}

					root.AddCommand(cmd)
					cmd.AddCommand(sub)

					return cmd
				}(),
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				name: "RootUse-Use.md",
				content: `## ` + "`RootUse Use`" + `

<sub><sup>Last updated on ` + time.Now().Format("2006-01-02") + `</sup></sub>

Short

### Synopsis

Long


` + "```" + `shell
RootUse Use [flags]
` + "```" + `

### Examples

` + "```" + `shell
Example
` + "```" + `

### Options

` + "```" + `shell
      --Flag1 string   Usage of Flag1
      --Flag2 string   Usage of Flag2 (default "DefaultValueOfFlag2")
  -h, --help           help for Use
` + "```" + `

### Options inherited from parent commands

` + "```" + `shell
      --PersistentFlag string   Usage of PersistentFlag
` + "```" + `

### See also

- [RootUse](RootUse.md) - RootShort
- [RootUse Use SubUse](RootUse-Use-SubUse.md) - SubShort
`,
				errIs: nil,
			},
		},
		{
			name: "happy path: short",
			fields: fields{
				Command: &cobra.Command{
					Use:   "Use",
					Short: "Short",
				},
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				name: "Use.md",
				content: `## ` + "`Use`" + `

<sub><sup>Last updated on ` + time.Now().Format("2006-01-02") + `</sup></sub>

Short

### Options

` + "```" + `shell
  -h, --help   help for Use
` + "```" + `
`,
				errIs: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			c := &xcommand{
				Command: tt.fields.Command,
			}

			// Act
			err := c.GenerateReference(tt.args.dir)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.NoError(t, err)

				data, err := os.ReadFile(filepath.Join(tt.args.dir, tt.expected.name))
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.content, string(data))
			} else {
				assert.Error(t, err)

				if tt.expected.errIs != nil {
					assert.ErrorIs(t, err, tt.expected.errIs)
				}
			}
		})
	}
}

func Test_xcommand_GenerateReferences(t *testing.T) {
	type fields struct {
		Command *cobra.Command
	}

	type args struct {
		dir string
	}

	type expected struct {
		names []string
		errIs error
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "happy path: has a parent and children",
			fields: fields{
				Command: func() *cobra.Command {
					root := &cobra.Command{
						Use:   "RootUse",
						Short: "RootShort",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}
					root.PersistentFlags().String("PersistentFlag", "", "Usage of PersistentFlag")

					cmd := &cobra.Command{
						Use:     "Use",
						Aliases: []string{"Alias1", "Alias2"},
						Short:   "Short",
						Long: `Long
`,
						Example: "Example",
						Version: "Version",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}
					cmd.Flags().String("Flag1", "", "Usage of Flag1")
					cmd.Flags().String("Flag2", "DefaultValueOfFlag2", "Usage of Flag2")

					sub := &cobra.Command{
						Use:   "SubUse",
						Short: "SubShort",
						RunE: func(cmd *cobra.Command, args []string) error {
							return nil
						},
					}

					root.AddCommand(cmd)
					cmd.AddCommand(sub)

					return cmd
				}(),
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				names: []string{
					"RootUse-Use.md",
					"RootUse-Use-SubUse.md",
				},
				errIs: nil,
			},
		},
		{
			name: "happy path: orphan",
			fields: fields{
				Command: &cobra.Command{
					Use:   "Use",
					Short: "Short",
				},
			},
			args: args{
				dir: t.TempDir(),
			},
			expected: expected{
				names: []string{
					"Use.md",
				},
				errIs: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			c := &xcommand{
				Command: tt.fields.Command,
			}

			// Act
			err := c.GenerateReferences(tt.args.dir)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.NoError(t, err)

				es, err := os.ReadDir(tt.args.dir)
				assert.NoError(t, err)
				names := make([]string, 0, len(es))
				for _, e := range es {
					names = append(names, e.Name())
				}
				assert.ElementsMatch(t, tt.expected.names, names)
			} else {
				assert.Error(t, err)

				if tt.expected.errIs != nil {
					assert.ErrorIs(t, err, tt.expected.errIs)
				}
			}
		})
	}
}
