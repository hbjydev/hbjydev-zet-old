package zet

import (
	"fmt"

	"github.com/hbjydev/zet/pkg/repo"
	Z "github.com/rwxrob/bonzai"
)

var list = &Z.Cmd{
	Name: `list`,

	Call: func(caller *Z.Cmd, args ...string) error {
		items, err := repo.List()
		if err != nil {
			return err
		}

		fmt.Println("=============================")
		for i := 0; i < len(items); i++ {
			fmt.Printf("|- %v\n", items[i])
		}

		return nil
	},
}
