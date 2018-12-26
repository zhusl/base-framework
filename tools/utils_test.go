// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package tools

import "testing"

func TestStringInSlice(t *testing.T) {
	testSlice := []string{"hello", "world", "china"}
	if !StringInSlice(testSlice, "hello") {
		t.Error("function StringInSlice failed to achieve the expected")
	}
	if StringInSlice(testSlice, "hello1") {
		t.Error("function StringInSlice failed to achieve the expected")
	}
}
