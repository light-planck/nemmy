package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/light-planck/nemmy/backend/internal"
	"github.com/light-planck/nemmy/backend/internal/migrations"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.Background()
	cfg, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}
	db, err := internal.NewDB(ctx, cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))

	cmds := newCommands(migrate.NewMigrator(db, migrations.Migrations))
	app := &cli.App{Commands: cmds}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newCommands(migrator *migrate.Migrator) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "init",
			Usage: "create migration tables",
			Action: func(c *cli.Context) error {
				return migrator.Init(c.Context)
			},
		},
		{
			Name:  "migrate",
			Usage: "migrate database",
			Action: func(c *cli.Context) error {
				if err := migrator.Lock(c.Context); err != nil {
					return err
				}
				defer migrator.Unlock(c.Context) //nolint:errcheck

				group, err := migrator.Migrate(c.Context)
				if err != nil {
					return err
				}
				if group.IsZero() {
					fmt.Printf("there are no new migrations to run (database is up to date)\n")
					return nil
				}
				fmt.Printf("migrated to %s\n", group)
				return nil
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the last migration group",
			Action: func(c *cli.Context) error {
				if err := migrator.Lock(c.Context); err != nil {
					return err
				}
				defer migrator.Unlock(c.Context) //nolint:errcheck

				group, err := migrator.Rollback(c.Context)
				if err != nil {
					return err
				}
				if group.IsZero() {
					fmt.Printf("there are no groups to roll back\n")
					return nil
				}
				fmt.Printf("rolled back %s\n", group)
				return nil
			},
		},
		{
			Name:  "lock",
			Usage: "lock migrations",
			Action: func(c *cli.Context) error {
				return migrator.Lock(c.Context)
			},
		},
		{
			Name:  "unlock",
			Usage: "unlock migrations",
			Action: func(c *cli.Context) error {
				return migrator.Unlock(c.Context)
			},
		},
		{
			Name:  "status",
			Usage: "print migrations status",
			Action: func(c *cli.Context) error {
				ms, err := migrator.MigrationsWithStatus(c.Context)
				if err != nil {
					return err
				}
				fmt.Printf("migrations: %s\n", ms)
				fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
				fmt.Printf("last migration group: %s\n", ms.LastGroup())
				return nil
			},
		},
		{
			Name:  "mark_applied",
			Usage: "mark migrations as applied without actually running them",
			Action: func(c *cli.Context) error {
				group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
				if err != nil {
					return err
				}
				if group.IsZero() {
					fmt.Printf("there are no new migrations to mark as applied\n")
					return nil
				}
				fmt.Printf("marked as applied %s\n", group)
				return nil
			},
		},
	}
}
