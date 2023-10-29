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
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeltaCommand(t *testing.T) {
	deltacast := []byte(`{"version":2,"width":80,"height":24,"timestamp":1504467315,"env":{"SHELL":"/bin/zsh","TERM":"xterm-256color"}}
[0,"o","h"]
[0.1,"o","e"]
[0.2,"o","l"]
[0.3,"o","l"]
[0.4,"o","o"]
[0.5,"o"," "]
[0.6,"o","w"]
[0.7,"o","o"]
[0.8,"o","r"]
[0.9,"o","l"]
[1,"o","d"]
`)

	type args struct {
		input  string
		output string
	}

	type expected struct {
		data  []byte
		errIs error
	}

	tests := []struct {
		name     string
		args     *args
		expected *expected
	}{
		{
			name: "happy path: input from file / output to file",
			args: &args{
				input:  "testdata/test.cast",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  deltacast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from stdin / output to file",
			args: &args{
				input:  "-",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  deltacast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from file / output to stdout",
			args: &args{
				input:  "testdata/test.cast",
				output: "-",
			},
			expected: &expected{
				data:  deltacast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from stdin / output to stdout",
			args: &args{
				input:  "-",
				output: "-",
			},
			expected: &expected{
				data:  deltacast,
				errIs: nil,
			},
		},
		{
			name: "edge path: input not exist",
			args: &args{
				input:  "testdata/not-exist/test.cast",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  nil,
				errIs: os.ErrNotExist,
			},
		},
		{
			name: "edge path: output not exist",
			args: &args{
				input:  "testdata/test.cast",
				output: filepath.Join(t.TempDir(), "not-exist/output.cast"),
			},
			expected: &expected{
				data:  nil,
				errIs: os.ErrNotExist,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()

			var stdin io.Reader
			if tt.args.input == "-" {
				data, _ := os.ReadFile("testdata/test.cast")
				stdin = bytes.NewReader(data)
			} else {
				stdin = new(bytes.Reader)
			}
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			cmd := newDeltaCommand(WithStdio(stdin, stdout, stderr))
			cmd.SetArgs([]string{"--input", tt.args.input, "--output", tt.args.output})

			// Act
			err := cmd.ExecuteContext(ctx)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				if tt.args.output == "-" {
					assert.Equal(t, tt.expected.data, stdout.Bytes())
				} else {
					data, _ := os.ReadFile(tt.args.output)
					assert.Equal(t, tt.expected.data, data)
				}
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.expected.errIs)
			}
		})
	}
}

func TestAccumulateCommand(t *testing.T) {
	acccast := []byte(`{"version":2,"width":80,"height":24,"timestamp":1504467315,"env":{"SHELL":"/bin/zsh","TERM":"xterm-256color"}}
[0,"o","h"]
[0.1,"o","e"]
[0.4,"o","l"]
[1,"o","l"]
[2,"o","o"]
[3.5,"o"," "]
[5.6,"o","w"]
[8.4,"o","o"]
[12,"o","r"]
[16.5,"o","l"]
[22,"o","d"]
`)

	type args struct {
		input  string
		output string
	}

	type expected struct {
		data  []byte
		errIs error
	}

	tests := []struct {
		name     string
		args     *args
		expected *expected
	}{
		{
			name: "happy path: input from file / output to file",
			args: &args{
				input:  "testdata/test.cast",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  acccast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from stdin / output to file",
			args: &args{
				input:  "-",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  acccast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from file / output to stdout",
			args: &args{
				input:  "testdata/test.cast",
				output: "-",
			},
			expected: &expected{
				data:  acccast,
				errIs: nil,
			},
		},
		{
			name: "happy path: input from stdin / output to stdout",
			args: &args{
				input:  "-",
				output: "-",
			},
			expected: &expected{
				data:  acccast,
				errIs: nil,
			},
		},
		{
			name: "edge path: input not exist",
			args: &args{
				input:  "testdata/not-exist/test.cast",
				output: filepath.Join(t.TempDir(), "output.cast"),
			},
			expected: &expected{
				data:  nil,
				errIs: os.ErrNotExist,
			},
		},
		{
			name: "edge path: output not exist",
			args: &args{
				input:  "testdata/test.cast",
				output: filepath.Join(t.TempDir(), "not-exist/output.cast"),
			},
			expected: &expected{
				data:  nil,
				errIs: os.ErrNotExist,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()

			var stdin io.Reader
			if tt.args.input == "-" {
				data, _ := os.ReadFile("testdata/test.cast")
				stdin = bytes.NewReader(data)
			} else {
				stdin = new(bytes.Reader)
			}
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			cmd := newAccumulateCommand(WithStdio(stdin, stdout, stderr))
			cmd.SetArgs([]string{"--input", tt.args.input, "--output", tt.args.output})

			// Act
			err := cmd.ExecuteContext(ctx)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				if tt.args.output == "-" {
					assert.Equal(t, tt.expected.data, stdout.Bytes())
				} else {
					data, _ := os.ReadFile(tt.args.output)
					assert.Equal(t, tt.expected.data, data)
				}
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.expected.errIs)
			}
		})
	}
}
