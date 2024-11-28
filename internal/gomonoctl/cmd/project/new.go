// Copyright 2024 Shiqinfeng &lt;150627601@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/shiqinfeng1/gomono-layout/internal/gomonoctl/util/base"
)

// Project is a project template.
type Project struct {
	Name string
}

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir, layout, branch, serviceName string) error {
	to := filepath.Join(dir, p.Name)

	fmt.Printf(
		"🚀 Creating project %s & Add service %s, layout repo is %s, please wait a moment.\n\n",
		p.Name,
		serviceName,
		layout,
	)
	repo := base.NewRepo(layout, branch, []string{"cmd/server", "internal/server"})
	if err := repo.CopyTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
		return err
	}

	e := os.Rename(
		filepath.Join(to, "cmd", "server"),
		filepath.Join(to, "cmd", serviceName),
	)
	if e != nil {
		return e
	}
	e = os.Rename(
		filepath.Join(to, "internal", "server"),
		filepath.Join(to, "internal", serviceName),
	)
	if e != nil {
		return e
	}
	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			🤝 Thanks for using Gomono")
	// fmt.Println("	📚 Tutorial: https://go-kratos.dev/docs/getting-started/start")
	return nil
}

// New new a project from remote repo.
func (p *Project) AddService(ctx context.Context, dir, layout, branch, serviceName string) error {
	to := filepath.Join(dir, p.Name)

	_, err1 := os.Stat(filepath.Join(to, "cmd", serviceName))
	_, err2 := os.Stat(filepath.Join(to, "internal", serviceName))
	if !os.IsNotExist(err1) || !os.IsNotExist(err2) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return nil
		}
		os.RemoveAll(filepath.Join(to, "cmd", serviceName))
		os.RemoveAll(filepath.Join(to, "internal", serviceName))
	}

	fmt.Printf("🚀 Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch, []string{"cmd/server", "internal/server"})
	if err := repo.CopyServiceTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
		return err
	}

	e := os.Rename(
		filepath.Join(to, "cmd", "server"),
		filepath.Join(to, "cmd", serviceName),
	)
	if e != nil {
		return e
	}
	e = os.Rename(
		filepath.Join(to, "internal", "server"),
		filepath.Join(to, "internal", serviceName),
	)
	if e != nil {
		return e
	}
	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			🤝 Thanks for using Gomono")
	// fmt.Println("	📚 Tutorial: https://go-kratos.dev/docs/getting-started/start")
	return nil
}
