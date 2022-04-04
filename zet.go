package zet

import (
	"fmt"
	"log"
	"time"

	"github.com/hbjydev/zet/pkg/repo"
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/inc/help"
	"github.com/rwxrob/config"
)

var Cmd = &Z.Cmd{
	Name:    `zet`,
	Version: `v0.0.1`,
	Summary: `manages a zettelkasten`,
	Usage:   `[init|new|view|edit|list|help|config]`,

	Commands: []*Z.Cmd{_init, new, help.Cmd, config.Cmd},
}

var _init = &Z.Cmd{
	Name:    `init`,
	Summary: `initializes a new zettelkasten`,

	Call: func(caller *Z.Cmd, args ...string) error {
		return repo.Init()
	},
}

var new = &Z.Cmd{
	Name: `new`,
	Call: func(caller *Z.Cmd, args ...string) error {
		now := time.Now()
		id := fmt.Sprintf(
			"%v%v%v%v%v%v",
			fmt.Sprintf("%04d", now.Year()),
			fmt.Sprintf("%02d", int(now.Month())),
			fmt.Sprintf("%02d", now.Day()),
			fmt.Sprintf("%02d", now.Hour()),
			fmt.Sprintf("%02d", now.Minute()),
			fmt.Sprintf("%02d", now.Second()),
		)

		if len(id) != 14 {
			return fmt.Errorf("invalid id length: %v", len(id))
		}

		log.Printf("Create a new Zettelkasten: %v\n", id)

		exists, err := repo.Exists(id)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("zettelkasten with id [%v] already exists", id)
		}

		if err := repo.New(id); err != nil {
			return err
		}

		return nil
	},
}
