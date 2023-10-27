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

package asciinema

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV2Header_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}

	type expected struct {
		data *V2Header
		err  error
	}

	tests := []struct {
		name     string
		args     *args
		expected *expected
	}{
		{
			name: "happy path: required",
			args: &args{
				b: []byte(`{"version": 2, "width": 80, "height": 24}`),
			},
			expected: &expected{
				data: &V2Header{
					Version: 2,
					Width:   80,
					Height:  24,
				},
				err: nil,
			},
		},
		{
			name: "happy path: optional",
			args: &args{
				b: []byte(`{"version": 2, "width": 80, "height": 24, "timestamp": 1504467315, "duration": 1.23, "idle_time_limit": 4.56, "command": "Command", "title": "Demo", "env": {"SHELL": "/bin/zsh", "TERM": "xterm-256color"}, "theme": {"fg": "#d0d0d0", "bg": "#212121", "palette": "#151515:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#d0d0d0:#505050:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#f5f5f5"}}`),
			},
			expected: &expected{
				data: &V2Header{
					Version:       2,
					Width:         80,
					Height:        24,
					Timestamp:     1504467315,
					Duration:      1.23,
					IdleTimeLimit: 4.56,
					Command:       "Command",
					Title:         "Demo",
					Env: map[string]string{
						"SHELL": "/bin/zsh",
						"TERM":  "xterm-256color",
					},
					Theme: &V2HeaderTheme{
						FG:      "#d0d0d0",
						BG:      "#212121",
						Palette: "#151515:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#d0d0d0:#505050:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#f5f5f5",
					},
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			var v V2Header
			err := json.Unmarshal(tt.args.b, &v)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.Equal(t, tt.expected.data, &v)
				assert.NoError(t, err)
			} else {
				assert.Zero(t, v)
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestV2Header_MarshalJSON(t *testing.T) {
	type expected struct {
		data []byte
		err  error
	}

	tests := []struct {
		name     string
		data     *V2Header
		expected *expected
	}{
		{
			name: "happy path: required",
			data: &V2Header{
				Version: 2,
				Width:   80,
				Height:  24,
			},
			expected: &expected{
				data: []byte(`{"version":2,"width":80,"height":24}`),
				err:  nil,
			},
		},
		{
			name: "happy path: optional",
			data: &V2Header{
				Version:       2,
				Width:         80,
				Height:        24,
				Timestamp:     1504467315,
				Duration:      1.23,
				IdleTimeLimit: 4.56,
				Command:       "Command",
				Title:         "Demo",
				Env: map[string]string{
					"SHELL": "/bin/zsh",
					"TERM":  "xterm-256color",
				},
				Theme: &V2HeaderTheme{
					FG:      "#d0d0d0",
					BG:      "#212121",
					Palette: "#151515:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#d0d0d0:#505050:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#f5f5f5",
				},
			},
			expected: &expected{
				data: []byte(`{"version":2,"width":80,"height":24,"timestamp":1504467315,"duration":1.23,"idle_time_limit":4.56,"command":"Command","title":"Demo","env":{"SHELL":"/bin/zsh","TERM":"xterm-256color"},"theme":{"fg":"#d0d0d0","bg":"#212121","palette":"#151515:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#d0d0d0:#505050:#ac4142:#7e8e50:#e5b567:#6c99bb:#9f4e85:#7dd6cf:#f5f5f5"}}`),
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual, err := json.Marshal(tt.data)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.Equal(t, tt.expected.data, actual)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, actual)
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestV2Event_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}

	type expected struct {
		data *V2Event
		err  error
	}

	tests := []struct {
		name     string
		args     *args
		expected *expected
	}{
		{
			name: "happy path",
			args: &args{
				b: []byte(`[0.123456789, "o", "hello world"]`),
			},
			expected: &expected{
				data: &V2Event{
					Time: 0.123456789,
					Code: "o",
					Data: "hello world",
				},
				err: nil,
			},
		},
		{
			name: "edge path: invalid event time",
			args: &args{
				b: []byte(`["0.123456789", "o", "hello world"]`),
			},
			expected: &expected{
				data: nil,
				err:  fmt.Errorf("invalid event time: %v", "0.123456789"),
			},
		},
		{
			name: "edge path: invalid event code",
			args: &args{
				b: []byte(`[0.123456789, 0, "hello world"]`),
			},
			expected: &expected{
				data: nil,
				err:  fmt.Errorf("invalid event code: %v", 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			var v V2Event
			err := json.Unmarshal(tt.args.b, &v)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.Equal(t, tt.expected.data, &v)
				assert.NoError(t, err)
			} else {
				assert.Zero(t, v)
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}

func TestV2Event_MarshalJSON(t *testing.T) {
	type expected struct {
		data []byte
		err  error
	}

	tests := []struct {
		name     string
		data     *V2Event
		expected *expected
	}{
		{
			name: "happy path",
			data: &V2Event{
				Time: 0.123456789,
				Code: "o",
				Data: "hello world",
			},
			expected: &expected{
				data: []byte(`[0.123456789,"o","hello world"]`),
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			actual, err := json.Marshal(tt.data)

			// Assert
			if strings.HasPrefix(tt.name, "happy") {
				assert.Equal(t, tt.expected.data, actual)
				assert.NoError(t, err)
			} else {
				assert.Nil(t, actual)
				assert.Equal(t, tt.expected.err, err)
			}
		})
	}
}
