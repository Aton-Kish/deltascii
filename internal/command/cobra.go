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
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"text/template"
	"time"

	"github.com/Aton-Kish/deltascii/internal/xfilepath"
	"github.com/spf13/cobra"
)

const (
	fileNameReadme = "README.md"
)

var (
	lastUpdatedPattern = regexp.MustCompile(`Last updated on \d{4}-\d{2}-\d{2}`)

	templateFuncMap = template.FuncMap{
		"replace": strings.ReplaceAll,
		"now":     time.Now,
	}

	//go:embed template/README.md.gotmpl
	readmeGoTemplate string
	readmeTemplate   = template.Must(
		template.
			New("readme").
			Funcs(templateFuncMap).
			Parse(readmeGoTemplate),
	)

	//go:embed template/reference.md.gotmpl
	referenceGoTemplate string
	referenceTemplate   = template.Must(
		template.
			New("reference").
			Funcs(templateFuncMap).
			Parse(referenceGoTemplate),
	)
)

type xcommand struct {
	*cobra.Command
}

func newCommand(cmd *cobra.Command) *xcommand {
	return &xcommand{
		Command: cmd,
	}
}

func (c *xcommand) AliasUseLines() []string {
	line := c.UseLine()
	lines := make([]string, 0, len(c.Aliases))
	for _, alias := range c.Aliases {
		var aliasLine string
		if c.HasParent() {
			aliasLine = strings.Replace(line, fmt.Sprintf("%s %s", c.Parent().CommandPath(), c.Name()), fmt.Sprintf("%s %s", c.Parent().CommandPath(), alias), 1)
		} else {
			aliasLine = strings.Replace(line, c.Name(), alias, 1)
		}

		lines = append(lines, aliasLine)
	}

	return lines
}

func (c *xcommand) GenerateReadme(dir string) error {
	return c.generateDocument(readmeTemplate, dir, fileNameReadme)
}

func (c *xcommand) GenerateReference(dir string) error {
	c.InitDefaultHelpCmd()
	c.InitDefaultHelpFlag()

	return c.generateDocument(referenceTemplate, dir, fmt.Sprintf("%s.md", strings.ReplaceAll(c.CommandPath(), " ", "-")))
}

func (c *xcommand) generateDocument(tmpl *template.Template, dir string, fileName string) error {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, c); err != nil {
		return err
	}

	name := filepath.Join(dir, fileName)
	if xfilepath.Exist(name) {
		data, err := os.ReadFile(name)
		if err != nil {
			return err
		}

		mask := "Last updated on 2006-01-02"
		if slices.Equal(lastUpdatedPattern.ReplaceAll(data, []byte(mask)), lastUpdatedPattern.ReplaceAll(buf.Bytes(), []byte(mask))) {
			// NOTE: no need to generate
			return nil
		}
	}

	if !xfilepath.Exist(dir) {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	if err := os.WriteFile(name, buf.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}

func (c *xcommand) GenerateReferences(dir string) error {
	if err := c.GenerateReference(dir); err != nil {
		return err
	}

	for _, sub := range c.Commands() {
		if !sub.IsAvailableCommand() || sub.IsAdditionalHelpTopicCommand() {
			continue
		}

		if err := newCommand(sub).GenerateReferences(dir); err != nil {
			return err
		}
	}

	return nil
}
