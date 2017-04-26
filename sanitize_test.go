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
	"fmt"

	"gopkg.in/check.v1"
)

type AllowSetSuite struct{}

var _ = check.Suite(&AllowSetSuite{})
var _ = fmt.Printf

func (s *AllowSetSuite) SetUpSuite(c *check.C)    {}
func (s *AllowSetSuite) TearDownSuite(c *check.C) {}
func (s *AllowSetSuite) SetUpTest(c *check.C)     {}
func (s *AllowSetSuite) TearDownTest(c *check.C)  {}

func (s *AllowSetSuite) TestCharAllowSet(c *check.C) {
	var tests = []struct {
		inChars           string
		inLen             int
		inInput           string
		outCreateAllowSet bool
		outIsSafe         bool
	}{
		// 0 - length too long
		{
			"a",
			1,
			"aa",
			true,
			false,
		},
		// 1 - invalid chars
		{
			"a",
			2,
			"ab",
			true,
			false,
		},
		// 2 - non-ascii allowset
		{
			"😎",
			0,
			"",
			false,
			false,
		},
		// 3 - non-ascii input
		{
			"a",
			4,
			"😎",
			true,
			false,
		},
		// 4 - all good
		{
			"a",
			2,
			"aa",
			true,
			true,
		},
	}

	for i, tt := range tests {
		comment := check.Commentf("Test %v", i)

		as, err := NewCharAllowSet(tt.inChars, tt.inLen)
		if tt.outCreateAllowSet {
			c.Assert(err, check.IsNil, comment)
		} else {
			c.Assert(err, check.NotNil, comment)
			continue
		}
		err = as.IsSafe(tt.inInput)
		if tt.outIsSafe {
			c.Assert(err, check.IsNil, comment)
		} else {
			c.Assert(err, check.NotNil, comment)
		}
	}
}
