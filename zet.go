package zet

import (
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/config"
)

var Cmd = &Z.Cmd{
	Name:    `zet`,
	Version: `v0.0.1`,
	Summary: `manages a zettelkasten`,

	Usage: `[init|new|view|edit|list|help|config]`,

	License:   `apache-2.0`,
	Copyright: `Copyright 2022 Hayden Young`,

	MinArgs: 1,

	Commands: []*Z.Cmd{_init, new, list, help.Cmd, config.Cmd},
}
