// Copyright 2024 slw <150657601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package genutil

import (
	"testing"
)

func TestValidDir(t *testing.T) {
	_, err := OutDir("./")
	if err != nil {
		t.Fatal(err)
	}
}

func TestInvalidDir(t *testing.T) {
	_, err := OutDir("./nondir")
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestNotDir(t *testing.T) {
	_, err := OutDir("./genutils_test.go")
	if err == nil {
		t.Fatal("expected an error")
	}
}
