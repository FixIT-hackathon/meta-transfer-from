package cli

import (
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service"
	"github.com/urfave/cli"
)

func Run(args []string) bool {
	//var cfg config.Config
	//log := logan.New()
	//
	//defer func() {
	//	if rvr := recover(); rvr != nil {
	//		log.WithRecover(rvr).Error("app panicked")
	//	}
	//}()
	//
	app := cli.NewApp()
	//
	//before := func(_ *cli.Context) error {
	//	getter, err := kv.FromEnv()
	//	if err != nil {
	//		return errors.Wrap(err, "failed to get config")
	//	}
	//	cfg = config.NewConfig(getter)
	//	log = cfg.Log()
	//	return nil
	//}

	app.Commands = cli.Commands{
		//{
		//	Name: "migrate",
		//	Subcommands: cli.Commands{
		//		{
		//			Name:   "up",
		//			Before: before,
		//			Action: func(ctx *cli.Context) error {
		//				return MigrateUp(cfg)
		//			},
		//		},
		//		{
		//			Name:   "down",
		//			Before: before,
		//			Action: func(ctx *cli.Context) error {
		//				return MigrateDown(cfg)
		//			},
		//		},
		//	},
		//},
		{
			Name:   "run",
			Action: func(_ *cli.Context) error {
				return service.Service{}.Run()
			},
		},
	}

	if err := app.Run(args); err != nil {
		return false
	}
	return true
}
