package cli

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/config"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service"
	"github.com/urfave/cli"
	"gitlab.com/distributed_lab/kit/kv"
)

func Run(args []string) bool {
	var cfg config.Config

	app := cli.NewApp()

	before := func(_ *cli.Context) error {
		getter, err := kv.FromEnv()
		if err != nil {
			return err
		}
		cfg = config.NewConfig(getter)
		return nil
	}

	app.Commands = cli.Commands{
		{
			Name: "migrate",
			Subcommands: cli.Commands{
				{
					Name:   "up",
					Before: before,
					Action: func(ctx *cli.Context) error {
						return MigrateUp(cfg)
					},
				},
				{
					Name:   "down",
					Before: before,
					Action: func(ctx *cli.Context) error {
						return MigrateDown(cfg)
					},
				},
			},
		},
		{
			Name: "run",
			Action: func(_ *cli.Context) error {
				return service.Service{}.Run(cfg)
			},
		},
	}

	if err := app.Run(args); err != nil {
		return false
	}
	return true
}
