package zet

import (
	"github.com/hbjydev/zet/pkg/repo"
	Z "github.com/rwxrob/bonzai"
)

var _init = &Z.Cmd{
	Name:    `init`,
	Summary: `initializes a new zettelkasten`,

	Call: func(caller *Z.Cmd, args ...string) error {
		return repo.Init()
	},
}
