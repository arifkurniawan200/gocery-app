package cmd

import (
	"github.com/spf13/cobra"
	"grocery-api/cmd/migration"
	"grocery-api/config"
	"grocery-api/db"
	"grocery-api/internal/app"
	"grocery-api/internal/repository"
	"log"
)

func Start() {
	cfg := config.ReadConfig()
	// root command
	root := &cobra.Command{}

	// command allowed
	cmds := []*cobra.Command{
		{
			Use:   "db:migrate",
			Short: "database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.RunMigration(cfg)
			},
		},
		{
			Use:   "db:seeding",
			Short: "database seeding",
			Run: func(cmd *cobra.Command, args []string) {
				//err := migration.SeedingData(cfg)
				//if err != nil {
				//	log.Fatal(err.Error())
				//}
			},
		},
		{
			Use:   "api",
			Short: "run api server",
			Run: func(cmd *cobra.Command, args []string) {
				db, err := db.NewDatabase(cfg.DB)
				if err != nil {
					log.Fatal(err)
				}

				userRepo := repository.NewUserRepository(db)
				productRepo := repository.NewProductRepository(db)
				app.Run(userRepo, productRepo)
			},
		},
	}
	root.AddCommand(cmds...)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
