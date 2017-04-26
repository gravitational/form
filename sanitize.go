/*
Copyright 2017 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package form

import (
	"unicode"

	"github.com/gravitational/trace"
)

// AllowSet is an interface that exposes a single function used to check if
// input is safe.
type AllowSet interface {
	IsSafe(string) error
}

// CharAllowSet is a set of allowed ASCII characters and a maximum length.
type CharAllowSet struct {
	len   int
	chars [256]bool
}

// NewCharAllowSet creates a new CharAllowSet.
func NewCharAllowSet(s string, n int) (CharAllowSet, error) {
	cas := CharAllowSet{len: n}
	for _, v := range s {
		if v > unicode.MaxASCII {
			return cas, trace.BadParameter("%v [%#x] is outside ASCII range", string(v), v)
		}
		cas.chars[int(v)] = true
	}
	return cas, nil
}

// IsSafe checks against CharAllowSet rules.
func (c CharAllowSet) IsSafe(s string) error {
	if len(s) > c.len {
		return trace.BadParameter("length %v is longer than maximum allowable length: %v", len(s), c.len)
	}

	for _, v := range s {
		if v > unicode.MaxASCII {
			return trace.BadParameter("non-ascii character %v [%#x] not allowed", string(v), v)
		}
		if c.chars[int(v)] == false {
			return trace.BadParameter("character %v [%#x] not allowed", string(v), v)
		}
	}

	return nil
}
