// MIT License

// Copyright (c) 2022 Project-Sparrow
// Copyright (c) 2023 Kelwing <kelwing@kelnet.org>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package snowflake_test

import (
	"testing"
	"time"

	"wumpgo.dev/snowflake"
)

func TestGenerate(t *testing.T) {
	epoch := time.Now()
	snowflake.Init(epoch, 1, 1)

	s := snowflake.Generate()

	if (int(s)&0x3E0000)>>17 != 1 {
		t.Fail()
	}

	if (int(s)&0x1F000)>>12 != 1 {
		t.Fail()
	}
}
