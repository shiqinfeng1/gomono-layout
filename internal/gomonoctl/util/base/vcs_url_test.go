// Copyright 2024 Shiqinfeng &lt;150627601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package base

import (
	"net"
	"strings"
	"testing"
)

func TestParseVCSUrl(t *testing.T) {
	repos := []string{
		// ssh://[user@]host.xz[:port]/path/to/repo.git/
		"ssh://git@github.com:7875/shiqinfeng1/gomono-layout.git",
		// git://host.xz[:port]/path/to/repo.git/
		"git://github.com:7875/shiqinfeng1/gomono-layout.git",
		// http[s]://host.xz[:port]/path/to/repo.git/
		"https://github.com:7875/shiqinfeng1/gomono-layout.git",
		// ftp[s]://host.xz[:port]/path/to/repo.git/
		"ftps://github.com:7875/shiqinfeng1/gomono-layout.git",
		//[user@]host.xz:path/to/repo.git/
		"git@github.com:shiqinfeng1/gomono-layout.git",
		// ssh://[user@]host.xz[:port]/~[user]/path/to/repo.git/
		"ssh://git@github.com:7875/shiqinfeng1/gomono-layout.git",
		// git://host.xz[:port]/~[user]/path/to/repo.git/
		"git://github.com:7875/shiqinfeng1/gomono-layout.git",
		//[user@]host.xz:/~[user]/path/to/repo.git/
		"git@github.com:shiqinfeng1/gomono-layout.git",
		///path/to/repo.git/
		"~/shiqinfeng1/gomono-layout.git",
		// file:///path/to/repo.git/
		"file://~/shiqinfeng1/gomono-layout.git",
	}
	for _, repo := range repos {
		url, err := ParseVCSUrl(repo)
		if err != nil {
			t.Fatal(repo, err)
		}
		urlPath := strings.TrimLeft(url.Path, "/")
		if urlPath != "shiqinfeng1/gomono-layout.git" {
			t.Fatal(repo, "parse url failed", urlPath)
		}
	}
}

func TestParseSsh(t *testing.T) {
	repo := "ssh://git@github.com:7875/shiqinfeng1/gomono-layout.git"
	url, err := ParseVCSUrl(repo)
	if err != nil {
		t.Fatal(err)
	}
	host, _, err := net.SplitHostPort(url.Host)
	if err != nil {
		host = url.Host
	}
	t.Log(host, url.Path)
}
