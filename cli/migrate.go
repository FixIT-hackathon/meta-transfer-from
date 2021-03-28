package cli

import (
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/assets"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/config"
)

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func MigrateUp(cfg config.Config) error {
	am, err := migrate.Exec(cfg.DB().RawDB(), "postgres", migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}

	fmt.Println("Applied", am)
	return nil
}

func MigrateDown(cfg config.Config) error {
	am, err := migrate.Exec(cfg.DB().RawDB(), "postgres", migrations, migrate.Down)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	fmt.Println("Down", am)

	return nil
}
