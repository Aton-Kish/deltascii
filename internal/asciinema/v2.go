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
)

type V2Header struct {
	// required
	Version int `json:"version"`
	Width   int `json:"width"`
	Height  int `json:"height"`
	// optional
	Timestamp     int               `json:"timestamp,omitempty"`
	Duration      float64           `json:"duration,omitempty"`
	IdleTimeLimit float64           `json:"idle_time_limit,omitempty"`
	Command       string            `json:"command,omitempty"`
	Title         string            `json:"title,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	Theme         *V2HeaderTheme    `json:"theme,omitempty"`
}

type V2HeaderTheme struct {
	FG      string `json:"fg"`
	BG      string `json:"bg"`
	Palette string `json:"palette"`
}

type V2Event struct {
	Time float64 `json:"time"`
	Code string  `json:"code"`
	Data any     `json:"data"`
}

func (e *V2Event) UnmarshalJSON(b []byte) error {
	var v [3]any
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	t, ok := v[0].(float64)
	if !ok {
		return fmt.Errorf("invalid event time: %v", v[0])
	}

	c, ok := v[1].(string)
	if !ok {
		return fmt.Errorf("invalid event code: %v", v[1])
	}

	e.Time = t
	e.Code = c
	e.Data = v[2]

	return nil
}

func (e V2Event) MarshalJSON() ([]byte, error) {
	return json.Marshal([3]any{e.Time, e.Code, e.Data})
}
