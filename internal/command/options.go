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
	"io"
	"os"
)

type stdio struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

type options struct {
	stdio stdio
}

func newOptions(optFns ...func(o *options)) *options {
	o := &options{
		stdio: stdio{in: os.Stdin, out: os.Stdout, err: os.Stderr},
	}

	for _, fn := range optFns {
		fn(o)
	}

	return o
}

func WithStdio(in io.Reader, out, err io.Writer) func(o *options) {
	return func(o *options) {
		o.stdio = stdio{in: in, out: out, err: err}
	}
}
