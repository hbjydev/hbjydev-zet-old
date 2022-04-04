package zet

import (
	"fmt"
	"log"
	"time"

	"github.com/hbjydev/zet/pkg/repo"
	Z "github.com/rwxrob/bonzai"
)

var new = &Z.Cmd{
	Name: `new`,
	Call: func(caller *Z.Cmd, args ...string) error {
		zetExists, err := repo.RepoExists()
		if err != nil {
			return fmt.Errorf("no zettelkasten is initialized in the path specified. have you run `zet init`?")
		} else if !zetExists {
			return fmt.Errorf("no zettelkasten is initialized in the path specified. have you run `zet init`?")
		}

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

		log.Printf("created a new zet: %v\n", id)

		return nil
	},
}
